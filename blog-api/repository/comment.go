package repository

import (
	"github.com/mstgnz/microservice/entity"
	"gorm.io/gorm"
)

// ICommentRepository interface
type ICommentRepository interface {
	InsertComment(b entity.Comment) (entity.Comment, error)
	UpdateComment(b entity.Comment) entity.Comment
	DeleteComment(b entity.Comment)
	FindCommentByID(commentID uint64) entity.Comment
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

// InsertComment create comment
func (db *commentRepository) InsertComment(b entity.Comment) (entity.Comment, error) {
	if err := db.connection.Save(&b).Error; err != nil {
		return b, err
	}
	db.connection.Save(&b)
	return b, nil
}

// UpdateComment update comment
func (db *commentRepository) UpdateComment(b entity.Comment) entity.Comment {
	db.connection.Save(&b)
	return b
}

// DeleteComment delete comment
func (db *commentRepository) DeleteComment(b entity.Comment) {
	db.connection.Delete(&b)
}

// FindCommentByID find id blog
func (db *commentRepository) FindCommentByID(commentID uint64) entity.Comment {
	var comment entity.Comment
	db.connection.Find(&comment, commentID)
	return comment
}
