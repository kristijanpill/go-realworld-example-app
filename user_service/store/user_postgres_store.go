package store

import (
	"github.com/kristijanpill/go-realworld-example-app/user_service/model"
	"gorm.io/gorm"
)

type UserPostgresStore struct {
	db *gorm.DB
}

func NewUserPostgresStore(db *gorm.DB) (*UserPostgresStore, error) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return &UserPostgresStore{
		db: db,
	}, nil
}

func (store *UserPostgresStore) Create(user *model.User) (*model.User, error) {
	result := store.db.Create(user)

	return user, result.Error
}

func (store *UserPostgresStore) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := store.db.Where("email = ?", email).First(&user)

	return &user, result.Error
}

func (store *UserPostgresStore) FindById(id string) (*model.User, error) {
	var user model.User
	result := store.db.Where("id = ?", id).First(&user)

	return &user, result.Error
}

func (store *UserPostgresStore) Update(user *model.User) (*model.User, error) {
	result := store.db.Save(user)

	return user, result.Error
}