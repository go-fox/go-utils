package entgo

import (
	"entgo.io/ent/dialect/sql"
	"github.com/go-fox/go-utils/pagination"
)

// BuildQuerySelect 构建查询条件
func BuildQuerySelect(req *pagination.PagingRequest, defaultOrderField string) (whereSelector func(s *sql.Selector), querySelector []func(s *sql.Selector), err error) {
	defer func() {
		if rec := recover(); rec != nil {
			recErr, ok := rec.(error)
			if ok {
				err = recErr
			}
		}
	}()

	if req.Where != nil {
		whereSelector = QueryCommandToWhereConditions(req.Where.LogicalOperator, req.Where.Conditions)
	}

	// 构建排序条件
	var orderSelector func(s *sql.Selector)
	err, orderSelector = BuildOrderSelector(req.OrderBy, defaultOrderField)
	if err != nil {
		return nil, nil, err
	}
	if whereSelector != nil {
		querySelector = append(querySelector, whereSelector)
	}
	// 添加排序条件
	querySelector = append(querySelector, orderSelector)

	return
}
