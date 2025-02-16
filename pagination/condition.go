package pagination

// LogicalOperator 逻辑运算符
type LogicalOperator int32

const (
	// LogicalOperatorAnd and查询
	LogicalOperatorAnd LogicalOperator = iota + 1
	// LogicalOperatorOr or查询
	LogicalOperatorOr
)

type QueryOperator int32

const (
	// QueryOperatorEqual 等于
	QueryOperatorEqual QueryOperator = iota + 1
	// QueryOperatorNotEqual 不等于
	QueryOperatorNotEqual
	// QueryOperatorGreater 大于
	QueryOperatorGreater
	// QueryOperatorGreaterEqual 大于等于
	QueryOperatorGreaterEqual
	// QueryOperatorLess 小于
	QueryOperatorLess
	// QueryOperatorLessEqual 小于等于
	QueryOperatorLessEqual
	// QueryOperatorIn 在其中
	QueryOperatorIn
	// QueryOperatorNotIn 不在其中
	QueryOperatorNotIn
	// QueryOperatorLike 模糊查询
	QueryOperatorLike
	// QueryOperatorIsNull 为空
	QueryOperatorIsNull
	// QueryOperatorIsNotNull 不为空
	QueryOperatorIsNotNull
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
