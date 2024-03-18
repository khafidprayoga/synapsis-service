package postgresRepository

import (
	"context"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"gorm.io/gorm"
)

func (p postgresRepository) CreateProduct(
	_ context.Context,
	productData *synapsisv1.Product,
) (*synapsisv1.CreateProductResponse, error) {
	errDoTX := p.orm.Transaction(func(tx *gorm.DB) error {
		if errInsertProduct := tx.Create(&productData); errInsertProduct != nil {
			return errInsertProduct.Error
		}

		var productCatRel []*synapsisv1.ProductCategoryRelation
		for _, category := range productData.GetProductCategories() {
			productCatRel = append(productCatRel, &synapsisv1.ProductCategoryRelation{
				ProductId:         productData.GetId(),
				ProductCategoryId: category.GetId(),
			})
		}
		if errInsertRelation := tx.Create(&productCatRel); errInsertRelation != nil {
			return errInsertRelation.Error
		}
		return nil
	})

	if errDoTX != nil {
		return nil, errDoTX
	}

	return &synapsisv1.CreateProductResponse{
		Product: productData,
	}, nil
}
