package store

import "github.com/kristijanpill/go-realworld-example-app/article_service/model"

type CommentStore interface {
	Create(*model.Comment) (*model.Comment, error)
}