package service

import (
	"errors"

	"github.com/mashingan/smapping"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/entity"
	"github.com/mstgnz/microservice/repository"
)

// IBlogService interface
type IBlogService interface {
	All() ([]dto.BlogList, error)
	Create(b dto.BlogCreate) (dto.BlogList, error)
	Update(b dto.BlogUpdate) (dto.BlogList, error)
	Delete(b dto.BlogDelete) error
	Find(slug string) (dto.BlogList, error)
}

// blogService struct
type blogService struct {
	blogRepository repository.IBlogRepository
}

// BlogService instance
func BlogService(blogRepo repository.IBlogRepository) IBlogService {
	return &blogService{
		blogRepository: blogRepo,
	}
}

func (service *blogService) Create(b dto.BlogCreate) (dto.BlogList, error) {
	blog := entity.Blog{}
	blogDto := dto.BlogList{}
	err := smapping.FillStruct(&blog, smapping.MapFields(&b))
	if err != nil {
		return blogDto, err
	}
	blog, err = service.blogRepository.Create(blog)
	if err != nil {
		return blogDto, err
	}
	err = smapping.FillStruct(&blogDto, smapping.MapFields(&blog))
	return blogDto, err
}

// Update blog service
func (service *blogService) Update(b dto.BlogUpdate) (dto.BlogList, error) {
	blog := entity.Blog{}
	blogDto := dto.BlogList{}
	// is owner
	find, err := service.blogRepository.FindByID(b.ID)
	if err != nil {
		return blogDto, err
	}
	if find.UserID != b.UserID {
		return blogDto, errors.New("this content does not belong to you")
	}
	// mapping
	err = smapping.FillStruct(&blog, smapping.MapFields(&b))
	if err != nil {
		return blogDto, err
	}
	// update
	blog, err = service.blogRepository.Update(blog)
	if err != nil {
		return blogDto, err
	}
	// mapping
	err = smapping.FillStruct(&blogDto, smapping.MapFields(&blog))
	return blogDto, err
}

// Delete blog service
func (service *blogService) Delete(b dto.BlogDelete) error {
	// is owner
	find, err := service.blogRepository.FindByID(b.ID)
	if err != nil {
		return err
	}
	if find.UserID != b.UserID {
		return errors.New("this content does not belong to you")
	}
	// delete
	return service.blogRepository.Delete(b.ID)
}

// All blog service
func (service *blogService) All() ([]dto.BlogList, error) {
	var blogDto []dto.BlogList
	var bd dto.BlogList
	blogs, err := service.blogRepository.All()
	if err != nil {
		return blogDto, err
	}
	for i := 0; i < len(blogs); i++ {
		err = smapping.FillStruct(&bd, smapping.MapFields(&blogs[i]))
		blogDto = append(blogDto, bd)
	}
	return blogDto, err
}

func (service *blogService) Find(slug string) (dto.BlogList, error) {
	var blogDto dto.BlogList
	blog, err := service.blogRepository.FindBySlug(slug)
	if err != nil {
		return blogDto, err
	}
	err = smapping.FillStruct(&blogDto, smapping.MapFields(&blog))
	return blogDto, err
}
