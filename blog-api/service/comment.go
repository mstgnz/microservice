package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/mstgnz/services/dto"
	"github.com/mstgnz/services/entity"
	"github.com/mstgnz/services/repository"
)

// ICommentService interface
type ICommentService interface {
	Insert(b dto.CommentCreateDTO) (entity.Comment, error)
	Update(b dto.CommentUpdateDTO) entity.Comment
	Delete(b entity.Comment)
	IsAllowedToEdit(userID string, commentID uint64) bool
}

// commentService struct
type commentService struct {
	commentRepository repository.ICommentRepository
}

// CommentService instance
func CommentService(commentRepo repository.ICommentRepository) ICommentService {
	return &commentService{
		commentRepository: commentRepo,
	}
}

// Insert comment service
func (service *commentService) Insert(b dto.CommentCreateDTO) (entity.Comment, error) {
	comment := entity.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.commentRepository.InsertComment(comment)
}

// Update comment service
func (service *commentService) Update(b dto.CommentUpdateDTO) entity.Comment {
	comment := entity.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.commentRepository.UpdateComment(comment)
	return res
}

// Delete comment service
func (service *commentService) Delete(b entity.Comment) {
	service.commentRepository.DeleteComment(b)
}

// IsAllowedToEdit blog service
func (service *commentService) IsAllowedToEdit(userID string, commentID uint64) bool {
	b := service.commentRepository.FindCommentByID(commentID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
