package entgo

import (
	"entgo.io/ent/dialect/sql"
	"github.com/go-fox/fox/api/gen/go/pagination"
)

// BuildQuerySelect 构建查询条件
func BuildQuerySelect(req *pagination.PagingParams, defaultOrderField string) (whereSelector []func(s *sql.Selector), querySelector []func(s *sql.Selector), err error) {
	defer func() {
		if rec := recover(); rec != nil {
			recErr, ok := rec.(error)
			if ok {
				err = recErr
			}
		}
	}()

	if req.Query != nil {
		conditions := QueryCommandToWhereConditions(req.Query.LogicalOperator, req.Query.Conditions)
		if conditions != nil {
			whereSelector = append(whereSelector, conditions)
		}
	}

	// 构建排序条件
	var orderSelector func(s *sql.Selector)
	orderSelector, err = BuildOrderSelector(req.OrderBy, defaultOrderField)
	if err != nil {
		return nil, nil, err
	}

	// 添加排序条件
	if orderSelector != nil {
		querySelector = append(querySelector, orderSelector)
	}

	// 添加分页条件
	pagingSelector := BuildPaginationSelector(req.GetNoPaging(), req.GetPage(), req.GetSize())
	if pagingSelector != nil {
		querySelector = append(querySelector, pagingSelector)
	}
	return
}
