package dto

// RegisterDTO Register Data Transfer Objects
type RegisterDTO struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// UserUpdateDTO User Data Transfer Objects
type UserUpdateDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

// PassUpdateDTO Password Data Transfer Objects
type PassUpdateDTO struct {
	ID         uint   `json:"id"`
	Password   string `json:"password"`
	RePassword string `json:"re-password"`
}

// LoginDTO Login Data Transfer Objects
type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
