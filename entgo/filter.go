package entgo

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/go-fox/fox/api/gen/go/pagination"
	"strings"
)

const (
	// JSONFieldDelimiter JSONB字段分隔符
	JSONFieldDelimiter = "." // JSONB字段分隔符
)

// splitJsonFieldKey 分割JSON字段键
func splitJsonFieldKey(key string) []string {
	return strings.Split(key, JSONFieldDelimiter)
}

// isJSONFieldKey 是否为JSON字段键
func isJSONFieldKey(key string) bool {
	return strings.Contains(key, JSONFieldDelimiter)
}

// QueryCommandToWhereConditions 将查询条件转换为sql.Selector
func QueryCommandToWhereConditions(logicalOperator pagination.Condition_LogicalOperator, conditions []*pagination.Condition) func(s *sql.Selector) {
	if len(conditions) == 0 {
		return nil
	}
	return func(s *sql.Selector) {
		predicate := processQueryCondition(s, logicalOperator, conditions)
		s.Where(predicate)
	}
}

// processQueryCondition 处理查询映射
func processQueryCondition(s *sql.Selector, logicOperator pagination.Condition_LogicalOperator, conditions []*pagination.Condition) *sql.Predicate {
	predicates := makeConditions(s, conditions)
	if logicOperator == pagination.Condition_OR {
		return sql.Or(predicates...)
	}
	return sql.And(predicates...)
}

func makeConditions(s *sql.Selector, conditions []*pagination.Condition) []*sql.Predicate {
	var ps []*sql.Predicate
	for _, condition := range conditions {
		p := makeFieldFilter(s, condition)
		if p != nil {
			ps = append(ps, p)
		}
		if len(condition.Conditions) > 0 {
			predicates := processQueryCondition(s, condition.LogicalOperator, condition.Conditions)
			if predicates != nil {
				ps = append(ps, predicates)
			}
		}
	}
	return ps
}

// makeFieldFilter 构建查询条件
func makeFieldFilter(s *sql.Selector, condition *pagination.Condition) *sql.Predicate {
	p := sql.P()
	field := condition.Field
	isJSONField := isJSONFieldKey(field)
	if isJSONField {
		splitField := splitJsonFieldKey(field)
		if len(splitField) == 2 {
			field = filterJSONField(s, splitField[0], splitField[1])
		}
	}
	return processQueryOperator(s, p, condition.QueryOperator, field, condition.Value)
}

// filterJSONField 处理JSON字段
func filterJSONField(s *sql.Selector, field, jsonbField string) string {
	p := sql.P()
	switch s.Builder.Dialect() {
	case dialect.Postgres:
		p.Ident(s.C(field)).WriteString(" ->> ").WriteString("'" + jsonbField + "'")
		break

	case dialect.MySQL:
		str := fmt.Sprintf("JSON_EXTRACT(%s, '$.%s')", s.C(field), jsonbField)
		p.WriteString(str)
		break
	}
	return p.String()
}

func processQueryOperator(selector *sql.Selector, p *sql.Predicate, op pagination.Condition_QueryOperator, field string, value any) *sql.Predicate {
	switch op {
	case pagination.Condition_EQUAL:
		return p.EQ(selector.C(field), value)
	case pagination.Condition_NOT_EQUAL:
		return p.NEQ(selector.C(field), value)
	case pagination.Condition_GREATER:
		return p.GT(selector.C(field), value)
	case pagination.Condition_GREATER_EQUAL:
		return p.GTE(selector.C(field), value)
	case pagination.Condition_LESS:
		return p.LT(selector.C(field), value)
	case pagination.Condition_LESS_EQUAL:
		return p.LTE(selector.C(field), value)
	case pagination.Condition_IN:
		args, ok := value.([]any)
		if !ok {
			return nil
		}
		return p.In(selector.C(field), args...)
	case pagination.Condition_NOT_IN:
		args, ok := value.([]any)
		if !ok {
			return nil
		}
		return p.NotIn(selector.C(field), args)
	case pagination.Condition_LIKE:
		s, ok := value.(string)
		if !ok {
			s = convertor.ToString(value)
		}
		return p.Like(selector.C(field), "%"+s+"%")
	case pagination.Condition_IS_NULL:
		p.IsNull(selector.C(field))
	case pagination.Condition_IS_NOT_NULL:
		p.NotNull(selector.C(field))
	default:
		return nil
	}
	return nil
}
