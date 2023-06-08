package repository

import (
	"strconv"
	"time"

	"github.com/gosimple/slug"
	"github.com/mstgnz/services/entity"
	"gorm.io/gorm"
)

// IBlogRepository interface
type IBlogRepository interface {
	InsertBlog(b entity.Blog) (entity.Blog, error)
	UpdateBlog(b entity.Blog) entity.Blog
	DeleteBlog(b entity.Blog)
	AllBlog() []entity.Blog
	FindBlogByID(blogID uint64) entity.Blog
	GenerateSlug(slug string) string
}

// blogRepository struct
type blogRepository struct {
	connection *gorm.DB
}

// BlogRepository instance
func BlogRepository(dbConn *gorm.DB) IBlogRepository {
	return &blogRepository{
		connection: dbConn,
	}
}

// InsertBlog create blog
func (db *blogRepository) InsertBlog(b entity.Blog) (entity.Blog, error) {

	// slug control
	b.Slug = db.GenerateSlug(b.Title)

	if err := db.connection.Save(&b).Error; err != nil {
		return b, err
	}
	db.connection.Save(&b)
	return b, nil
}

// UpdateBlog update blog
func (db *blogRepository) UpdateBlog(b entity.Blog) entity.Blog {
	db.connection.Save(&b)
	return b
}

// DeleteBlog delete blog
func (db *blogRepository) DeleteBlog(b entity.Blog) {
	db.connection.Delete(&b)
}

// FindBlogByID find id blog
func (db *blogRepository) FindBlogByID(blogID uint64) entity.Blog {
	var blog entity.Blog
	db.connection.Find(&blog, blogID)
	return blog
}

// AllBlog get all blog
func (db *blogRepository) AllBlog() []entity.Blog {
	var blogs []entity.Blog
	db.connection.Preload("User").Find(&blogs)
	return blogs
}

// GenerateSlug generate slug with blog title
func (db *blogRepository) GenerateSlug(title string) string {
	generateSlug := slug.Make(title)
	var blog entity.Blog
	db.connection.Where("slug = ?", generateSlug).First(&blog)
	if len(blog.Slug) > 0 {
		return generateSlug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}
	return generateSlug
}
