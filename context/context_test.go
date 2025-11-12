package econtext

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimeoutContext(t *testing.T) {
	type args struct {
		ctx        context.Context
		cancelFunc func()
		timeout    time.Duration
	}
	valueContext := context.WithValue(context.Background(), "key", "value")
	timeoutCtx, cancelFunc := context.WithTimeout(valueContext, time.Second)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				ctx:        timeoutCtx,
				cancelFunc: cancelFunc,
				timeout:    1 * time.Minute,
			},
		},
	}
	wg := sync.WaitGroup{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := NewTimeoutContext(tt.args.ctx, tt.args.timeout)
			wg.Add(1)
			go func(ctx context.Context) {
				fmt.Printf("There is a %v in the context\n", ctx.Value("key"))
				defer wg.Done()
				select {
				case <-time.After(10 * time.Second): // 模拟耗时操作（2秒）
					fmt.Println("Work has been running for 10 seconds now")
				case <-ctx.Done(): // 如果超时或被取消，退出
					fmt.Println("Work canceled due to:", ctx.Err())
				}
				return
			}(ctx)
			tt.args.cancelFunc()
			wg.Wait()
		})
	}
}

func TestNewTimeoutContext(t *testing.T) {
	type args struct {
		ctx        context.Context
		cancelFunc func()
		timeout    time.Duration
	}
	valueContext := context.WithValue(context.Background(), "key", "value")
	timeoutCtx, cancelFunc := context.WithTimeout(valueContext, time.Second)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				ctx:        timeoutCtx,
				cancelFunc: cancelFunc,
				timeout:    30 * time.Second,
			},
		},
	}
	wg := sync.WaitGroup{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := NewTimeoutContext(tt.args.ctx, tt.args.timeout)
			wg.Add(1)
			go func(ctx context.Context) {
				fmt.Printf("There is a %v in the context\n", ctx.Value("key"))
				defer wg.Done()
				select {
				case <-time.After(10 * time.Second): // 模拟耗时操作（2秒）
					fmt.Println("Work has been running for 10 seconds now")
				case <-ctx.Done(): // 如果超时或被取消，退出
					fmt.Println("Work canceled due to:", ctx.Err())
				}
				return
			}(ctx)
			tt.args.cancelFunc()
			wg.Wait()
		})
	}
}
