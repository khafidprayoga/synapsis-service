package seed

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
)

func NewProductCategory() []*synapsisv1.ProductCategory {
	return []*synapsisv1.ProductCategory{
		{
			Name:        "Electronics",
			Description: faker.Sentence(),
		},
		{
			Name:        "Fashion",
			Description: faker.Sentence(),
		},
		{
			Id:          uuid.New().String(),
			Name:        "Apple",
			Description: faker.Sentence(),
		},
	}
}
