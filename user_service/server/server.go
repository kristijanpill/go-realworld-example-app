package server

import (
	"fmt"
	"log"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kristijanpill/go-realworld-example-app/common/proto/pb"
	"github.com/kristijanpill/go-realworld-example-app/user_service/config"
	"github.com/kristijanpill/go-realworld-example-app/user_service/handler"
	"github.com/kristijanpill/go-realworld-example-app/user_service/service"
	"github.com/kristijanpill/go-realworld-example-app/user_service/store"
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
	userStore := server.initUserStore()
	jwtManager := server.initJWTManager()
	profileServiceClient := server.initProfileServiceClient()
	userService := service.NewUserService(userStore, jwtManager, profileServiceClient)
	userHandler := handler.NewUserHandler(userService)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}

func (server *Server) initUserStore() store.UserStore {
	userStore, err := store.NewUserPostgresStore(server.config.UserDatabaseHost, server.config.UserDatabasePort, server.config.UserDatabaseName, server.config.UserDatabaseUser, server.config.UserDatabasePassword)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	return userStore
}

func (server *Server) initJWTManager() *service.JWTManager {
	jwtManager, err := service.NewJWTManager(server.config.PrivateKey, server.config.PublicKey)
	if err != nil {
		log.Fatal("cannot create jwt manager: ", err)
	}

	return jwtManager
}

func (server *Server) initProfileServiceClient() pb.ProfileServiceClient {
	address := fmt.Sprintf("%s:%s", server.config.ProfileServiceHost, server.config.ProfileServicePort)
	profileServiceClient, err := service.NewProfileServiceClient(address)
	if err != nil {
		log.Fatal("cannot create profile service client: ", err)
	}

	return profileServiceClient
}
