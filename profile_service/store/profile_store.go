package store

import (
	"github.com/google/uuid"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
)

type ProfileStore interface {
	Save(*model.Profile) (*model.Profile, error)
	FindById(uuid.UUID) (*model.Profile, error)
}