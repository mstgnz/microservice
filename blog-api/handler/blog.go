package handler

import (
	"net/http"

	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/service"
)

// IBlogHandler interface
type IBlogHandler interface {
	All(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

// blogHandler struct
type blogHandler struct {
	blogService service.IBlogService
}

// BlogHandler instance
func BlogHandler(blogService service.IBlogService) IBlogHandler {
	return &blogHandler{
		blogService: blogService,
	}
}

// All get all blogs
func (c *blogHandler) All(w http.ResponseWriter, r *http.Request) {
	_ = config.WriteJSON(w, 200, config.Response{Status: true, Message: "All"})
	/*var blogs = c.blogService.All()
	res := helper.BuildResponse(true, "OK", blogs)
	context.JSON(http.StatusOK, res)*/
}

// FindByID get by id for blog
func (c *blogHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	_ = config.WriteJSON(w, 200, config.Response{Status: true, Message: "FindBy"})

	/*id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var blog = c.blogService.FindByID(id)
	if (blog == entity.Blog{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", blog)
		context.JSON(http.StatusOK, res)
	}*/
}

// Insert create blog
func (c *blogHandler) Insert(w http.ResponseWriter, r *http.Request) {
	_ = config.WriteJSON(w, 200, config.Response{Status: true, Message: "Insert"})

	/*var blogCreateDTO dto.BlogCreateDTO
	errDTO := context.ShouldBind(&blogCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			blogCreateDTO.UserID = convertedUserID
		}
		result, err := c.blogService.Insert(blogCreateDTO)
		if err != nil {
			response := helper.BuildErrorResponse("ERROR", err.Error(), err.Error())
			context.JSON(http.StatusBadRequest, response)
		} else {
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)
		}
	}*/
}

// Update blog
func (c *blogHandler) Update(w http.ResponseWriter, r *http.Request) {
	_ = config.WriteJSON(w, 200, config.Response{Status: true, Message: "Update"})

	/*var blogUpdateDTO dto.BlogUpdateDTO
	errDTO := context.ShouldBind(&blogUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.blogService.IsAllowedToEdit(userID, blogUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			blogUpdateDTO.UserID = id
		}
		result := c.blogService.Update(blogUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}*/
}

// Delete blog
func (c *blogHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_ = config.WriteJSON(w, 200, config.Response{Status: true, Message: "Delete"})

	/*var blog entity.Blog
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	blog.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.blogService.IsAllowedToEdit(userID, blog.ID) {
		c.blogService.Delete(blog)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}*/
}
