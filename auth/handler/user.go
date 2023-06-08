package handler

import (
	"net/http"

	"github.com/mstgnz/services/config"
	"github.com/mstgnz/services/service"
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
	_ = config.WriteJSON(w, 200, config.Response{Status: true, Message: "Update"})

	/*var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)*/
}

// Profile user
func (c *userHandler) Profile(w http.ResponseWriter, r *http.Request) {
	_ = config.WriteJSON(w, 200, config.Response{Status: true, Message: "Profile"})

	/*authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)*/

}
