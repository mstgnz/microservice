package handler

import (
	"net/http"

	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/service"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (c *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginDTO dto.LoginDTO
	// body to struct
	err := config.ReadJSON(w, r, &loginDTO)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	// struct to validate
	err = config.Validate(loginDTO)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	user, err := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	generatedToken, err := config.GenerateToken(user.ID)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	user.Token = generatedToken
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Login successful", Data: user})
}

func (c *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerDTO dto.RegisterDTO
	// body to struct
	err := config.ReadJSON(w, r, &registerDTO)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	// struct to validate
	err = config.Validate(registerDTO)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	if !c.authService.FindByEmail(registerDTO.Email) {
		_ = config.WriteJSON(w, http.StatusConflict, config.Response{Status: false, Message: "Email already exists"})
		return
	} else {
		user, err := c.authService.CreateUser(registerDTO)
		if err != nil {
			_ = config.WriteJSON(w, http.StatusCreated, config.Response{Status: false, Message: "Register error", Error: err.Error()})
			return
		}
		token, err := config.GenerateToken(user.ID)
		if err != nil {
			_ = config.WriteJSON(w, http.StatusCreated, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
			return
		}
		user.Token = token
		_ = config.WriteJSON(w, http.StatusCreated, config.Response{Status: true, Message: "Register successful", Data: user})
		return
	}
}
