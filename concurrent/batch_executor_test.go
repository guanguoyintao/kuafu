package econcurrent

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
)

func TestBatchExecutor_Run(t *testing.T) {
	contextKey := "k"
	contextValue := "v"

	type fields struct {
		HasTimeout   bool
		callables    []*Callable[any]
		results      map[string]any
		successCount uint32
		errorCount   uint32
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "all success",
			fields: fields{
				HasTimeout: false,
				callables: []*Callable[any]{
					{
						TaskName: "task1",
						Task: func(ctx context.Context, l *log.Helper) (any, error) {
							time.Sleep(1 * time.Second)
							return 1, nil
						},
					},
					{
						TaskName: "task2",
						Task: func(ctx context.Context, l *log.Helper) (any, error) {
							time.Sleep(1 * time.Second)
							return 2, nil
						},
					},
					{
						TaskName: "task3",
						Task: func(ctx context.Context, l *log.Helper) (any, error) {
							time.Sleep(1 * time.Second)
							return 3, nil
						},
					},
				},
				results: map[string]any{
					"task1": 1,
					"task2": 2,
					"task3": 3,
				},
				successCount: 3,
				errorCount:   0,
			},
			args: args{timeout: 2 * time.Second},
		},
		{
			name: "one timeout",
			fields: fields{
				HasTimeout: true,
				callables: []*Callable[any]{
					{
						TaskName: "task1",
						Task: func(ctx context.Context, l *log.Helper) (any, error) {
							time.Sleep(900 * time.Millisecond)
							return 1, nil
						},
					},
					{
						TaskName: "task2",
						Task: func(ctx context.Context, l *log.Helper) (any, error) {
							time.Sleep(2 * time.Second)
							return 2, nil
						},
					},
					{
						TaskName: "task3",
						Task: func(ctx context.Context, l *log.Helper) (any, error) {
							time.Sleep(900 * time.Millisecond)
							return 3, nil
						},
					},
				},
				results: map[string]any{
					"task1": 1,
					"task3": 3,
				},
				successCount: 2,
				errorCount:   0,
			},
			args: args{timeout: 1 * time.Second},
		},
		{
			name: "all timeout",
			fields: fields{
				HasTimeout: true,
				callables: []*Callable[any]{
					{
						TaskName: "task1",
						Task: func(ctx context.Context, l *log.Helper) (any, error) {
							time.Sleep(900 * time.Millisecond)
							return 1, nil
						},
					},
					{
						TaskName: "task2",
						Task: func(ctx context.Context, l *log.Helper) (any, error) {
							time.Sleep(2 * time.Second)
							return 2, nil
						},
					},
					{
						TaskName: "task3",
						Task: func(ctx context.Context, l *log.Helper) (any, error) {
							time.Sleep(900 * time.Millisecond)
							return 3, nil
						},
					},
				},
				results:      map[string]any{},
				successCount: 0,
				errorCount:   0,
			},
			args: args{timeout: 100 * time.Millisecond},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBatchExecutor[any](log.NewHelper(log.DefaultLogger))
			for _, callback := range tt.fields.callables {
				b.AddCallable(callback)
			}
			ctx := context.WithValue(context.Background(), contextKey, contextValue)
			b.Run(ctx, tt.args.timeout)
			resultMap := make(map[string]any)
			for _, result := range b.taskResults {
				resultMap[result.TaskName] = result.Result
			}
			for taskName, taskResult := range resultMap {
				assert.Equal(t, tt.fields.results[taskName], taskResult)
			}
			fmt.Println(b.taskResults, b.SuccessCount, b.ErrorCount)
			assert.Equal(t, tt.fields.HasTimeout, b.HasTimeout)
		})
	}
}
