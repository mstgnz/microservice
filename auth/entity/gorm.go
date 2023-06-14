package entity

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
