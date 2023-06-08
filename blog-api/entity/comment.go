package entity

import (
	"gorm.io/gorm"
)

// Comment Entity
type Comment struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey"`
	UserID  uint64 `gorm:"NOT NULL"`
	BlogID  uint64 `gorm:"NOT NULL"`
	Content string `gorm:"NOT NULL"`
	User    User   `gorm:"constraint:OnDelete:CASCADE"`
	Blog    Blog   `gorm:"constraint:OnDelete:CASCADE"`
}
