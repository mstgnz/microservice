package entity

import (
	"gorm.io/gorm"
)

// Blog Entity
type Blog struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey"`
	UserID    uint64 `gorm:"NOT NULL"`
	Title     string `gorm:"NOT NULL;size:255"`
	Slug      string `gorm:"NOT NULL;size:255;index:blogs_slug,unique"`
	ShortText string `gorm:"NOT NULL;size:500"`
	LongText  string `gorm:"NOT NULL"`
	User      User   `gorm:"constraint:OnDelete:CASCADE"`
}
