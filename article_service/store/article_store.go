package store

import "github.com/kristijanpill/go-realworld-example-app/article_service/model"

type ArticleStore interface {
	Create(*model.Article) (*model.Article, error)
}