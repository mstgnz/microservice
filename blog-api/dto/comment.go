package dto

// CommentCreateDTO Comment Create Data Transfer Objects
type CommentCreateDTO struct {
	UserID  uint   `json:"user_id"`
	BlogID  uint   `json:"blog_id"`
	Content string `json:"content"`
}

// CommentUpdateDTO Comment Update Data Transfer Objects
type CommentUpdateDTO struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	BlogID  uint   `json:"blog_id"`
	Content string `json:"content"`
}

// CommentListDTO Comment List Data Transfer Objects
type CommentListDTO struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	BlogID    uint   `json:"blog_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
