package service

import (
	"github.com/kristijanpill/go-realworld-example-app/article_service/store"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
)

type TagService struct {
	store store.TagStore
}

func NewTagService(store store.TagStore) *TagService {
	return &TagService{
		store: store,
	}
}

func (service *TagService) GetTags() (*pb.TagsResponse, error) {
	tags, err := service.store.FindAll()
	if err != nil {
		return nil, err
	}

	response := &pb.TagsResponse{
		Tags: []string{},
	}

	for _, tag := range tags {
		response.Tags = append(response.Tags, tag.Name)
	}

	return response, nil
}