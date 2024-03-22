package mongoRepository

import (
	"errors"
	"github.com/khafidprayoga/synapsis-service/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type mongoRepository struct {
	log   *zap.Logger
	mongo *mongo.Database
}

func NewRepository(l *zap.Logger, m *mongo.Database) (repository.UserRepository, error) {
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
