package dto

// RegisterDTO Register Data Transfer Objects
type RegisterDTO struct {
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email" `
	Password  string `json:"password" form:"password" binding:"required"`
}

// UserUpdateDTO User Data Transfer Objects
type UserUpdateDTO struct {
	ID        uint64 `json:"id" form:"id"`
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password,omitempty" form:"password,omitempty"`
}

// LoginDTO Login Data Transfer Objects
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
