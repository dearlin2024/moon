package alarmmodel

import (
	"github.com/aide-family/moon/pkg/palace/model"
)

const tableNameRealtimeAlarmReceiver = "realtime_alarm_receiver"

// RealtimeAlarmReceiver represents the realtime alarm receiver model
type RealtimeAlarmReceiver struct {
	model.EasyModel
	RealtimeAlarmID    uint32 `gorm:"column:realtime_alarm_id;type:int;not null;comment:告警ID" json:"realtime_alarm_id"`
	AlarmNoticeGroupID uint32 `gorm:"column:alarm_notice_group_id;type:int;not null;comment:告警通知组ID" json:"alarm_notice_group_id"`
}

// TableName overrides the default table name generated by gorm
func (*RealtimeAlarmReceiver) TableName() string {
	return tableNameRealtimeAlarmReceiver
}
