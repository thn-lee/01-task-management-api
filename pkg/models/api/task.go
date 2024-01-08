package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Task struct {
	ID          uint                  `form:"id" json:"id" gorm:"size:32;primaryKey"`
	Title       string                `form:"title" json:"title" gorm:"size:64;index"`
	Description string                `form:"description" json:"description" gorm:"type:text"`
	Status      string                `form:"status" json:"status" gorm:"size:64;index"`
	CreatedAt   time.Time             `form:"created_at" json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt   time.Time             `form:"updated_at" json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   soft_delete.DeletedAt `form:"deleted_at" json:"deleted_at,omitempty" gorm:"index"`
}

type TaskBody struct {
	Title       string `form:"title" json:"title" gorm:"size:64;index"`
	Description string `form:"description" json:"description" gorm:"type:text"`
	Status      int    `form:"status" json:"status" gorm:"size:64;index"`
}
