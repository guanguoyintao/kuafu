// package evalidator 提供了基于责任链模式和泛型的统一参数校验器。
package evalidator

import (
	"fmt"
)

// ValidationErrorType 是一个枚举类型，用于表示不同的校验错误类型。
type ValidationErrorType int

const (
	ErrRequired ValidationErrorType = iota
	ErrInvalidFormat
	ErrMinLength
	ErrMaxLength
	ErrCustom // 用于自定义规则
)

// ValidationError 包含了详细的错误信息。
type ValidationError struct {
	Field string
	Type  ValidationErrorType
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("field: %s, type: %d, error: %s", e.Field, e.Type, e.Msg)
}

// ValidationHandler 是所有校验器的通用接口。
type ValidationHandler[T any] interface {
	SetNext(handler ValidationHandler[T]) ValidationHandler[T]
	Validate(field string, value T) *ValidationError
}

// baseHandler 实现了责任链的基本逻辑，可被其他具体校验器嵌入。
type baseHandler[T any] struct {
	next ValidationHandler[T]
}

func (b *baseHandler[T]) SetNext(handler ValidationHandler[T]) ValidationHandler[T] {
	b.next = handler
	return handler
}

// Validate 将校验任务传递给责任链中的下一个节点。
func (b *baseHandler[T]) Validate(field string, value T) *ValidationError {
	if b.next != nil {
		return b.next.Validate(field, value)
	}
	return nil
}

// Validate 是校验器的统一入口函数。
// 它接收字段名、值以及一个或多个校验器，然后在内部构建责任链并执行校验。
func Validate[T any](field string, value T, handlers ...ValidationHandler[T]) *ValidationError {
	if len(handlers) == 0 {
		return nil
	}

	// 动态构建责任链
	for i := 0; i < len(handlers)-1; i++ {
		handlers[i].SetNext(handlers[i+1])
	}

	// 从链的头部开始校验
	return handlers[0].Validate(field, value)
}
