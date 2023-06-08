package handler

import (
	"net/http"

	"github.com/mstgnz/services/config"
	"github.com/mstgnz/services/service"
)

// ICommentHandler interface
type ICommentHandler interface {
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

// commentHandler struct
type commentHandler struct {
	commentService service.ICommentService
}

// CommentHandler instance
func CommentHandler(commentService service.ICommentService) ICommentHandler {
	return &commentHandler{
		commentService: commentService,
	}
}

// Insert create comment
func (c *commentHandler) Insert(w http.ResponseWriter, r *http.Request) {
	_ = config.WriteJSON(w, 200, config.Response{Status: true, Message: "Insert"})

	/*var commentCreateDTO dto.CommentCreateDTO
	errDTO := context.ShouldBind(&commentCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			commentCreateDTO.UserID = convertedUserID
		}
		result, err := c.commentService.Insert(commentCreateDTO)
		if err != nil {
			response := helper.BuildErrorResponse("ERROR", err.Error(), err.Error())
			context.JSON(http.StatusBadRequest, response)
		} else {
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)
		}
	}*/
}

// Update comment
func (c *commentHandler) Update(w http.ResponseWriter, r *http.Request) {
	_ = config.WriteJSON(w, 200, config.Response{Status: true, Message: "Update"})

	/*var commentUpdateDTO dto.CommentUpdateDTO
	errDTO := context.ShouldBind(&commentUpdateDTO)
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
	if c.commentService.IsAllowedToEdit(userID, commentUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			commentUpdateDTO.UserID = id
		}
		result := c.commentService.Update(commentUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}*/
}

// Delete comment
func (c *commentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_ = config.WriteJSON(w, 200, config.Response{Status: true, Message: "Delete"})

	/*var comment entity.Comment
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	comment.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.commentService.IsAllowedToEdit(userID, comment.ID) {
		c.commentService.Delete(comment)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}*/
}
