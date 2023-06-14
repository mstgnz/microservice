package repository

import (
	"github.com/mstgnz/microservice/entity"
	"gorm.io/gorm"
)

// ICommentRepository interface
type ICommentRepository interface {
	Create(b entity.Comment) (entity.Comment, error)
	Update(b entity.Comment) (entity.Comment, error)
	Delete(id uint) error
	Find(commentID uint) (entity.Comment, error)
}

// commentRepository struct
type commentRepository struct {
	connection *gorm.DB
}

// CommentRepository instance
func CommentRepository(dbConn *gorm.DB) ICommentRepository {
	return &commentRepository{
		connection: dbConn,
	}
}

// Create comment
func (db *commentRepository) Create(b entity.Comment) (entity.Comment, error) {
	tx := db.connection.Save(&b)
	return b, tx.Error
}

// Update comment
func (db *commentRepository) Update(b entity.Comment) (entity.Comment, error) {
	// it only updates the filled values
	tx := db.connection.Model(&b).Updates(&b)
	return b, tx.Error
}

// Delete comment
func (db *commentRepository) Delete(id uint) error {
	tx := db.connection.Delete(entity.Comment{}, id)
	return tx.Error
}

// Find id blog
func (db *commentRepository) Find(commentID uint) (entity.Comment, error) {
	var comment entity.Comment
	tx := db.connection.Find(&comment, commentID)
	return comment, tx.Error
}
