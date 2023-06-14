package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/service"
)

// IBlogHandler interface
type IBlogHandler interface {
	All(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
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

func (c *blogHandler) All(w http.ResponseWriter, _ *http.Request) {
	blogs, err := c.blogService.All()
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "All", Data: blogs})
}

func (c *blogHandler) Find(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	blog, err := c.blogService.Find(slug)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Blog successful", Data: blog})

}

func (c *blogHandler) Create(w http.ResponseWriter, r *http.Request) {
	var blogCreate dto.BlogCreate
	// body to struct
	err := config.ReadJSON(w, r, &blogCreate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	// struct to validate
	err = config.Validate(blogCreate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	blogCreate.UserID = userID

	blogs, err := c.blogService.Create(blogCreate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Blog create successful", Data: blogs})
}

// Update blog
func (c *blogHandler) Update(w http.ResponseWriter, r *http.Request) {
	var blogUpdate dto.BlogUpdate
	// body to struct
	err := config.ReadJSON(w, r, &blogUpdate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	// struct to validate
	err = config.Validate(blogUpdate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusBadRequest, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	id := chi.URLParam(r, "id")
	i, _ := strconv.Atoi(id)
	blogUpdate.ID = uint(i)
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	blogUpdate.UserID = userID

	blog, err := c.blogService.Update(blogUpdate)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Blog update successful", Data: blog})
}

// Delete blog
func (c *blogHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, _ := strconv.Atoi(id)
	var blogDelete dto.BlogDelete
	blogDelete.ID = uint(i)
	userID, _ := config.GetUserIDByToken(r.Header.Get("Authorization"))
	blogDelete.UserID = userID
	err := c.blogService.Delete(blogDelete)
	if err != nil {
		_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: false, Message: "Failed to process request", Error: err.Error()})
		return
	}
	_ = config.WriteJSON(w, http.StatusOK, config.Response{Status: true, Message: "Blog delete successful", Data: blogDelete})
}
