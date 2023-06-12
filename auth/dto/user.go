package dto

// RegisterDTO Register Data Transfer Objects
type RegisterDTO struct {
	FirstName string `json:"firstname" validate:"required,min=2"`
	LastName  string `json:"lastname" validate:"required,min=2"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

// UserUpdateDTO User Data Transfer Objects
type UserUpdateDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname" validate:"required,min=2"`
	LastName  string `json:"lastname" validate:"required,min=2"`
	Email     string `json:"email" validate:"required,email"`
}

// PassUpdateDTO Password Data Transfer Objects
type PassUpdateDTO struct {
	ID         uint   `json:"id"`
	Password   string `json:"password" validate:"required,min=6"`
	RePassword string `json:"re-password" validate:"required,min=6"`
}

// LoginDTO Login Data Transfer Objects
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
