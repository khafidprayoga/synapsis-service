package postgresRepository

import (
	"context"
	"github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
)

func (p postgresRepository) GetProductRelations(
	ctx context.Context,
	productId string) ([]*synapsisv1.ProductCategoryRelation, error) {

	relations := []*synapsisv1.ProductCategoryRelation{}
	findRelations := p.orm.
		WithContext(ctx).
		Where("product_id = ?", productId).
		Find(&relations).Debug()
	if findRelations.Error != nil {
		return nil, unwrapError(findRelations)
	}

	return relations, nil
}
