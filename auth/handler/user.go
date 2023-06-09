package handler

import (
	"net/http"

	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/service"
)

// IUserHandler interface
type IUserHandler interface {
	Update(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
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
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err})
		return
	}
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	userUpdateDTO.ID = userID
	user, err := c.userService.Update(userUpdateDTO)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Update failed", Error: err})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Update successful", Data: user})
}

// Profile user
func (c *userHandler) Profile(w http.ResponseWriter, r *http.Request) {
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	user, err := c.userService.Profile(userID)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Profile failed", Error: err})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Profile successful", Data: user})
}

func (c *userHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var passUpdateDTO dto.PassUpdateDTO
	err := config.ReadJSON(w, r, &passUpdateDTO)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err})
		return
	}
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	passUpdateDTO.ID = userID
	err = c.userService.UpdatePassword(passUpdateDTO)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Update failed", Error: err})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Update successful"})
}
