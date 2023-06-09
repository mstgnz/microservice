package service

import (
	"github.com/mashingan/smapping"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/entity"
	"github.com/mstgnz/microservice/repository"
)

// IUserService interface
type IUserService interface {
	Update(user dto.UserUpdateDTO) (entity.User, error)
	Profile(userID uint) entity.User
}

// userService struct
type userService struct {
	userRepository repository.IUserRepository
}

// UserService user
func UserService(userRepo repository.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepo,
	}
}

// Update user
func (service *userService) Update(user dto.UserUpdateDTO) (entity.User, error) {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		return userToUpdate, err
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser, nil
}

// Profile user
func (service *userService) Profile(userID uint) entity.User {
	return service.userRepository.ProfileUser(userID)
}
