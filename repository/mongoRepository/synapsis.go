package mongoRepository

import (
	"errors"
	"github.com/khafidprayoga/synapsis-service/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type mongoRepository struct {
	log   *zap.Logger
	mongo *mongo.Database
}

func NewRepository(l *zap.Logger, m *mongo.Database) (repository.SynapsisRepository, error) {
	if l == nil {
		return nil, errors.New("logger is required")
	}

	if m == nil {
		return nil, errors.New("mongo database is required")
	}

	return &mongoRepository{
		log:   l,
		mongo: m,
	}, nil
}
