package evalidator

import (
	"fmt"
	"unicode/utf8"
)

type minLengthHandler[T any] struct {
	baseHandler[T]
	Min int
}

// MinLength 返回一个最小长度校验器。
// 它支持 string 和 slice 类型。
func MinLength[T any](min int) ValidationHandler[T] {
	return &minLengthHandler[T]{Min: min}
}

func (h *minLengthHandler[T]) Validate(field string, value T) *ValidationError {
	var length int
	validType := true

	switch val := any(value).(type) {
	case string:
		length = utf8.RuneCountInString(val)
	case []any: // 仅作为示例，可以扩展到 []int, []string 等
		length = len(val)
	default:
		validType = false
	}

	if validType && length < h.Min {
		return &ValidationError{
			Field: field,
			Type:  ErrMinLength,
			Msg:   fmt.Sprintf("length must be at least %d", h.Min),
		}
	}
	return h.baseHandler.Validate(field, value)
}
