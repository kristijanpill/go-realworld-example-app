package store

import "github.com/kristijanpill/go-realworld-example-app/user_service/model"

type UserStore interface {
	Create(*model.User) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindById(string) (*model.User, error)
	Update(*model.User) (*model.User, error)
}