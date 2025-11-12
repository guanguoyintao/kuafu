package evalidator

import (
	"fmt"
	"unicode/utf8"
)

type maxLengthHandler[T any] struct {
	baseHandler[T]
	Max int
}

// MaxLength 返回一个最大长度校验器。
// 它支持 string 和 slice 类型。
func MaxLength[T any](max int) ValidationHandler[T] {
	return &maxLengthHandler[T]{Max: max}
}

func (h *maxLengthHandler[T]) Validate(field string, value T) *ValidationError {
	var length int
	validType := true

	switch val := any(value).(type) {
	case string:
		length = utf8.RuneCountInString(val)
	case []any:
		length = len(val)
	default:
		validType = false
	}

	if validType && length > h.Max {
		return &ValidationError{
			Field: field,
			Type:  ErrMaxLength,
			Msg:   fmt.Sprintf("length must be at most %d", h.Max),
		}
	}
	return h.baseHandler.Validate(field, value)
}
