package service

import (
	"errors"

	"github.com/mashingan/smapping"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/entity"
	"github.com/mstgnz/microservice/repository"
)

// ICommentService interface
type ICommentService interface {
	Create(b dto.CommentCreate) (dto.Comment, error)
	Update(b dto.CommentUpdate) (dto.Comment, error)
	Delete(b dto.CommentDelete) error
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

// Create comment service
func (service *commentService) Create(b dto.CommentCreate) (dto.Comment, error) {
	comment := entity.Comment{}
	commentDto := dto.Comment{}
	// mapping
	err := smapping.FillStruct(&comment, smapping.MapFields(&b))
	if err != nil {
		return commentDto, err
	}
	// create
	comment, err = service.commentRepository.Create(comment)
	if err != nil {
		return commentDto, err
	}
	// mapping
	err = smapping.FillStruct(&commentDto, smapping.MapFields(&comment))
	return commentDto, err
}

// Update comment service
func (service *commentService) Update(b dto.CommentUpdate) (dto.Comment, error) {
	comment := entity.Comment{}
	commentDto := dto.Comment{}
	// is owner
	find, err := service.commentRepository.Find(b.ID)
	if err != nil {
		return commentDto, err
	}
	if find.UserID != b.UserID {
		return commentDto, errors.New("this content does not belong to you")
	}
	if find.BlogID != b.BlogID {
		return commentDto, errors.New("this content does not belong to blog")
	}
	// mapping
	err = smapping.FillStruct(&comment, smapping.MapFields(&b))
	if err != nil {
		return commentDto, err
	}
	// update
	comment, err = service.commentRepository.Update(comment)
	if err != nil {
		return commentDto, err
	}
	// mapping
	err = smapping.FillStruct(&commentDto, smapping.MapFields(&comment))
	return commentDto, err
}

// Delete comment service
func (service *commentService) Delete(b dto.CommentDelete) error {
	// is owner
	find, err := service.commentRepository.Find(b.ID)
	if err != nil {
		return err
	}
	if find.UserID != b.UserID {
		return errors.New("this content does not belong to you")
	}
	// delete
	return service.commentRepository.Delete(b.ID)
}
