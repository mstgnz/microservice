package dto

import (
	"time"
)

type BlogCreate struct {
	UserID    uint   `json:"user_id"`
	Title     string `json:"title" validate:"required,min=10,max=100"`
	Slug      string `json:"slug" validate:"omitempty"`
	ShortText string `json:"short_text" validate:"required,min=100,max=200"`
	LongText  string `json:"long_text" validate:"required,min=200"`
}

type BlogUpdate struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Title     string `json:"title" validate:"omitempty,min=10,max=100"`
	ShortText string `json:"short_text" validate:"omitempty,min=100,max=200"`
	LongText  string `json:"long_text" validate:"omitempty,min=200"`
}

type BlogList struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	ShortText string    `json:"short_text"`
	LongText  string    `json:"long_text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
}

type Blog struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	ShortText string    `json:"short_text"`
	LongText  string    `json:"long_text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
	Comments  []Comment `json:"comments"`
}

type BlogDelete struct {
	ID     uint   `json:"id" validate:"required"`
	UserID uint   `json:"user_id"`
	Slug   string `json:"slug"`
}
