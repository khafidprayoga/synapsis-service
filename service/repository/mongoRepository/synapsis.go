package mongoRepository

import (
	"errors"
	"github.com/khafidprayoga/synapsis-service/service/repository"
	"go.uber.org/zap"
)

type mongoRepository struct {
	Log *zap.Logger
}

func NewRepository(l *zap.Logger) (repository.SynapsisRepository, error) {
	if l == nil {
		return nil, errors.New("logger is required")
	}

	return &mongoRepository{
		Log: l,
	}, nil
}
