package entgo

import (
	"entgo.io/ent/dialect/sql"

	paging "github.com/go-fox/go-utils/pagination"
)

// BuildPaginationSelector 构建分页查询条件
func BuildPaginationSelector(pagination bool, page uint32, size uint32) func(selector *sql.Selector) {
	if !pagination {
		return nil
	}
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return func(selector *sql.Selector) {
		selector.Offset(paging.GetPageOffset(page, size)).Limit(int(size))
	}
}
