package evalidator

import (
	"fmt"
	"regexp"
)

type regexHandler[T any] struct {
	baseHandler[T]
	Pattern *regexp.Regexp
}

// Regex 返回一个正则表达式校验器。
// 该校验器只对 string 类型的值生效。
func Regex[T any](pattern *regexp.Regexp) ValidationHandler[T] {
	return &regexHandler[T]{Pattern: pattern}
}

func (h *regexHandler[T]) Validate(field string, value T) *ValidationError {
	strValue, ok := any(value).(string)
	if !ok {
		// 如果不是字符串类型，则跳过此校验
		return h.baseHandler.Validate(field, value)
	}

	if !h.Pattern.MatchString(strValue) {
		return &ValidationError{
			Field: field,
			Type:  ErrInvalidFormat,
			Msg:   fmt.Sprintf("does not match regex pattern: %s", h.Pattern.String()),
		}
	}
	return h.baseHandler.Validate(field, value)
}
