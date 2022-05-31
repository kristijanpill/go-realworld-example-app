package store

import (
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
	"gorm.io/gorm"
)

type FollowPostgresStore struct {
	db *gorm.DB
}

func NewFollowPostgresStore(db *gorm.DB) (*FollowPostgresStore, error) {
	err := db.AutoMigrate(&model.Follow{})	
	if err != nil {
		return nil, err
	}

	return &FollowPostgresStore{
		db: db,
	}, nil
}

func (store *FollowPostgresStore) Create(follow *model.Follow) (*model.Follow, error) {
	result := store.db.Create(follow)

	return follow, result.Error
}

func (store *FollowPostgresStore) Delete(follow *model.Follow) error {
	result := store.db.Delete(follow)

	return result.Error
}

func (store *FollowPostgresStore) ExistsByProfileIdAndTargetId(profileId, targetId string) bool {
	var follow model.Follow
	return store.db.Where("profile_id = ? AND target_id = ?", profileId, targetId).First(&follow).RowsAffected == 1
}