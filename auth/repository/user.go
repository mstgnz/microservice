package repository

import (
	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/entity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	VerifyCredential(email string, password string) (entity.User, error)
	IsDuplicateEmail(email string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	ProfileUser(userID uint) (entity.User, error)
	UpdatePassword(pass dto.PassUpdateDTO) error
}

type userRepository struct {
	conn *gorm.DB
}

func UserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		conn: db,
	}
}

func (db *userRepository) InsertUser(user entity.User) (entity.User, error) {
	user.Password = config.HashAndSalt([]byte(user.Password))
	tx := db.conn.Save(&user)
	return user, tx.Error
}

func (db *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	tx := db.conn.Model(&user).Updates(entity.User{FirstName: user.FirstName, LastName: user.LastName})
	return user, tx.Error
}

func (db *userRepository) VerifyCredential(email string, password string) (entity.User, error) {
	var user entity.User
	tx := db.conn.Where("email = ?", email, "password = ?", password).Take(&user)
	return user, tx.Error
}

func (db *userRepository) IsDuplicateEmail(email string) (entity.User, error) {
	var user entity.User
	tx := db.conn.Where("email = ?", email).Take(&user)
	return user, tx.Error
}

func (db *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	tx := db.conn.Where("email = ?", email).Take(&user)
	return user, tx.Error
}

func (db *userRepository) ProfileUser(userID uint) (entity.User, error) {
	var user entity.User
	tx := db.conn.Find(&user, userID)
	return user, tx.Error

}

func (db *userRepository) UpdatePassword(pass dto.PassUpdateDTO) error {
	tx := db.conn.Model(&entity.User{}).Where("id = ?", pass.ID).Update("password", pass.Password)
	return tx.Error
}
