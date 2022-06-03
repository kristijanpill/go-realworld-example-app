package store

import "github.com/kristijanpill/go-realworld-example-app/article_service/model"

type TagStore interface {
	Create(*model.Tag) (*model.Tag, error)
	FindByName(name string) (*model.Tag, error)
}