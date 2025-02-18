package pagination

// LogicalOperator 逻辑运算符
type LogicalOperator string

const (
	// LogicalOperatorAnd and查询
	LogicalOperatorAnd LogicalOperator = "and"
	// LogicalOperatorOr or查询
	LogicalOperatorOr LogicalOperator = "or"
)

// QueryOperator 查询运算符
type QueryOperator string

const (
	// QueryOperatorEqual 等于查询
	QueryOperatorEqual QueryOperator = "="
	// QueryOperatorNotEqual 不等于查询
	QueryOperatorNotEqual QueryOperator = "!="
	// QueryOperatorGreater 大于查询
	QueryOperatorGreater QueryOperator = ">"
	// QueryOperatorGreaterEqual 大于等于查询
	QueryOperatorGreaterEqual QueryOperator = ">="
	// QueryOperatorLess 小于查询
	QueryOperatorLess QueryOperator = "<"
	// QueryOperatorLessEqual 小于等于查询
	QueryOperatorLessEqual QueryOperator = "<="
	// QueryOperatorIn in查询
	QueryOperatorIn QueryOperator = "in"
	// QueryOperatorNotIn not in查询
	QueryOperatorNotIn QueryOperator = "not_in"
	// QueryOperatorLike like查询
	QueryOperatorLike QueryOperator = "like"
	// QueryOperatorIsNull null查询
	QueryOperatorIsNull QueryOperator = "null"
	// QueryOperatorIsNotNull not null查询
	QueryOperatorIsNotNull QueryOperator = "not_null"
)

type Where struct {
	LogicalOperator LogicalOperator `json:"logical_operator"`
	Conditions      []Condition     `json:"conditions"`
}

type Condition struct {
	Field           string          `json:"field"`
	Operator        QueryOperator   `json:"operator"`
	Value           interface{}     `json:"value"`
	LogicalOperator LogicalOperator `json:"logical_operator"`
	Conditions      []Condition     `json:"conditions"`
}
