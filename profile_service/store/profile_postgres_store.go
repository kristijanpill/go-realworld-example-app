package store

import (
	"github.com/google/uuid"
	"github.com/kristijanpill/go-realworld-example-app/common/db"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
	"gorm.io/gorm"
)

type ProfilePostgresStore struct {
	db *gorm.DB
}

func NewProfilePostgresStore(host, port, dbname, user, password string) (*ProfilePostgresStore, error) {
	db, err := db.NewPostgresConnection(host, port, dbname, user, password)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Profile{})
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

func (store *ProfilePostgresStore) FindById(id uuid.UUID) (*model.Profile, error) {
	var profile model.Profile
	result := store.db.Where("id = ?", id.String()).First(&profile)

	return &profile, result.Error
}

func (store *ProfilePostgresStore) Update(profile *model.Profile) (*model.Profile, error) {
	result := store.db.Save(profile)

	return profile, result.Error
}