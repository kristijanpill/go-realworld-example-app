package service

import (
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProfileServiceClient(address string) (pb.ProfileServiceClient, error) {
	con, err := getConnection(address)
	if err != nil {
		return nil, err
	}

	return pb.NewProfileServiceClient(con), nil
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
