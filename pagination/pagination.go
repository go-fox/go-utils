package pagination

// PagingRequest 分页请求
type PagingRequest struct {
	Page     int      `json:"page"`      // 当前页码
	Size     int      `json:"size"`      // 每页条数
	Where    *Where   `json:"where"`     // 查询条件
	OrderBy  []string `json:"order_by"`  // 排序字段
	NoPaging bool     `json:"no_paging"` // 是否分页
	Fields   []string `json:"fields"`    // 查询字段
}

// PagingResponse 分页响应
type PagingResponse struct {
	Total     int64 `json:"total"`      // 总条数
	TotalPage int64 `json:"total_page"` // 总页数
	Page      int64 `json:"page"`       // 当前页码
	Size      int64 `json:"size"`       // 每页条数
	Records   any   `json:"records"`    // 数据
}
