package redisRepository

import (
	"errors"
	"github.com/khafidprayoga/synapsis-service/repository"
	"go.uber.org/zap"
)

type redisRepository struct {
	log *zap.Logger
}

func NewRepository(l *zap.Logger) (repository.AuthRepository, error) {
	if l == nil {
		return nil, errors.New("logger is required")
	}

	return &redisRepository{
		log: l,
	}, nil
}
