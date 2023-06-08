package dto

import (
	"time"
)

// BlogCreateDTO Blog Create Data Transfer Objects
type BlogCreateDTO struct {
	UserID    uint64 `json:"user_id" form:"user_id"`
	Title     string `json:"title" form:"title" binding:"required"`
	ShortText string `json:"short_text" form:"short_text" binding:"required"`
	LongText  string `json:"long_text" form:"long_text" binding:"required"`
}

// BlogUpdateDTO Blog Update Data Transfer Objects
type BlogUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	UserID    uint64 `json:"user_id" form:"user_id"`
	Title     string `json:"title" form:"title" binding:"required"`
	ShortText string `json:"short_text" form:"short_text" binding:"required"`
	LongText  string `json:"long_text" form:"long_text" binding:"required"`
}

// BlogListDTO Blog List Data Transfer Objects
type BlogListDTO struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	ShortText string    `json:"short_text"`
	LongText  string    `json:"long_text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
