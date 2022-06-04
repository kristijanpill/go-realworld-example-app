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