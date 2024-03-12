package service

import (
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"github.com/khafidprayoga/synapsis-service/repository"
	"go.uber.org/zap"
)

type synapsisService struct {
	log  *zap.Logger
	repo repository.SynapsisRepository
	synapsisv1.UnimplementedSynapsisServiceServer
}

func NewSynapsisService(l *zap.Logger, repo repository.SynapsisRepository) synapsisv1.SynapsisServiceServer {
	return &synapsisService{
		log:  l,
		repo: repo,
	}
}
