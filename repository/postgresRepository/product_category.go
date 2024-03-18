package postgresRepository

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
)

func (p postgresRepository) GetProductCategoryById(
	_ context.Context,
	categoryId ...string,
) (cat []*synapsisv1.ProductCategory, e error) {
	result := p.orm.
		Where("id IN (?)", categoryId).
		Find(&cat)

	return cat, unwrapError(result)
}
