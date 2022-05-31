package server

import (
	"fmt"
	"log"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/config"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/handler"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/service"
	"github.com/kristijanpill/go-realworld-example-app/profile_service/store"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
	mux    *runtime.ServeMux
}

func NewServer(config *config.Config) *Server {
	server := &Server{
		config: config,
		mux: runtime.NewServeMux(),
	}

	return server;
}

func (server *Server) Start() {
	profileStore := server.initProfileStore()
	profileService := service.NewProfileService(profileStore)
	profileHandler := handler.NewProfileHandler(profileService)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProfileServiceServer(grpcServer, profileHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}

func (server *Server) initProfileStore() store.ProfileStore {
	profileStore, err := store.NewProfilePostgresStore(server.config.ProfileDatabaseHost, server.config.ProfileDatabasePort, server.config.ProfileDatabaseName, server.config.ProfileDatabaseUser, server.config.ProfileDatabasePassword)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	return profileStore
}
