package repository

import (
	"strconv"
	"time"

	"github.com/gosimple/slug"
	"github.com/mstgnz/microservice/entity"
	"gorm.io/gorm"
)

// BlogRepository interface
type BlogRepository interface {
	All() ([]entity.Blog, error)
	Delete(id uint) error
	Create(b entity.Blog) (entity.Blog, error)
	Update(b entity.Blog) (entity.Blog, error)
	FindByID(id uint) (entity.Blog, error)
	FindBySlug(slug string) (entity.Blog, error)
	GenerateSlug(slug string) string
}

// blogRepository struct
type blogRepository struct {
	connection *gorm.DB
}

// NewBlogRepository instance
func NewBlogRepository(dbConn *gorm.DB) BlogRepository {
	return &blogRepository{
		connection: dbConn,
	}
}

func (db *blogRepository) All() ([]entity.Blog, error) {
	var blogs []entity.Blog
	tx := db.connection.Preload("User").Find(&blogs)
	return blogs, tx.Error
}

func (db *blogRepository) Create(b entity.Blog) (entity.Blog, error) {
	// slug control
	b.Slug = db.GenerateSlug(b.Title)
	tx := db.connection.Save(&b)
	return b, tx.Error
}

func (db *blogRepository) Update(b entity.Blog) (entity.Blog, error) {
	// slug control
	b.Slug = db.GenerateSlug(b.Title)
	// it only updates the filled values
	tx := db.connection.Model(&b).Updates(&b)
	return b, tx.Error
}

func (db *blogRepository) Delete(id uint) error {
	tx := db.connection.Delete(entity.Blog{}, id)
	return tx.Error
}

func (db *blogRepository) FindByID(id uint) (entity.Blog, error) {
	var blog entity.Blog
	tx := db.connection.Preload("User").Preload("Comments").First(&blog, id)
	return blog, tx.Error
}

func (db *blogRepository) FindBySlug(slug string) (entity.Blog, error) {
	var blog entity.Blog
	tx := db.connection.Preload("User").Preload("Comments").First(&blog, "slug = ?", slug)
	return blog, tx.Error
}

// GenerateSlug generate slug with blog title
func (db *blogRepository) GenerateSlug(title string) string {
	if len(title) > 0 {
		generateSlug := slug.Make(title)
		var blog entity.Blog
		db.connection.First(&blog, "slug = ?", generateSlug)
		if len(blog.Slug) > 0 {
			return generateSlug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
		}
		return generateSlug
	}
	return title
}
