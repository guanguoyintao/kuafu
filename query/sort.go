package equery

type SortOrder string

// Sort 排序字段的定义
type Sort struct {
	Field  string // 排序字段
	IsDesc bool
}

// Sorts 排序规则数组
type Sorts []*Sort
