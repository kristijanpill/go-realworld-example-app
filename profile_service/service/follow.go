package service

import (
	"context"
	"errors"

	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/store"
)

var ErrProfileAlreadyFollowed = errors.New("profile already followed")

type FollowService struct {
	followStore store.FollowStore
	profileStore store.ProfileStore
}

func NewFollowService(followStore store.FollowStore, profileStore store.ProfileStore) *FollowService {
	return &FollowService{
		followStore: followStore,
		profileStore: profileStore,
	}
}

func (service *FollowService) FollowUserByUsername(ctx context.Context, request *pb.FollowRequest) (*pb.ProfileResponse, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)

	currentUserProfile, err := service.profileStore.FindById(currentUserIdString)
	if err != nil {
		return nil, err
	}

	targetUserProfile, err := service.profileStore.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	if (service.followStore.ExistsByProfileIdAndTargetId(currentUserProfile.ID.String(), targetUserProfile.ID.String())) {
		return nil, ErrProfileAlreadyFollowed
	}
	
	follow := model.NewFollow(currentUserProfile.ID, targetUserProfile.ID)
	_, err = service.followStore.Create(follow)
	if err != nil {
		return nil, err
	}

	return &pb.ProfileResponse{
		Profile: &pb.Profile{
			Username: targetUserProfile.Username,
			Bio: targetUserProfile.Bio,
			Image: targetUserProfile.Image,
			Following: true,
		},
	}, nil
}