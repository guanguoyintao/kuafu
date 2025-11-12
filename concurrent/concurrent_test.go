package econcurrent

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
)

func TestGoExecHandlerWithoutTimeout(t *testing.T) {
	valueContext := context.WithValue(context.Background(), "key", "value")
	timeoutCtx, cancelFunc := context.WithTimeout(valueContext, time.Second)
	wg := sync.WaitGroup{}
	type args struct {
		ctx        context.Context
		cancelFunc func()
		l          *log.Helper
		runnable   *Runnable
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				ctx:        timeoutCtx,
				l:          log.NewHelper(log.DefaultLogger),
				cancelFunc: cancelFunc,
				runnable: &Runnable{
					Task: func(ctx context.Context, l *log.Helper) error {
						defer wg.Done()
						fmt.Printf("There is a %v in the context\n", ctx.Value("key"))
						select {
						case <-time.After(10 * time.Second): // 模拟耗时操作（2秒）
							fmt.Println("Work has been running for 10 seconds now")
						case <-ctx.Done(): // 如果超时或被取消，退出
							fmt.Println("Work canceled due to:", ctx.Err())
						}
						return nil
					},
					TaskName: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg.Add(1)
			GoExecHandlerWithoutTimeout(tt.args.ctx, tt.args.l, tt.args.runnable)
			cancelFunc()
			wg.Wait()
		})
	}
}

func TestGoExecHandlerWithTimeout(t *testing.T) {
	valueContext := context.WithValue(context.Background(), "key", "value")
	timeoutCtx, cancelFunc := context.WithTimeout(valueContext, 10*time.Second)
	wg := sync.WaitGroup{}
	type args struct {
		timeout    time.Duration
		cancelFunc func()
		l          *log.Helper
		runnable   *Runnable
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				l:          log.NewHelper(log.DefaultLogger),
				cancelFunc: cancelFunc,
				timeout:    time.Hour,
				runnable: &Runnable{
					Task: func(ctx context.Context, l *log.Helper) error {
						defer wg.Done()
						fmt.Printf("There is a %v in the context\n", ctx.Value("key"))
						select {
						case <-time.After(20 * time.Second): // 模拟耗时操作（2秒）
							fmt.Println("Work has been running for 20 seconds now")
						case <-ctx.Done(): // 如果超时或被取消，退出
							fmt.Println("Work canceled due to:", ctx.Err())
						}
						return nil
					},
					TaskName: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg.Add(1)
			go GoExecHandlerWithTimeout(timeoutCtx, tt.args.l, tt.args.runnable, tt.args.timeout)
			go cancelFunc()
			wg.Wait()
		})
	}
}

func TestGoDelayedExecHandler(t *testing.T) {
	// 准备测试数据
	ctx := context.Background()
	logger := log.NewHelper(log.DefaultLogger)
	executed := false
	delay := 100 * time.Millisecond

	// 创建测试任务
	task := &Runnable{
		TaskName: "test_delayed_task",
		Task: func(ctx context.Context, l *log.Helper) error {
			executed = true
			return nil
		},
	}

	// 执行延迟任务
	err := GoDelayedExecHandler(ctx, logger, task, delay)
	assert.NoError(t, err)

	// 等待任务执行完成（需要等待比延迟时间更长一点）
	time.Sleep(delay + 50*time.Millisecond)

	// 验证任务是否执行
	assert.True(t, executed, "任务应该被执行")

	// 测试错误情况
	errorTask := &Runnable{
		TaskName: "test_error_task",
		Task: func(ctx context.Context, l *log.Helper) error {
			return fmt.Errorf("测试错误")
		},
	}

	err = GoDelayedExecHandler(ctx, logger, errorTask, delay)
	assert.NoError(t, err)                  // 初始返回应该是nil
	time.Sleep(delay + 50*time.Millisecond) // 等待错误任务执行
}
