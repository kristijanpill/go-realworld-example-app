package store

import (
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
	"gorm.io/gorm"
)

type ProfilePostgresStore struct {
	db *gorm.DB
}

func NewProfilePostgresStore(db *gorm.DB) (*ProfilePostgresStore, error) {
	err := db.AutoMigrate(&model.Profile{})
	if err != nil {
		return nil, err
	}

	return &ProfilePostgresStore{
		db: db,
	}, nil
}

func (store *ProfilePostgresStore) Create(profile *model.Profile) (*model.Profile, error) {
	result := store.db.Create(profile)

	return profile, result.Error
}

func (store *ProfilePostgresStore) FindById(id string) (*model.Profile, error) {
	var profile model.Profile
	result := store.db.Where("id = ?", id).First(&profile)

	return &profile, result.Error
}

func (store *ProfilePostgresStore) FindByUsername(username string) (*model.Profile, error) {
	var profile model.Profile
	result := store.db.Where("username = ?", username).First(&profile)

	return &profile, result.Error
}

func (store *ProfilePostgresStore) Update(profile *model.Profile) (*model.Profile, error) {
	result := store.db.Save(profile)

	return profile, result.Error
}
