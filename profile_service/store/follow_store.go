package store

import (
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
)

type FollowStore interface {
	Create(*model.Follow) (*model.Follow, error)
	DeleteByProfileIdAndTargetId(profileId, targetId string) error
	ExistsByProfileIdAndTargetId(profileId, targetId string) bool
	FindAllByProfileId(profileId string) ([]*model.Follow, error)
}