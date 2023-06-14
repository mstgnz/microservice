package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/service"
)

// ICommentHandler interface
type ICommentHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
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

// Create comment
func (c *commentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var commentCreate dto.CommentCreate
	// body to struct
	err := config.ReadJSON(w, r, &commentCreate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	// struct to validate
	err = config.Validate(commentCreate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	commentCreate.UserID = userID

	comment, err := c.commentService.Create(commentCreate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Comment create successful", Data: comment})
}

// Update comment
func (c *commentHandler) Update(w http.ResponseWriter, r *http.Request) {
	var commentUpdate dto.CommentUpdate
	// body to struct
	err := config.ReadJSON(w, r, &commentUpdate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	// struct to validate
	err = config.Validate(commentUpdate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	id := chi.URLParam(r, "id")
	i, _ := strconv.Atoi(id)
	commentUpdate.ID = uint(i)
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	commentUpdate.UserID = userID

	comment, err := c.commentService.Update(commentUpdate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Comment update successful", Data: comment})
}

// Delete comment
func (c *commentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, _ := strconv.Atoi(id)
	var commentDelete dto.CommentDelete
	commentDelete.ID = uint(i)
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	commentDelete.UserID = userID
	err := c.commentService.Delete(commentDelete)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Comment delete successful", Data: commentDelete})
}
