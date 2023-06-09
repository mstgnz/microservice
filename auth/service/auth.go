package service

import (
	"github.com/mashingan/smapping"
	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/entity"
	"github.com/mstgnz/microservice/repository"
)

type IAuthService interface {
	VerifyCredential(email string, password string) (entity.User, error)
	CreateUser(user dto.RegisterDTO) (entity.User, error)
	FindByEmail(email string) bool
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.IUserRepository
}

func AuthService(userRep repository.IUserRepository) IAuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) (entity.User, error) {
	user, err := service.userRepository.VerifyCredential(email, password)
	if err == nil && user.Password != "" {
		comparedPassword := config.ComparePassword(user.Password, []byte(password))
		if user.Email == email && comparedPassword {
			return user, nil
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
	return service.userRepository.InsertUser(user)
}

func (service *authService) FindByEmail(email string) bool {
	user, err := service.userRepository.FindByEmail(email)
	if err != nil || user.Email == "" {
		return false
	}
	return true
}

func (service *authService) IsDuplicateEmail(email string) bool {
	user, err := service.userRepository.IsDuplicateEmail(email)
	return err != nil || user.Email == ""
}
