package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/mstgnz/microservice/dto"
	"github.com/mstgnz/microservice/entity"
	"github.com/mstgnz/microservice/repository"
)

// IBlogService interface
type IBlogService interface {
	Insert(b dto.BlogCreateDTO) (entity.Blog, error)
	Update(b dto.BlogUpdateDTO) entity.Blog
	Delete(b entity.Blog)
	All() []dto.BlogListDTO
	FindByID(blogID uint64) entity.Blog
	IsAllowedToEdit(userID string, blogID uint64) bool
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

// Insert blog service
func (service *blogService) Insert(b dto.BlogCreateDTO) (entity.Blog, error) {
	blog := entity.Blog{}
	err := smapping.FillStruct(&blog, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.blogRepository.InsertBlog(blog)
}

// Update blog service
func (service *blogService) Update(b dto.BlogUpdateDTO) entity.Blog {
	blog := entity.Blog{}
	err := smapping.FillStruct(&blog, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.blogRepository.UpdateBlog(blog)
	return res
}

// Delete blog service
func (service *blogService) Delete(b entity.Blog) {
	service.blogRepository.DeleteBlog(b)
}

// All blog service
func (service *blogService) All() []dto.BlogListDTO {
	var list []dto.BlogListDTO
	getList := service.blogRepository.AllBlog()
	for i := 0; i < len(getList); i++ {
		list = append(list, dto.BlogListDTO{
			ID:        getList[i].ID,
			Slug:      getList[i].Slug,
			UserID:    getList[i].UserID,
			Title:     getList[i].Title,
			ShortText: getList[i].ShortText,
			LongText:  getList[i].LongText,
			CreatedAt: getList[i].CreatedAt,
			UpdatedAt: getList[i].UpdatedAt,
		})
	}
	return list
}

// FindByID blog service
func (service *blogService) FindByID(blogID uint64) entity.Blog {
	return service.blogRepository.FindBlogByID(blogID)
}

// IsAllowedToEdit blog service
func (service *blogService) IsAllowedToEdit(userID string, blogID uint64) bool {
	b := service.blogRepository.FindBlogByID(blogID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
