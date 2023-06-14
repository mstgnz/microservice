package entity

import (
	"time"

	"gorm.io/gorm"
)

// Comment Entity - Note: When using the GormModel, the mapping is not smooth.
type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"NOT NULL"`
	BlogID    uint           `gorm:"NOT NULL"`
	Content   string         `gorm:"NOT NULL"`
	User      User           `gorm:"constraint:OnDelete:CASCADE"`
	Blog      Blog           `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
