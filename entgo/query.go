package entgo

import (
	"entgo.io/ent/dialect/sql"

	"github.com/go-fox/go-utils/pagination"
)

// BuildQuerySelect 构建查询条件
func BuildQuerySelect(req *pagination.PagingRequest, defaultOrderField string) (whereSelector []func(s *sql.Selector), querySelector []func(s *sql.Selector), err error) {
	defer func() {
		if rec := recover(); rec != nil {
			recErr, ok := rec.(error)
			if ok {
				err = recErr
			}
		}
	}()

	if req.Condition != nil {
		conditions := QueryCommandToWhereConditions(req.Condition.LogicalOperator, req.Condition.Conditions)
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
	pagingSelector := BuildPaginationSelector(req.Pagination, req.Page, req.Size)
	if pagingSelector != nil {
		querySelector = append(querySelector, pagingSelector)
	}
	return
}
