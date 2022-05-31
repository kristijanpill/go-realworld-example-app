package server

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kristijanpill/go-realworld-example-app/common/interceptor"
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
	privateKey, publicKey := server.initKeyPair()
	jwtManager := service.NewJWTManager(privateKey, publicKey)
	profileServiceClient := server.initProfileServiceClient()
	userService := service.NewUserService(userStore, jwtManager, profileServiceClient)
	userHandler := handler.NewUserHandler(userService)
	authInterceptor := interceptor.NewAuthInterceptor("Token", server.config.RestrictedPaths, publicKey)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary()))
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

func (server *Server) initProfileServiceClient() pb.ProfileServiceClient {
	address := fmt.Sprintf("%s:%s", server.config.ProfileServiceHost, server.config.ProfileServicePort)
	profileServiceClient, err := service.NewProfileServiceClient(address)
	if err != nil {
		log.Fatal("cannot create profile service client: ", err)
	}

	return profileServiceClient
}

func (server *Server) initKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(server.config.PrivateKey))
	if err != nil {
		log.Fatal("cannot parse private key: ", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
	if err != nil {
		log.Fatal("cannot parse public key: ", err)
	}

	return privateKey, publicKey
}