package commonHelper

import (
	"errors"
	"github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
	"github.com/samber/lo"
)

var PaginationSortSQL = map[synapsisv1.Pagination_SortBy]string{
	synapsisv1.Pagination_ASC:  "ASC",
	synapsisv1.Pagination_DESC: "DESC",
}

func ValidatePaginationRequest(req *synapsisv1.Pagination, allowedSort ...string) error {
	var (
		start = req.GetStartAt().AsTime()
		end   = req.GetEndAt().AsTime()
	)

	if req.GetStartAt() != nil && req.GetEndAt() == nil {
		return errors.New("filter endAt must not be empty")
	}

	if req.GetEndAt() != nil && req.GetStartAt() == nil {
		return errors.New("filter startAt must not be empty")
	}

	if req.GetStartAt() != nil && req.GetEndAt() != nil {
		if start.After(end) {
			return errors.New("filter startAt must be before endAt")
		}
	}

	if req.GetPage() == nil && req.GetCursor() == nil {
		return errors.New("pagination page (sql) or cursor (nosql) must not be empty")
	}

	if req.GetPage() != nil {
		if req.GetPage().GetLimit() == 0 || req.GetPage().GetLimit() > 20 {
			return errors.New("pagination limit must not be zero and not more than 20")
		}
	}

	if req.GetOrderBy() != "" {
		var isValid bool
		lo.ForEach(allowedSort, func(in string, index int) {
			if in == req.GetOrderBy() {
				isValid = true
				req.OrderBy = &in
				return
			}
		})

		if !isValid {
			return errors.New("invalid order by, check the field on `allowedOrderBy` details")
		}

	}
	return nil
}
