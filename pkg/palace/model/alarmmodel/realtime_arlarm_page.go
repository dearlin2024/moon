package alarmmodel

import (
	"github.com/aide-family/moon/pkg/palace/model"
)

const tableNameRealtimeAlarmPage = "realtime_alarm_page"

// RealtimeAlarmPage represents the realtime alarm pages model
type RealtimeAlarmPage struct {
	model.EasyModel
	RealtimeAlarmID uint32 `gorm:"column:realtime_alarm_id;type:int;not null;comment:告警ID" json:"realtime_alarm_id"`
	PageID          uint32 `gorm:"column:page_id;type:int;not null;comment:页面ID" json:"page_id"`
}

// TableName overrides the default table name generated by gorm
func (*RealtimeAlarmPage) TableName() string {
	return tableNameRealtimeAlarmPage
}
