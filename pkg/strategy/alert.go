package strategy

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"
	"prometheus-manager/pkg/util/times"
)

var _ Alerter = (*Alerting)(nil)

type Alerter interface {
	Eval(ctx context.Context) ([]*Alarm, error)
}

type (
	Alerting struct {
		datasourceName DatasourceName
		groups         []*Group
	}

	AlertingOption func(*Alerting)
)

var alarmCache = NewDefaultAlarmCache()

// NewAlerting 初始化策略告警实例
func NewAlerting(groups ...*Group) Alerter {
	a := &Alerting{
		datasourceName: PrometheusDatasource,
		groups:         groups,
	}

	return a
}

func (a *Alerting) Eval(ctx context.Context) ([]*Alarm, error) {
	eg := new(errgroup.Group)
	eg.SetLimit(100)
	alarms := NewAlarmList()
	timeNowUnix := time.Now().Unix()
	strategyIds := make(map[uint32]struct{})
	for _, groupItem := range a.groups {
		group := groupItem
		for _, strategyItem := range group.Rules {
			strategyIds[strategyItem.Id] = struct{}{}
			strategyInfo := &*strategyItem
			eg.Go(func() error {
				datasource := NewDatasource(a.datasourceName, strategyInfo.Endpoint())
				queryResponse, err := datasource.Query(ctx, strategyInfo.Expr, timeNowUnix)
				if err != nil {
					return err
				}
				newAlarmInfo := NewAlarm(group, strategyInfo, queryResponse.Data.Result)
				// 获取该策略下所有已经产生的告警数据
				existAlarmInfo, exist := alarmCache.Get(strategyInfo.Id)
				if !exist {
					log.Info("不存在数据, 存入缓存")
					// 不存在历史数据, 则直接把新告警数据缓存到alarmCache
					alarmCache.Set(strategyInfo.Id, newAlarmInfo)
					// 不需要立即告警
					//alarms.Append(newAlarmInfo)
					return nil
				}

				// 比较两次告警数据, 新数据需要加入alerts, 旧数据需要删除, 并标记为告警恢复
				usableAlarmInfo := a.mergeAlarm(strategyInfo, newAlarmInfo, existAlarmInfo)
				alarms.Append(usableAlarmInfo)
				return nil
			})
		}
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	timeUnix := time.Now().Unix()
	endsAt := time.Unix(timeUnix, 0).Format(times.ParseLayout)
	resolvedAlarmMap := make(map[uint32]*Alarm)
	// 不存在的告警规则直接发送告警恢复通知
	alarmCache.RangeNotifyAlerts(func(alertItem *Alert) {
		ruleId := alertItem.Labels.StrategyId()
		if _, ok := strategyIds[ruleId]; !ok {
			// 告警恢复
			resolvedAlarm, stored := resolvedAlarmMap[ruleId]
			if !stored {
				alarmInfo, exist := alarmCache.Get(ruleId)
				if !exist || alarmInfo == nil {
					return
				}
				resolvedAlarm = alarmInfo
				resolvedAlarm.Alerts = make([]*Alert, 0, 10)
				resolvedAlarmMap[ruleId] = resolvedAlarm
			}
			alertResolvedDetail := alertItem
			alertResolvedDetail.Status = AlarmStatusResolved
			alertResolvedDetail.EndsAt = endsAt
			resolvedAlarm.Alerts = append(resolvedAlarm.Alerts, alertResolvedDetail)
			alarms.Append(resolvedAlarm)
			alarmCache.RemoveNotifyAlert(alertResolvedDetail)
		}
	})
	for ruleId := range resolvedAlarmMap {
		alarmCache.Remove(ruleId)
	}

	return alarms.List(), nil
}

// mergeAlarm 根据新告警数据和旧告警数据, 生成告警数据
func (a *Alerting) mergeAlarm(ruleInfo *Rule, newAlarmInfo, existAlarmInfo *Alarm) *Alarm {
	if newAlarmInfo == nil || existAlarmInfo == nil {
		return nil
	}
	alarm := newAlarmInfo
	existAlertMap := make(map[string]*Alert)
	for _, alert := range existAlarmInfo.Alerts {
		existAlertMap[alert.Fingerprint] = alert
	}

	log.Infow("ruleId", ruleInfo.Id, "existAlarmInfo.Alerts.len", len(existAlarmInfo.Alerts))

	// 初始化一个最大值
	alerts := make([]*Alert, 0, len(newAlarmInfo.Alerts)+len(existAlarmInfo.Alerts))
	// 把新告警set到缓存中
	alarmCache.Set(ruleInfo.Id, newAlarmInfo)
	nowTimeUnix := time.Now().Unix()
	ruleDuration := BuildDuration(ruleInfo.For)
	newAlertMap := make(map[string]*Alert)
	for _, alert := range newAlarmInfo.Alerts {
		alertInfo := alert
		newAlertMap[alert.Fingerprint] = alertInfo
		// 判断告警时常是否满足告警条件, 满足则加入新告警列表
		var eventAt int64
		// 判断告警是否已存在
		if existAlert, ok := existAlertMap[alert.Fingerprint]; ok {
			eventAt = times.ParseAlertTimeUnix(existAlert.StartsAt)
		} else {
			eventAt = times.ParseAlertTimeUnix(alertInfo.StartsAt)
		}

		//log.Infow("eventAt", eventAt, "ruleDuration", ruleDuration, "nowTimeUnix", nowTimeUnix)

		if nowTimeUnix-eventAt >= ruleDuration {
			alerts = append(alerts, alertInfo)
			alarmCache.SetNotifyAlert(alertInfo)
		}
	}

	endsAt := time.Unix(nowTimeUnix, 0).Format(times.ParseLayout)
	for _, oldAlert := range existAlarmInfo.Alerts {
		log.Info("判段告警恢复")
		existAlertTmp, ok := newAlertMap[oldAlert.Fingerprint]
		log.Infow("existAlertTmp", existAlertTmp, "exist", ok, "strategyId", oldAlert.Labels.StrategyId())
		if !ok {
			log.Infow("===============告警恢复通知")
			// 判断是否发送过告警, 如果没有发送过, 不算告警恢复事件
			notifyAlert, notifyOK := alarmCache.GetNotifyAlert(oldAlert)
			if notifyOK && notifyAlert.Status == AlarmStatusFiring {
				notifyAlert.Status = AlarmStatusResolved
				notifyAlert.EndsAt = endsAt
				alerts = append(alerts, notifyAlert)
				alarmCache.RemoveNotifyAlert(notifyAlert)
			}
			continue
		}

	}

	alarm.Alerts = alerts

	return alarm
}

// BuildDuration 字符串转为api时间
func BuildDuration(duration string) int64 {
	durationLen := len(duration)
	if duration == "" || durationLen < 2 {
		return 0
	}
	value, _ := strconv.Atoi(duration[:durationLen-1])
	// 获取字符串最后一个字符
	unit := string(duration[durationLen-1])
	switch unit {
	case "s":
		return int64(value)
	case "m":
		return int64(value) * 60
	case "h":
		return int64(value) * 60 * 60
	case "d":
		return int64(value) * 60 * 60 * 24
	default:
		return 0
	}
}

// WithDatasourceName 设置数据源名称
func WithDatasourceName(datasourceName DatasourceName) AlertingOption {
	return func(a *Alerting) {
		a.datasourceName = datasourceName
	}
}

// SetAlarmCache 设置告警缓存
func SetAlarmCache(cache AlarmCache) {
	alarmCache = cache
}