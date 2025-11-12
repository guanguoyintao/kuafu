package evalidator

import "reflect"

// requiredHandler 检查值是否为其类型的零值。
type requiredHandler[T any] struct {
	baseHandler[T]
}

// Required 返回一个必填项校验器。
// 它会检查值是否是其类型的零值（例如，"" for string, 0 for int, nil for pointers/slices）。
func Required[T any]() ValidationHandler[T] {
	return &requiredHandler[T]{}
}

func (h *requiredHandler[T]) Validate(field string, value T) *ValidationError {
	// 使用反射来判断是否为零值，更具通用性
	if reflect.ValueOf(&value).Elem().IsZero() {
		return &ValidationError{
			Field: field,
			Type:  ErrRequired,
			Msg:   "is required",
		}
	}
	return h.baseHandler.Validate(field, value)
}
