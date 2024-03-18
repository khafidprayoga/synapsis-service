package postgresRepository

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
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
