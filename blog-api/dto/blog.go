package dto

import (
	"time"
)

// BlogCreateDTO Blog Create Data Transfer Objects
type BlogCreateDTO struct {
	UserID    uint   `json:"user_id"`
	Title     string `json:"title"`
	ShortText string `json:"short_text"`
	LongText  string `json:"long_text"`
}

// BlogUpdateDTO Blog Update Data Transfer Objects
type BlogUpdateDTO struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Title     string `json:"title"`
	ShortText string `json:"short_text"`
	LongText  string `json:"long_text"`
}

// BlogListDTO Blog List Data Transfer Objects
type BlogListDTO struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	ShortText string    `json:"short_text"`
	LongText  string    `json:"long_text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
