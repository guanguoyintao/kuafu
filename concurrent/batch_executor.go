package econcurrent

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	econtext "github.com/guanguoyintao/kuafu/context"
)

type BatchExecutorResult[T any] struct {
	TaskName string
	Result   T
	Err      error
}

type BatchExecutor[T any] struct {
	log *log.Helper
	// 是否超时
	HasTimeout bool
	// 结果集
	taskResultMap map[string]*BatchExecutorResult[T]
	taskResults   []*BatchExecutorResult[T]
	// 操作成功个数
	SuccessCount uint32
	// 操作返回error个数
	ErrorCount uint32

	callables   []*Callable[T]
	resultsChan chan *BatchExecutorResult[T]
	doneChan    chan struct{}
	errChan     chan error
}

func NewBatchExecutor[T any](l *log.Helper) *BatchExecutor[T] {
	return &BatchExecutor[T]{
		log:           l,
		taskResultMap: make(map[string]*BatchExecutorResult[T]),
		taskResults:   make([]*BatchExecutorResult[T], 0),
		doneChan:      make(chan struct{}),
		errChan:       make(chan error),
	}
}

// AddCallable 添加需并行执行的可执行单元
func (b *BatchExecutor[T]) AddCallable(callable *Callable[T]) {
	b.callables = append(b.callables, callable)
}

func (b *BatchExecutor[T]) Run(ctx context.Context, timeout time.Duration) error {
	if b.log == nil {
		b.log = log.NewHelper(log.DefaultLogger)
	}
	// 初始化channels
	b.resultsChan = make(chan *BatchExecutorResult[T], len(b.callables))
	// 设置超时context
	timeoutCtx, cancel := econtext.NewTimeoutContext(ctx, timeout)
	defer cancel()
	// 启动所有任务
	remainingTasks := len(b.callables)
	for i, callable := range b.callables {
		go b.executeTask(timeoutCtx, i, callable)
	}
	// 等待所有任务完成或超时
	for remainingTasks > 0 {
		select {
		case result := <-b.resultsChan:
			b.taskResultMap[result.TaskName] = result
			b.taskResults = append(b.taskResults, result)
			atomic.AddUint32(&b.SuccessCount, 1)
			remainingTasks--
		case <-b.errChan:
			atomic.AddUint32(&b.ErrorCount, 1)
			remainingTasks--
		case <-timeoutCtx.Done():
			b.HasTimeout = true
			return timeoutCtx.Err()
		}
	}
	close(b.doneChan)
	return nil
}

func (b *BatchExecutor[T]) executeTask(ctx context.Context, index int, task *Callable[T]) {
	defer GoroutineRecover(ctx, b.log)
	result, err := goExecTracerCallbackHandler(ctx, b.log, task)
	if err != nil {
		b.log.WithContext(ctx).Error(err.Error())
		b.errChan <- err
	}
	b.resultsChan <- &BatchExecutorResult[T]{
		TaskName: task.TaskName,
		Result:   result,
		Err:      err,
	}
}

// GetResult 获取指定任务名称的执行结果
func (b *BatchExecutor[T]) GetResult(taskName string) (T, error) {
	select {
	case <-b.doneChan:
		if result, exists := b.taskResultMap[taskName]; exists {
			return result.Result, nil
		}
		var zero T
		return zero, fmt.Errorf("task not found: %s", taskName)
	}
}

// GetResultList 获取指定任务名称的执行结果
func (b *BatchExecutor[T]) GetResultList() []*BatchExecutorResult[T] {
	select {
	case <-b.doneChan:
		return b.taskResults
	}
}
