package services

import (
	"github.com/Hot-One/firstms_go_api_gateway/config"
	"github.com/Hot-One/firstms_go_api_gateway/genproto/user_service"

	// "github.com/golang/protobuf/protoc-gen-go/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	UserService() user_service.UserServiceClient
}

type grpcClients struct {
	userService user_service.UserServiceClient
}

func NewgrpcClients(cfg config.Config) (ServiceManagerI, error) {
	// Order Microservice
	connOrderService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService: user_service.NewUserServiceClient(connOrderService),
	}, nil
}

func (g *grpcClients) UserService() user_service.UserServiceClient {
	return g.userService
}
