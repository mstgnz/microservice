package handler

import (
	"net/http"
	"strconv"

	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/entity"
	"github.com/mstgnz/microservice/service"
)

type IAuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authService service.IAuthService
}

func AuthHandler(authService service.IAuthService) IAuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (c *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginDTO dto.LoginDTO
	err := config.ReadJSON(w, r, &loginDTO)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := config.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Login successful", Data: generatedToken})
		return
	}
	_ = config.WriteJSON(w, http.StatusUnauthorized, config.Response{Status: false, Message: "Invalid credential"})
	return
}

func (c *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerDTO dto.RegisterDTO
	err := config.ReadJSON(w, r, &registerDTO)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		_ = config.WriteJSON(w, http.StatusConflict, config.Response{Status: false, Message: "Email already exists"})
		return
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := config.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		_ = config.WriteJSON(w, http.StatusCreated, config.Response{Status: true, Message: "Register successful", Data: token})
		return
	}
}
