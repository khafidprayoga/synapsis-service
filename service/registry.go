package service

import (
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"go.uber.org/zap"
)

type synapsisService struct {
	Log *zap.Logger
	synapsisv1.UnimplementedSynapsisServiceServer
}

func NewSynapsisServivce(l *zap.Logger) synapsisv1.SynapsisServiceServer {
	return &synapsisService{
		Log: l,
	}

}
