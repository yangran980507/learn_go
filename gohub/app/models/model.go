// Package models 模型通用属性和方法
package models

import (
	"github.com/spf13/cast"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`

	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at_at,omitempty"`
}

func (m BaseModel) GetStringID() string {
	return cast.ToString(m.ID)
}
