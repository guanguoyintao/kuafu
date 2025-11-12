package erand

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestRandomRangeFloat64(t *testing.T) {
	type args struct {
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "0-0.4",
			args: args{
				min: 0,
				max: 0.4,
			},
		},
		{
			name: "0.5-0.6",
			args: args{
				min: 0.5,
				max: 0.6,
			},
		},
		{
			name: "-0.1-0.1",
			args: args{
				min: -0.1,
				max: 0.1,
			},
		},
		{
			name: "-0.8--0.5",
			args: args{
				min: -0.8,
				max: -0.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomRangeFloat64(tt.args.min, tt.args.max)
			fmt.Println(tt.args.min, tt.args.max, got)
		})
	}
}

func TestGenRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{
			name:    "Test length 0",
			args:    args{length: 0},
			wantLen: 0,
		},
		{
			name:    "Test length 10",
			args:    args{length: 10},
			wantLen: 10,
		},
		{
			name:    "Test length 100",
			args:    args{length: 100},
			wantLen: 100,
		},
		{
			name:    "Test length -1 (invalid)",
			args:    args{length: -1},
			wantLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenRandomString(tt.args.length)
			// 打印生成的哈希字符串并检查长度
			fmt.Println("Generated string:", got)
			assert.Equal(t, tt.wantLen, len(got), "Expected length does not match")
		})
	}
}

func TestGenRandomAlphaNumeric(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{
			name:   "最小长度2",
			length: 2,
		},
		{
			name:   "常规长度6",
			length: 6,
		},
		{
			name:   "较长长度12",
			length: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 生成多次测试随机性
			for i := 0; i < 100; i++ {
				result := GenRandomAlphaNumeric(tt.length)
				// 检查长度
				assert.Equal(t, tt.length, len(result))
				// 检查是否同时包含字母和数字
				hasLetter := false
				hasDigit := false
				for _, char := range result {
					if unicode.IsLetter(char) {
						hasLetter = true
					}
					if unicode.IsDigit(char) {
						hasDigit = true
					}
				}
				assert.True(t, hasLetter, "生成的代码必须包含至少一个字母")
				assert.True(t, hasDigit, "生成的代码必须包含至少一个数字")
			}
		})
	}
}
