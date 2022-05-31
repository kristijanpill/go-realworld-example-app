package store

import "github.com/kristijanpill/go-realworld-example-app/user_service/model"

type UserStore interface {
	Save(*model.User) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}