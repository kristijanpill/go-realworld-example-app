package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kristijanpill/go-realworld-example-app/api_gateway/config"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *config.Config
	mux *runtime.ServeMux
}

func NewServer(config *config.Config) *Server {
	server := &Server{
		config: config,
		mux: runtime.NewServeMux(),
	}

	server.initGatewayHandlers()

	return server;
}

func (server *Server) initGatewayHandlers() {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	userServiceEndpoint := formatEndpoint(server.config.UserServiceHost, server.config.UserServicePort)
	err := pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), server.mux, userServiceEndpoint, dialOptions)
	if err != nil {
		log.Fatal("cannot register user service handler: ", err)
	}

	profileServiceEndpoint := formatEndpoint(server.config.ProfileServiceHost, server.config.ProfileServicePort)
	err = pb.RegisterProfileServiceHandlerFromEndpoint(context.Background(), server.mux, profileServiceEndpoint, dialOptions)
	if err != nil {
		log.Fatal("cannot register profile service handler: ", err)
	}
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}

func formatEndpoint(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}