package econtext

import (
	"context"
	"time"
)

type ValueContext struct {
	context.Context
}

func (c *ValueContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *ValueContext) Done() <-chan struct{} {

	return nil
}

func (c *ValueContext) Err() error {
	return nil
}

func NewValueContext(ctx context.Context) context.Context {
	return &ValueContext{ctx}
}

func NewTimeoutContext(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	valueCxt := &ValueContext{ctx}
	newCtx, cancel := context.WithTimeout(valueCxt, timeout)

	return newCtx, cancel
}
