package history

import (
	query "github.com/aide-cloud/gorm-normalize"
	"gorm.io/gorm"
)

// LikeInstance 根据字典名称模糊查询
func LikeInstance(keyword string) query.ScopeMethod {
	return func(db *gorm.DB) *gorm.DB {
		if keyword == "" {
			return db
		}
		return db.Where("instance LIKE ?", keyword+"%")
	}
}

// TimeRange 根据时间范围查询
func TimeRange(startTime, endTime int64) query.ScopeMethod {
	return func(db *gorm.DB) *gorm.DB {
		if startTime > 0 && endTime > 0 {
			return db.Where("start_at BETWEEN ? AND ?", startTime, endTime)
		}
		return db
	}
}

// WhereAlarmPages 根据告警页面查询
func WhereAlarmPages(ids []uint) query.ScopeMethod {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) == 0 {
			return db
		}
		return db.Where("alarm_page_id IN (?)", ids)
	}
}
