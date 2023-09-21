package service

import (
	"errors"

	"github.com/mashingan/smapping"
	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/entity"
	"github.com/mstgnz/microservice/repository"
)

type AuthService interface {
	VerifyCredential(email string, password string) (entity.User, error)
	CreateUser(user dto.RegisterDTO) (entity.User, error)
	FindByEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) (entity.User, error) {
	user, err := service.userRepository.FindByEmail(email)
	if err == nil {
		comparedPassword := config.ComparePassword(user.Password, password)
		if !comparedPassword {
			return user, errors.New("information could not be verified")
		}
	}
	return user, err
}

func (service *authService) CreateUser(userDto dto.RegisterDTO) (entity.User, error) {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&userDto))
	if err != nil {
		return user, err
	}
	user.Password = config.HashAndSalt(user.Password)
	return service.userRepository.InsertUser(user)
}

func (service *authService) FindByEmail(email string) bool {
	user, err := service.userRepository.FindByEmail(email)
	if err != nil || user.Email == "" {
		return false
	}
	return true
}
