package entgo

import (
	"context"
	"github.com/go-fox/fox/api/gen/go/pagination"
)

// Pagination 分页查询
func Pagination[T PagingQueryBuilder[T, V, M], V any, M any](ctx context.Context, query T, params *pagination.PagingParams, defaultOrderField ...string) (PagingResponse[V], error) {
	var orderField string
	if len(defaultOrderField) > 0 {
		orderField = defaultOrderField[0]
	}
	whereSelector, querySelector, err := BuildQuerySelect(params, orderField)
	if err != nil {
		return PagingResponse[V]{}, err
	}
	// 查询总条数
	query.Modify(whereSelector...)
	count, err := query.Count(ctx)
	if err != nil {
		return PagingResponse[V]{}, err
	}
	// 添加分页条件
	query.Modify(querySelector...)
	all, err := query.All(ctx)
	if err != nil {
		return PagingResponse[V]{}, err
	}

	return PagingResponse[V]{
		Total:   uint32(count),
		Records: all,
	}, nil
}
