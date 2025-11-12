package evalidator

// RuleFunc 是一个简单的自定义校验函数类型。
type RuleFunc[T any] func(field string, value T) *ValidationError

type ruleHandler[T any] struct {
	baseHandler[T]
	rule RuleFunc[T]
}

// Rule 将一个自定义的校验函数适配成一个 ValidationHandler。
func Rule[T any](fn RuleFunc[T]) ValidationHandler[T] {
	return &ruleHandler[T]{rule: fn}
}

func (h *ruleHandler[T]) Validate(field string, value T) *ValidationError {
	if err := h.rule(field, value); err != nil {
		return err
	}
	return h.baseHandler.Validate(field, value)
}
