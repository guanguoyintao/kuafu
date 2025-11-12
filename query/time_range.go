package equery

import (
	"fmt"
	"time"
)

// TimeRange 过滤器结构体，包含开始时间和结束时间
type TimeRange struct {
	Start   *time.Time
	End     *time.Time
	IsStart bool // 是否有有效的开始时间
	IsEnd   bool // 是否有有效的结束时间
}

// NewTimeRange 创建一个新的 TimeRange 实例
// 可以设置仅开始时间、仅结束时间，或开始到结束时间
func NewTimeRange(start, end *time.Time) *TimeRange {
	return &TimeRange{
		Start:   start,
		End:     end,
		IsStart: start != nil,
		IsEnd:   end != nil,
	}
}

// IsInRange 检查给定的时间是否在时间范围内
// 支持单范围(>start 或 <end)以及双范围(start - end)
func (tr *TimeRange) IsInRange(t time.Time) bool {
	if tr.IsStart && tr.IsEnd {
		// 双范围检查 (start - end)
		return t.After(*tr.Start) && t.Before(*tr.End)
	} else if tr.IsStart {
		// 只有开始时间，检查 t > start
		return t.After(*tr.Start)
	} else if tr.IsEnd {
		// 只有结束时间，检查 t < end
		return t.Before(*tr.End)
	}
	return false
}

// Format 返回时间范围的字符串表示，格式为 "Start - End"
func (tr *TimeRange) Format() string {
	if tr.IsStart && tr.IsEnd {
		return fmt.Sprintf("%s - %s", tr.Start.Format(time.RFC3339), tr.End.Format(time.RFC3339))
	} else if tr.IsStart {
		return fmt.Sprintf("> %s", tr.Start.Format(time.RFC3339))
	} else if tr.IsEnd {
		return fmt.Sprintf("< %s", tr.End.Format(time.RFC3339))
	}
	return "Invalid TimeRange"
}

// FilterRecords 根据时间范围过滤记录
// 允许单范围（>start 或 <end）或双范围（start - end）
func (tr *TimeRange) FilterRecords(records []time.Time) []time.Time {
	var filteredRecords []time.Time
	for _, record := range records {
		if tr.IsInRange(record) {
			filteredRecords = append(filteredRecords, record)
		}
	}
	return filteredRecords
}
