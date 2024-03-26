package commonHelper

//
//import (
//	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
//	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"strings"
//)
//
//type Filter struct {
//	Pagination   *synapsisv1.Pagination
//	Condition    primitive.A
//	AllowedOrder []string
//}
//
//type PagingQuery struct {
//	After  *string
//	Before *string
//	Limit  *int
//	Order  *string
//}
//
//type CursorTransform struct {
//	After  *string `json:"after,omitempty" query:"after"`
//	Before *string `json:"before,omitempty" query:"before"`
//}
//
//var DefaultPage int = 1
//var DefaultLimit int = 10
//
//func Resolver(q PagingQuery) *paginator.Paginator {
//	p := paginator.New()
//
//	p.SetKeys("ID") // [default: "ID"] (supporting multiple keys, order of keys matters)
//
//	if q.After != nil {
//		p.SetAfterCursor(*q.After) // [default: nil]
//	}
//
//	if q.Before != nil {
//		p.SetBeforeCursor(*q.Before) // [default: nil]
//	}
//
//	if q.Limit != nil {
//		p.SetLimit(*q.Limit) // [default: 10]
//	}
//
//	if q.Order != nil && *q.Order == "asc" {
//		p.SetOrder(paginator.ASC) // [default: paginator.DESC]
//	}
//	return p
//}
//
//func (f *Filter) Statements() primitive.D {
//	p := f.Pagination
//	next := bson.D{}
//	orderField := f.GetOrderField()
//	opCursor := "$gt"
//
//	if p != nil {
//		if p.GetPage().GetOrderBy() != "" {
//			_, order := SplitOrderBy(p.GetPage().GetOrderBy())
//
//			if order < 0 {
//				opCursor = "$lt"
//			}
//		}
//
//		if p.GetCursor() != nil && p.GetCursor().SortId != "" {
//			objectID, _ := primitive.ObjectIDFromHex(p.GetCursor().SortId)
//
//			if orderField == "" || p.Cursor.GetFieldValue() == "" {
//				next = bson.D{
//					{"_id", bson.D{
//						{opCursor, objectID}},
//					},
//				}
//			} else {
//				next = bson.D{
//					{"$or",
//						bson.A{
//							bson.D{
//								{orderField,
//									bson.D{
//										{opCursor, p.Cursor.GetFieldValue()},
//									},
//								},
//							},
//							bson.D{
//								{
//									orderField, p.Cursor.GetFieldValue(),
//								},
//								{
//									"_id", bson.D{
//										{opCursor, objectID}},
//								},
//							},
//						},
//					},
//				}
//			}
//		}
//	}
//
//	return bson.D{
//		{
//			"$and", bson.A{
//				f.Conditions(),
//				next,
//			},
//		},
//	}
//}
//
//func (f *Filter) SetConditions(c primitive.D) primitive.A {
//	f.Condition = append(f.Condition, c)
//	return f.Condition
//}
//
//func (f *Filter) SetConditionIn(c primitive.M) primitive.A {
//	f.Condition = append(f.Condition, c)
//	return f.Condition
//}
//
//func (f *Filter) Conditions() primitive.D {
//	f.Condition = append(f.Condition, bson.D{{
//		"$and", bson.A{
//			bson.D{{"deleted_at", nil}},
//			bson.D{{"dt.deleted_at", nil}},
//		},
//	}})
//
//	return bson.D{
//		{
//			"$and", f.Condition,
//		},
//	}
//}
//
//func (f *Filter) Options() []*options.FindOptions {
//	p := f.Pagination
//	sort := bson.D{}
//	opts := []*options.FindOptions{}
//	limit := int64(DefaultLimit)
//	if p != nil {
//		if p.GetPage().GetOrderBy() != "" {
//			orderby, order := SplitOrderBy(p.GetPage().GetOrderBy())
//			sort = bson.D{{orderby, order}}
//		}
//
//		if p.GetPage().GetLimit() > 0 {
//			limit = p.GetPage().GetLimit()
//		}
//	}
//	opts = append(opts, options.Find().SetLimit(limit))
//	opts = append(opts, options.Find().SetSort(sort))
//	return opts
//}
//
//func (f *Filter) GetOrderField() string {
//	orderField := ""
//	orderby, _ := SplitOrderBy(f.Pagination.GetPage().GetOrderBy())
//	for _, a := range f.AllowedOrder {
//		if a == orderby {
//			orderField = orderby
//		}
//	}
//
//	return orderField
//}
//
//func SplitOrderBy(orderBy string) (string, int) {
//	orderBy = strings.TrimSpace(orderBy)
//	desc := false
//	field := orderBy
//	if field != "" {
//		if strings.HasPrefix(orderBy, "-") {
//			field = strings.TrimPrefix(orderBy, "-")
//			desc = true
//		}
//	}
//
//	if field == "" {
//		field = "_id"
//	}
//
//	sort := 1
//	if desc {
//		sort = -1
//	}
//	return field, sort
//}
