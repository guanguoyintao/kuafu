package equery

// Pagination 通用分页结构体
type Pagination struct {
	Page     int `json:"page"`      // 当前页码
	PageSize int `json:"page_size"` // 每页数量
}

// NewPagination 创建分页对象
func NewPagination(page, pageSize int) *Pagination {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // 默认每页10条
	}
	return &Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

// GetOffset 获取数据库查询的偏移量
func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

// GetLimit 获取查询限制数
func (p *Pagination) GetLimit() int {
	return p.PageSize
}
