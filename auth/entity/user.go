package entity

// User Entity
type User struct {
	GormModel
	FirstName string `gorm:"NOT NULL;size:25" json:"firstname"`
	LastName  string `gorm:"NOT NULL;size:25" json:"lastname"`
	Email     string `gorm:"NOT NULL;size:100;index:users_email,unique" json:"email"`
	Password  string `gorm:"NOT NULL;size:255" json:"-"`
	Token     string `gorm:"NOT NULL;size:255" json:"token"`
}
