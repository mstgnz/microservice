package entity

import (
	"time"

	"gorm.io/gorm"
)

// Blog Entity - Note: When using the GormModel, the mapping is not smooth.
type Blog struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"NOT NULL"`
	Title     string         `gorm:"NOT NULL;size:255"`
	Slug      string         `gorm:"NOT NULL;size:255;index:blogs_slug,unique"`
	ShortText string         `gorm:"NOT NULL;size:500"`
	LongText  string         `gorm:"NOT NULL"`
	User      User           `gorm:"constraint:OnDelete:CASCADE"`
	Comments  []Comment      `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
