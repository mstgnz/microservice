package dto

// RegisterDTO Register Data Transfer Objects
type RegisterDTO struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

// UserUpdateDTO User Data Transfer Objects
type UserUpdateDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required"`
}

// PassUpdateDTO Password Data Transfer Objects
type PassUpdateDTO struct {
	ID         uint   `json:"id"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"re-password" validate:"required"`
}

// LoginDTO Login Data Transfer Objects
type LoginDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
