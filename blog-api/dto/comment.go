package dto

import (
	"time"
)

type CommentCreate struct {
	UserID  uint   `json:"user_id"`
	BlogID  uint   `json:"blog_id" validate:"required"`
	Content string `json:"content" validate:"required,min=10"`
}

type CommentUpdate struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	BlogID  uint   `json:"blog_id" validate:"required"`
	Content string `json:"content" validate:"required,min=10"`
}

type Comment struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	BlogID    uint      `json:"blog_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user,omitempty"`
}

type CommentDelete struct {
	ID     uint `json:"id" validate:"required"`
	UserID uint `json:"user_id"`
}
