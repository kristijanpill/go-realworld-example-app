package store

import (
	"github.com/kristijanpill/go-realworld-example-app/article_service/model"
	"gorm.io/gorm"
)

type ArticlePostgresStore struct {
	db *gorm.DB
}

func NewArticlePostgresStore(db *gorm.DB) (*ArticlePostgresStore, error) {
	err := db.AutoMigrate(&model.Article{})	
	if err != nil {
		return nil, err
	}

	return &ArticlePostgresStore{
		db: db,
	}, nil
}

func (store *ArticlePostgresStore) Create(article *model.Article) (*model.Article, error) {
	result := store.db.Create(article)

	return article, result.Error
}
