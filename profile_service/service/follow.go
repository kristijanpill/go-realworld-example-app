package service

import (
	"context"
	"errors"

	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/model"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/store"
)

var ErrUserAlreadyFollowed = errors.New("user already followed")
var ErrUserNotFollowed = errors.New("user not followed")
var ErrSelfFollow = errors.New("cannot follow self")

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

func (service *FollowService) ExistsByProfileIdAndTargetId(profileId, targetId string) bool {
	return service.followStore.ExistsByProfileIdAndTargetId(profileId, targetId)
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

	if (currentUserIdString == targetUserProfile.ID.String()) {
		return nil, ErrSelfFollow
	}

	if (service.followStore.ExistsByProfileIdAndTargetId(currentUserProfile.ID.String(), targetUserProfile.ID.String())) {
		return nil, ErrUserAlreadyFollowed
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

func (service *FollowService) UnfollowUserByUsername(ctx context.Context, request *pb.UnfollowRequest) (*pb.ProfileResponse, error) {
	currentUserIdString := ctx.Value(interceptor.CurrentUserKey{}).(string)

	targetUserProfile, err := service.profileStore.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	if (!service.followStore.ExistsByProfileIdAndTargetId(currentUserIdString, targetUserProfile.ID.String())) {
		return nil, ErrUserNotFollowed
	}

	err = service.followStore.DeleteByProfileIdAndTargetId(currentUserIdString, targetUserProfile.ID.String())
	if err != nil {
		return nil, err
	}

	return &pb.ProfileResponse{
		Profile: &pb.Profile{
			Username: targetUserProfile.Username,
			Bio: targetUserProfile.Bio,
			Image: targetUserProfile.Image,
			Following: false,
		},
	}, nil
}