package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/service"
)

// IUserHandler interface
type IUserHandler interface {
	Update(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
}

// userHandler struct
type userHandler struct {
	userService service.IUserService
}

// UserHandler instance
func UserHandler(userService service.IUserService) IUserHandler {
	return &userHandler{
		userService: userService,
	}
}

// Update user
func (c *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	var userUpdateDTO dto.UserUpdateDTO
	err := config.ReadJSON(w, r, &userUpdateDTO)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request"})
		return
	}
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	userUpdateDTO.ID = userID
	user, _ := c.userService.Update(userUpdateDTO)
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Update successful", Data: user})
}

// Profile user
func (c *userHandler) Profile(w http.ResponseWriter, r *http.Request) {
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	log.Printf("USER ID %v", userID)
	user := c.userService.Profile(userID)
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: fmt.Sprintf("Profile: %d", userID), Data: user})
}
