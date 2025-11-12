package econcurrent

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	econtext "github.com/guanguoyintao/kuafu/context"
	"github.com/guanguoyintao/kuafu/kratos-x/kxmiddleware/tracing"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// Runnable 没有返回值的标准可执行单元
type Runnable struct {
	cancel       context.CancelFunc
	Task         func(ctx context.Context, l *log.Helper) error
	SubRunnables []*Runnable
	TaskName     string
	OnError      func(ctx context.Context, l *log.Helper, err error)
}

// Cancel 提供取消任务的功能
func (r *Runnable) Cancel() {
	if r.cancel != nil {
		// 调用取消函数，通知任务取消
		r.cancel()
	}
}

// Callable 带返回值和错误的标准可执行单元
type Callable[T any] struct {
	Task     func(ctx context.Context, l *log.Helper) (T, error)
	TaskName string
}

func GoroutineRecover(ctx context.Context, l *log.Helper) {
	r := recover()
	if r != nil {
		pc, file, line, ok := runtime.Caller(2)
		if !ok {
			return
		}
		funcName := runtime.FuncForPC(pc).Name()
		l.WithContext(ctx).Error(fmt.Sprintf("goroutine err is: %+v", r),
			"file", file,
			"line", line,
			"funcName", funcName)
	}
}

func goExecTracerRunnableHandler(ctx context.Context, l *log.Helper, runnable *Runnable) error {
	var span trace.Span
	tracer, ok := tracing.GetTracerFromContext(ctx)
	if ok {
		// 使用父 context 创建新的 Span
		ctx, span = tracer.Start(ctx, runnable.TaskName)
		defer span.End()
		// 设置 Span 的一些属性
		span.SetAttributes(attribute.String("task", runnable.TaskName))
		// 打印当前 Span 的 TraceID 和 SpanID
		l.WithContext(ctx).Debugf("Trace ID: %s, Span ID: %s\n", span.SpanContext().TraceID(), span.SpanContext().SpanID())
	}
	// 执行worker
	err := runnable.Task(ctx, l)
	if err != nil {
		buf := make([]byte, 64<<10) //nolint:mnd
		n := runtime.Stack(buf, false)
		buf = buf[:n]
		// 输出详细的错误信息
		l.WithContext(ctx).Errorw("stack", string(buf), "msg", fmt.Sprintf("task name %v, runnable error: %v", runnable.TaskName, err.Error()))
		return err
	}
	return nil
}

func goExecTracerCallbackHandler[T any](ctx context.Context, l *log.Helper, callback *Callable[T]) (T, error) {
	var span trace.Span
	tracer, ok := tracing.GetTracerFromContext(ctx)
	if ok {
		// 使用父 context 创建新的 Span
		ctx, span = tracer.Start(ctx, callback.TaskName)
		defer span.End()
		// 设置 Span 的一些属性
		span.SetAttributes(attribute.String("task", callback.TaskName))
		// 打印当前 Span 的 TraceID 和 SpanID
		l.WithContext(ctx).Debugf("Trace ID: %s, Span ID: %s\n", span.SpanContext().TraceID(), span.SpanContext().SpanID())
	}
	// 执行worker
	res, err := callback.Task(ctx, l)
	if err != nil {
		var zero T
		l.WithContext(ctx).Errorf("runnable err is %v", err.Error())
		return zero, err
	}

	return res, nil
}

func GoExecHandlerWithoutTimeout(ctx context.Context, l *log.Helper, runnable *Runnable) {
	ctx = econtext.NewValueContext(ctx)
	go func() {

		defer GoroutineRecover(ctx, l)
		_ = goExecTracerRunnableHandler(ctx, l, runnable)
		// 子任务
		if runnable.SubRunnables != nil && len(runnable.SubRunnables) > 0 {
			for _, subtask := range runnable.SubRunnables {
				sub := subtask
				GoExecHandlerWithoutTimeout(ctx, l, sub)
			}
		}
	}()
	return
}

func GoExecHandlerWithTimeout(ctx context.Context, l *log.Helper, runnable *Runnable, timeout time.Duration) {
	// 创建一个带有超时的上下文
	ctx, cancel := econtext.NewTimeoutContext(ctx, timeout)
	runnable.cancel = cancel
	go func() {
		defer cancel() // 确保在函数退出时取消上下文
		// 子任务
		defer GoroutineRecover(ctx, l)
		var exit bool
		for !exit {
			select {
			case <-ctx.Done(): // 检查上下文是否被取消
				l.WithContext(ctx).Errorf("task was canceled due to context cancellation: %s", ctx.Err())
				if runnable.OnError == nil {
					return
				}
				runnable.OnError(ctx, l, ctx.Err())
			default:
				err := goExecTracerRunnableHandler(ctx, l, runnable)
				if err != nil {
					if runnable.OnError == nil {
						return
					}
					runnable.OnError(ctx, l, err)
					return
				}
				exit = true
			}
			break
		}
		if runnable.SubRunnables != nil && len(runnable.SubRunnables) > 0 {
			for _, subtask := range runnable.SubRunnables {
				sub := subtask
				GoExecHandlerWithTimeout(ctx, l, sub, timeout)
			}
		}
	}()
}

// GoDelayedExecHandler Delay task
func GoDelayedExecHandler(ctx context.Context, l *log.Helper, runnable *Runnable, delay time.Duration) error {
	ctx = econtext.NewValueContext(ctx)
	go func() {
		time.After(delay)
		// 使用 AfterFunc 创建另一个定时器
		time.AfterFunc(delay, func() {
			err := goExecTracerRunnableHandler(ctx, l, runnable)
			if err != nil {
				l.WithContext(ctx).Errorf("worker err is %v", err.Error())
			}
		})
	}()
	return nil
}

func GoExecWithSemaphore(ctx context.Context, l *log.Helper, runnable *Runnable, wg *sync.WaitGroup) {
	ctx = econtext.NewValueContext(ctx)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer GoroutineRecover(ctx, l)
		err := goExecTracerRunnableHandler(ctx, l, runnable)
		if err != nil {
			l.WithContext(ctx).Errorf("worker err is %v", err.Error())
		}
	}()
}

func GoExecWithErrSemaphore(ctx context.Context, l *log.Helper, runnable *Runnable, eg *errgroup.Group) {
	ctx = econtext.NewValueContext(ctx)
	eg.Go(func() error {
		defer GoroutineRecover(ctx, l)
		err := goExecTracerRunnableHandler(ctx, l, runnable)
		if err != nil {
			l.WithContext(ctx).Errorf("worker err is %v", err.Error())
			return err
		}

		return nil
	})
}

func RetryWithBackoff(ctx context.Context, l *log.Helper, task func(ctx context.Context, l *log.Helper) error, maxRetries int, maxDelay time.Duration) error {
	ctx = econtext.NewValueContext(ctx)
	var retryDelay = time.Second
	var retries int
	for {
		err := task(ctx, l)
		if err == nil {
			return nil
		}
		l.WithContext(ctx).Error(err.Error())
		retries++
		if retries > maxRetries {
			return err
		}
		// Calculate the next retry delay using exponential backoff.
		retryDelay *= 2
		if retryDelay > maxDelay {
			retryDelay = maxDelay
		}
		time.Sleep(retryDelay)
	}
}
