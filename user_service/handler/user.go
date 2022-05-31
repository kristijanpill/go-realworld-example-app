package handler

import (
	"context"

	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/user_service/service"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) CreateUser(ctx context.Context, request *pb.NewUserRequest) (*pb.UserResponse, error) {
	return handler.service.Register(request)
}

func (handler *UserHandler) Login(ctx context.Context, request *pb.LoginUserRequest) (*pb.UserResponse, error) {
	return handler.service.Login(request)
}