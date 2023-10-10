package entity

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created-at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated-at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
