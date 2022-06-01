package store

import (
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
)

type ProfileStore interface {
	Create(*model.Profile) (*model.Profile, error)
	FindById(id string) (*model.Profile, error)
	FindByUsername(username string) (*model.Profile, error)
	Update(*model.Profile) (*model.Profile, error)
}