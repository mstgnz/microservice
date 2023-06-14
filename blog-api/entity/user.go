package entity

import (
	"time"

	"gorm.io/gorm"
)

// User Entity - Note: When using the GormModel, the mapping is not smooth.
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	FirstName string         `gorm:"NOT NULL;size:25"`
	LastName  string         `gorm:"NOT NULL;size:25"`
	Email     string         `gorm:"NOT NULL;size:100;index:users_email,unique"`
	Password  string         `gorm:"NOT NULL;size:255"`
	Token     string         `gorm:"NOT NULL;size:255"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
