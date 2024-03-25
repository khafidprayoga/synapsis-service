package service

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
	"github.com/khafidprayoga/synapsis-service/src/repository"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type synapsisService struct {
	log         *zap.Logger
	userRepo    repository.UserRepository
	productRepo repository.ProductRepository
	authRepo    repository.AuthRepository
	synapsisv1.UnimplementedSynapsisServiceServer
}

func (s synapsisService) Ping(_ context.Context, _ *emptypb.Empty) (*synapsisv1.PingResponse, error) {
	return &synapsisv1.PingResponse{Message: "pong"}, nil
}

type ServiceRepository struct {
	User    repository.UserRepository
	Product repository.ProductRepository
	Auth    repository.AuthRepository
}

func NewSynapsisService(
	l *zap.Logger,
	serviceRepo ServiceRepository,
) synapsisv1.SynapsisServiceServer {
	return &synapsisService{
		log:         l,
		userRepo:    serviceRepo.User,
		productRepo: serviceRepo.Product,
		authRepo:    serviceRepo.Auth,
	}
}
