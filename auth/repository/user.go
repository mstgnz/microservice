package repository

import (
	"log"

	"github.com/mstgnz/microservice/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// IUserRepository is contract what userRepository can do to db
type IUserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID uint) entity.User
}

type userRepository struct {
	conn *gorm.DB
}

// UserRepository is creates a new instance of IUserRepository
func UserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		conn: db,
	}
}

func (db *userRepository) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.conn.Save(&user)
	return user
}

func (db *userRepository) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.conn.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.conn.Save(&user)
	return user
}

func (db *userRepository) VerifyCredential(email string, _ string) interface{} {
	var user entity.User
	res := db.conn.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.conn.Where("email = ?", email).Take(&user)
}

func (db *userRepository) FindByEmail(email string) entity.User {
	var user entity.User
	db.conn.Where("email = ?", email).Take(&user)
	return user
}

func (db *userRepository) ProfileUser(userID uint) entity.User {
	var user entity.User
	db.conn.Preload("Blogs").Preload("Blogs.User").Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
