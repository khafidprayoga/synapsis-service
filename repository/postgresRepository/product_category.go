package postgresRepository

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
)

func (p postgresRepository) GetProductCategoryById(
	ctx context.Context,
	categoryId ...string,
) (cat []*synapsisv1.ProductCategory, e error) {
	result := p.orm.
		WithContext(ctx).
		Where("id IN (?)", categoryId).
		Find(&cat)

	return cat, unwrapError(result)
}
