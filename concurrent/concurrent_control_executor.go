package econcurrent

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ConcurrentControlExecutor[T any] struct {
	concurrencyControl chan struct{}
}

func NewConcurrentControlExecutor[T any](maxConcurrency uint32) *ConcurrentControlExecutor[T] {
	return &ConcurrentControlExecutor[T]{
		concurrencyControl: make(chan struct{}, maxConcurrency),
	}
}

func (q *ConcurrentControlExecutor[T]) Run(ctx context.Context, l *log.Helper, callback *Callable[T]) (any, error) {
	q.concurrencyControl <- struct{}{}
	defer func() {
		<-q.concurrencyControl
	}()
	res, err := callback.Task(ctx, l)
	if err != nil {
		return nil, err
	}

	return res, nil
}
