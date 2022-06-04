package store

import (
	"github.com/kristijanpill/go-realworld-example-app/article_service/model"
	"gorm.io/gorm"
)

type CommentPostgresStore struct {
	db *gorm.DB
}

func NewCommentPostgresStore(db *gorm.DB) (*CommentPostgresStore, error) {
	err := db.AutoMigrate(&model.Comment{})
	if err != nil {
		return nil, err
	}

	return &CommentPostgresStore{
		db: db,
	}, nil
}

func (store *CommentPostgresStore) Create(comment *model.Comment) (*model.Comment, error) {
	result := store.db.Create(comment)

	return comment, result.Error
}

func (store *CommentPostgresStore) FindById(id int32) (*model.Comment, error) {
	var comment model.Comment
	result := store.db.Where("id = ?", id).First(&comment)

	return &comment, result.Error
}
func (store *CommentPostgresStore) Delete(comment *model.Comment) error {
	result := store.db.Delete(comment)

	return result.Error
}