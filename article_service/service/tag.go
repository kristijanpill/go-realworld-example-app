package service

import "github.com/kristijanpill/go-realworld-example-app/article_service/store"

type TagService struct {
	store store.TagStore
}

func NewTagService(store store.TagStore) *TagService {
	return &TagService{
		store: store,
	}
}