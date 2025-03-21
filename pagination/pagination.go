package pagination

import (
	"context"
	"encoding/json"

	"entgo.io/ent/dialect/sql"
)

const (
	// DefaultPage 默认页码
	DefaultPage uint32 = 1
	// DefaultSize 默认每页条数
	DefaultSize uint32 = 10
)

// PagingRequest 分页请求
type PagingRequest struct {
	Page       uint32     `json:"page"`       // 当前页码
	Size       uint32     `json:"size"`       // 每页条数
	Condition  *Condition `json:"condition"`  // 查询条件
	OrderBy    []string   `json:"order_by"`   // 排序字段
	Pagination bool       `json:"pagination"` // 是否分页
	Fields     []string   `json:"fields"`     // 查询字段
}

// ToJSON 转换为json
func (p *PagingRequest) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

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
