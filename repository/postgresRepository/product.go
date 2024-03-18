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
		errInsertProduct := tx.Create(&productData)
		if errInsertProduct.Error != nil {
			return errInsertProduct.Error
		}

		var productCatRel []*synapsisv1.ProductCategoryRelation
		for _, category := range productData.GetProductCategories() {
			productCatRel = append(productCatRel, &synapsisv1.ProductCategoryRelation{
				ProductId:         productData.GetId(),
				ProductCategoryId: category.GetId(),
			})
		}

		insertRelation := tx.Create(&productCatRel)
		if insertRelation.Error != nil {
			return insertRelation.Error
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
