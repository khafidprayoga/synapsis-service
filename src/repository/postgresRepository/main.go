package postgresRepository

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/khafidprayoga/synapsis-service/src/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type postgresRepository struct {
	log *zap.Logger
	orm *gorm.DB
}

func NewRepository(l *zap.Logger, o *gorm.DB) (repository.ProductRepository, error) {
	if l == nil {
		return nil, errors.New("logger is required")
	}

	if o == nil {
		return nil, errors.New("postgres database is required")
	}

	return &postgresRepository{
		log: l,
		orm: o,
	}, nil
}

func unwrapError(result *gorm.DB) error {
	if result.Error != nil {
		var sqlErr *pgconn.PgError
		_ = errors.As(result.Error, &sqlErr)
		if sqlErr != nil {
			replaced := strings.Replace(sqlErr.Message, `"`, "`", -1)
			return errors.New(replaced)
		}

		return result.Error
	}
	return nil
}
