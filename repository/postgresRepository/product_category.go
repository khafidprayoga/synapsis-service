package postgresRepository

import (
	"context"
	"fmt"
	commonHelper "github.com/khafidprayoga/synapsis-service/common/helper"
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

func (p postgresRepository) CreateProductCategory(
	ctx context.Context,
	data []*synapsisv1.ProductCategory,
) ([]*synapsisv1.ProductCategory, error) {
	createdProductCategory := p.
		orm.
		WithContext(ctx).
		Create(data)

	if createdProductCategory.Error != nil {
		return data, unwrapError(createdProductCategory)
	}

	return data, nil
}

func (p postgresRepository) GetProductCategories(
	ctx context.Context,
	paging *synapsisv1.Pagination,
) (int64, []*synapsisv1.ProductCategory, error) {
	categories := []*synapsisv1.ProductCategory{}
	q := p.orm.
		WithContext(ctx).Table("product_category").Debug()

	if paging.GetSearch() != "" {
		search := "%" + paging.GetSearch() + "%"
		q = q.Where("name ILIKE ? or description ILIKE ?", search, search)
	}

	if paging.GetStartAt() != nil && paging.GetEndAt() != nil {
		q = q.Where("created_at BETWEEN ? AND ?", paging.GetStartAt().AsTime(), paging.GetEndAt().AsTime())
	}

	if paging.GetOrderBy() != "" {
		q = q.Order(
			fmt.Sprintf(
				"%s %s",
				paging.GetOrderBy(),
				commonHelper.PaginationSortSQL[paging.GetSort()]),
		)
	} else {
		q = q.Order("name ASC")
	}

	var categoriesCount int64
	count := q.Count(&categoriesCount)
	if count.Error != nil {
		return 0, categories, unwrapError(count)
	}

	findAll := q.
		Limit(int(paging.GetPage().GetLimit())).
		Offset(int(paging.GetPage().GetOffset())).
		Find(&categories)
	if findAll.Error != nil {
		return 0, categories, unwrapError(findAll)
	}

	return categoriesCount, categories, nil
}
