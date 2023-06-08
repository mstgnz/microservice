package entity

import (
	"gorm.io/gorm"
)

// User Entity
type User struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey"`
	FirstName string `gorm:"NOT NULL;size:25"`
	LastName  string `gorm:"NOT NULL;size:25"`
	Email     string `gorm:"NOT NULL;size:100;index:users_email,unique"`
	Password  string `gorm:"NOT NULL;size:255"`
	Token     string `gorm:"NOT NULL;size:255"`
}
