package entgo

import (
	"context"
	"entgo.io/ent/dialect/sql"
)

const (
	// DefaultPage 默认页码
	DefaultPage uint32 = 1
	// DefaultSize 默认每页条数
	DefaultSize uint32 = 10
)

type ModifyBuilder[T any] interface {
	Modify(modifiers ...func(s *sql.Selector)) T
}

// PagingQueryBuilder 查询构建器
type PagingQueryBuilder[T any, V any, M any] interface {
	Count(ctx context.Context) (int, error)
	Limit(limit int) T
	Offset(offset int) T
	All(ctx context.Context) ([]V, error)
	Modify(modifiers ...func(s *sql.Selector)) M
}

// PagingResponse 分页响应
type PagingResponse[T any] struct {
	Total   uint32 `json:"total"`   // 总条数
	Records []T    `json:"records"` // 数据
}

// GetPageOffset 获取分页偏移量
func GetPageOffset(page, size uint32) int {
	return int((page - 1) * size)
}

// BuildPaginationSelector 构建分页查询条件
func BuildPaginationSelector(pagination bool, page uint32, size uint32) func(selector *sql.Selector) {
	if !pagination {
		return nil
	}
	if page < 1 {
		page = DefaultPage
	}
	if size < 1 {
		size = DefaultSize
	}
	return func(selector *sql.Selector) {
		selector.Offset(GetPageOffset(page, size)).Limit(int(size))
	}
}
