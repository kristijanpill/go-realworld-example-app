package store

import (
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
)

type FollowStore interface {
	Create(*model.Follow) (*model.Follow, error)
	Delete(*model.Follow) error
	ExistsByProfileIdAndTargetId(profileId, targetId string) bool
}