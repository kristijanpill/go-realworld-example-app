package store

import "github.com/kristijanpill/go-realworld-example-app/article_service/model"

type ArticleStore interface {
	Create(*model.Article) (*model.Article, error)
	Find(offset, limit int32) ([]*model.Article, error)
	FindByTag(offset, limit int32, tag string) ([]*model.Article, error)
	FindByAuthor(offset, limit int32, userId string) ([]*model.Article, error)
}