package service

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"github.com/khafidprayoga/synapsis-service/repository"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type synapsisService struct {
	log  *zap.Logger
	repo repository.SynapsisRepository
	synapsisv1.UnimplementedSynapsisServiceServer
}

func (s synapsisService) Ping(_ context.Context, _ *emptypb.Empty) (*synapsisv1.PingResponse, error) {
	return &synapsisv1.PingResponse{Message: "pong"}, nil
}

func NewSynapsisService(
	l *zap.Logger,
	repo repository.SynapsisRepository) synapsisv1.SynapsisServiceServer {
	return &synapsisService{
		log:  l,
		repo: repo,
	}
}
