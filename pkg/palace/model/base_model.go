package model

import (
	"context"

	"github.com/aide-family/moon/pkg/helper/middleware"
	"github.com/aide-family/moon/pkg/util/types"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	ctx context.Context `gorm:"-"`

	ID        uint32     `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt types.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt types.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`

	// 创建人
	CreatorID uint32   `gorm:"column:creator;type:int unsigned;not null;comment:创建者" json:"creator_id"`
	Creator   *SysUser `gorm:"foreignKey:CreatorID" json:"creator"`
}

type DeleteAtFieldModel struct {
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint;not null;default:0;" json:"deleted_at"`
}

type AllFieldModel struct {
	BaseModel
	DeleteAtFieldModel
}

// WithContext 获取上下文
func (u *BaseModel) WithContext(ctx context.Context) *BaseModel {
	u.ctx = ctx
	return u
}

func (u *BaseModel) BeforeCreate(_ *gorm.DB) (err error) {
	claims, _ := middleware.ParseJwtClaims(u.ctx)
	u.CreatorID = claims.GetUser()
	return
}
