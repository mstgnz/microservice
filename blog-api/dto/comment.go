package dto

// CommentCreateDTO Comment Create Data Transfer Objects
type CommentCreateDTO struct {
	UserID  uint64 `json:"user_id" form:"user_id"`
	BlogID  uint64 `json:"blog_id" form:"blog_id"`
	Content string `json:"content" form:"content" binding:"required"`
}

// CommentUpdateDTO Comment Update Data Transfer Objects
type CommentUpdateDTO struct {
	ID      uint64 `json:"id" form:"id" binding:"required"`
	UserID  uint64 `json:"user_id" form:"user_id"`
	BlogID  uint64 `json:"blog_id" form:"blog_id"`
	Content string `json:"content" form:"content" binding:"required"`
}

// CommentListDTO Comment List Data Transfer Objects
type CommentListDTO struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	BlogID    uint64 `json:"blog_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
