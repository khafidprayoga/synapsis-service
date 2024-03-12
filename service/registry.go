package service

import (
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"github.com/khafidprayoga/synapsis-service/service/repository"
	"go.uber.org/zap"
)

type synapsisService struct {
	Log  *zap.Logger
	Repo repository.SynapsisRepository
	synapsisv1.UnimplementedSynapsisServiceServer
}

func NewSynapsisService(l *zap.Logger, repo repository.SynapsisRepository) synapsisv1.SynapsisServiceServer {
	return &synapsisService{
		Log:  l,
		Repo: repo,
	}
}
