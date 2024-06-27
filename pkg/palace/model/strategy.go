package model

import (
	"encoding/json"

	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"

	"gorm.io/plugin/soft_delete"
)

const TableNameStrategy = "strategies"

type Strategy struct {
	ID        uint32                `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt types.Time            `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt types.Time            `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint;not null;comment:删除时间" json:"deleted_at"`

	// 模板ID
	StrategyTemplateID uint32 `gorm:"column:strategy_template_id;type:int unsigned;not null;comment:策略模板ID" json:"strategy_template_id"`
	// 策略模板
	StrategyTemplate *StrategyTemplate `gorm:"foreignKey:StrategyTemplateID" json:"strategy_template"`

	Alert       string           `gorm:"column:alert;type:varchar(64);not null;comment:策略名称" json:"alert"`
	Expr        string           `gorm:"column:expr;type:text;not null;comment:告警表达式" json:"expr"`
	Status      vobj.Status      `gorm:"column:status;type:int;not null;comment:策略状态" json:"status"`
	Remark      string           `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Labels      vobj.Labels      `gorm:"column:labels;type:JSON;not null;comment:标签" json:"labels"`
	Annotations vobj.Annotations `gorm:"column:annotations;type:JSON;not null;comment:注解" json:"annotations"`
	// 告警等级数据
	StrategyLevelTemplates []*StrategyLevelTemplate `gorm:"foreignKey:StrategyID" json:"strategy_level_templates"`

	CreatorID uint32   `gorm:"column:creator;type:int unsigned;not null;comment:创建者" json:"creator_id"`
	Creator   *SysUser `gorm:"foreignKey:CreatorID" json:"creator"`

	// 公共通知对象
}

// String json string
func (c *Strategy) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

func (c *Strategy) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *Strategy) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

// TableName Strategy's table name
func (*Strategy) TableName() string {
	return TableNameStrategy
}