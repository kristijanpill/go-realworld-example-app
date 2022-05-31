package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/user_service/model"
	"github.com/kristijanpill/go-realworld-example-app/user_service/store"
)

var ErrUserNotActivated = errors.New("user is not activated")

type UserService struct {
	store store.UserStore
	jwtManager *JWTManager
	profileServiceClient pb.ProfileServiceClient
}

func NewUserService(store store.UserStore, jwtManager *JWTManager, profileServiceClient pb.ProfileServiceClient) *UserService {
	return &UserService{
		store: store,
		jwtManager: jwtManager,
		profileServiceClient: profileServiceClient,
	}
}

func (service *UserService) Register(request *pb.NewUserRequest) (*pb.UserResponse, error) {
	user, err := model.NewUser(request.User.Email, request.User.Password)
	if err != nil {
		return nil, err
	}

	createProfileRequest := &pb.CreateProfileRequest{
		Id: user.ID.String(),
		Profile: &pb.ProfileInfo{
			Username: request.User.Username,
		},
	}
	_, err = service.profileServiceClient.CreateProfile(context.Background(), createProfileRequest)
	if err != nil {
		return nil, err
	}

	_, err = service.store.Save(user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		User: &pb.User{
			Email: request.User.Email,
			Username: request.User.Username,
		},
	}, nil
}

func (service *UserService) Login(request *pb.LoginUserRequest) (*pb.UserResponse, error) {
	user, err := service.store.FindByEmail(request.User.Email)
	if err != nil {
		return nil, err
	}

	if (!user.Active) {
		return nil, ErrUserNotActivated
	}

	profile, err := service.profileServiceClient.GetProfileById(context.Background(), &pb.ProfileIdRequest{Id: user.ID.String()})
	if err != nil {
		return nil, err
	}

	token, err := service.jwtManager.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		User: &pb.User{
			Email: user.Email,
			Token: token,
			Username: profile.Profile.Username,
			Bio: profile.Profile.Bio,
			Image: profile.Profile.Image,
		},
	}, nil
}

func (service *UserService) GetCurrentUser(ctx context.Context) (*pb.UserResponse, error) {
	email := ctx.Value(interceptor.CurrentUserKey{}).(string)
	user, err := service.store.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	profile, err := service.getUserProfile(user.ID)
	if err != nil {
		return nil, err
	}

	token := ctx.Value(interceptor.TokenKey{}).(string)

	return &pb.UserResponse{User: &pb.User{
		Email: user.Email,
		Token: token,
		Username: profile.Profile.Username,
		Bio: profile.Profile.Bio,
		Image: profile.Profile.Image,
	}}, nil
}

func (service *UserService) getUserProfile(id uuid.UUID) (*pb.ProfileResponse, error) {
	return service.profileServiceClient.GetProfileById(context.Background(), &pb.ProfileIdRequest{Id: id.String()})
}