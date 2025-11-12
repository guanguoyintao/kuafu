package eisoduration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *Duration
		wantErr  error
	}{
		{
			name:  "完整格式",
			input: "P1Y2M3W4DT5H6M7S",
			expected: &Duration{
				Years:   1,
				Months:  2,
				Weeks:   3,
				Days:    4,
				Hours:   5,
				Minutes: 6,
				Seconds: 7,
			},
			wantErr: nil,
		},
		{
			name:  "只有日期部分",
			input: "P1Y2M3W4D",
			expected: &Duration{
				Years:  1,
				Months: 2,
				Weeks:  3,
				Days:   4,
			},
			wantErr: nil,
		},
		{
			name:  "只有时间部分",
			input: "PT5H6M7S",
			expected: &Duration{
				Hours:   5,
				Minutes: 6,
				Seconds: 7,
			},
			wantErr: nil,
		},
		{
			name:  "零持续时间",
			input: "P0D",
			expected: &Duration{
				Days: 0,
			},
			wantErr: nil,
		},
		{
			name:  "零持续时间",
			input: "P",
			expected: &Duration{
				Days: 0,
			},
			wantErr: nil,
		},
		{
			name:     "格式错误 - 缺少P",
			input:    "1Y2M",
			expected: nil,
			wantErr:  ErrBadFormat,
		},
		{
			name:     "格式错误 - 负数",
			input:    "P-1Y",
			expected: nil,
			wantErr:  ErrBadFormat,
		},
		{
			name:     "格式错误 - 无效字符",
			input:    "P1X",
			expected: nil,
			wantErr:  ErrBadFormat,
		},
		{
			name:     "格式错误 - T后无时间",
			input:    "P1YT",
			expected: nil,
			wantErr:  ErrBadFormat,
		},
		{
			name:     "格式错误 - 空字符串",
			input:    "",
			expected: nil,
			wantErr:  ErrBadFormat,
		},
		{
			name:  "所有单位为零",
			input: "P0Y0M0W0DT0H0M0S",
			expected: &Duration{
				Years:   0,
				Months:  0,
				Weeks:   0,
				Days:    0,
				Hours:   0,
				Minutes: 0,
				Seconds: 0,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromString(tt.input)
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}

func TestDuration_ToDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration *Duration
		validate func(t *testing.T, d time.Duration)
	}{
		{
			name: "基本单位测试",
			duration: &Duration{
				Days:    1,
				Hours:   2,
				Minutes: 3,
				Seconds: 4,
			},
			validate: func(t *testing.T, d time.Duration) {
				expected := 26*time.Hour + 3*time.Minute + 4*time.Second
				assert.Equal(t, expected, d)
			},
		},
		{
			name: "一个月测试",
			duration: &Duration{
				Months: 1,
			},
			validate: func(t *testing.T, d time.Duration) {
				now := time.Now()
				nextMonth := now.AddDate(0, 1, 0)
				expected := nextMonth.Sub(now)
				assert.Equal(t, expected, d)
			},
		},
		{
			name: "一年测试",
			duration: &Duration{
				Years: 1,
			},
			validate: func(t *testing.T, d time.Duration) {
				now := time.Now()
				nextYear := now.AddDate(1, 0, 0)
				expected := nextYear.Sub(now)
				assert.Equal(t, expected, d)
			},
		},
		{
			name: "复合时间测试",
			duration: &Duration{
				Years:   1,
				Months:  2,
				Days:    3,
				Hours:   4,
				Minutes: 5,
				Seconds: 6,
			},
			validate: func(t *testing.T, d time.Duration) {
				now := time.Now()
				future := now.AddDate(1, 2, 0).
					AddDate(0, 0, 3).
					Add(4*time.Hour + 5*time.Minute + 6*time.Second)
				expected := future.Sub(now)
				assert.Equal(t, expected, d)
			},
		},
		{
			name: "跨多年月份测试",
			duration: &Duration{
				Years:  2,
				Months: 15, // 1年3个月
			},
			validate: func(t *testing.T, d time.Duration) {
				now := time.Now()
				future := now.AddDate(3, 3, 0) // 2年+15个月 = 3年3个月
				expected := future.Sub(now)
				assert.Equal(t, expected, d)
			},
		},
		{
			name: "周测试",
			duration: &Duration{
				Weeks: 2,
			},
			validate: func(t *testing.T, d time.Duration) {
				expected := 14 * 24 * time.Hour
				assert.Equal(t, expected, d)
			},
		},
		{
			name:     "零持续时间",
			duration: &Duration{},
			validate: func(t *testing.T, d time.Duration) {
				assert.Equal(t, time.Duration(0), d)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.duration.ToDuration()
			tt.validate(t, got)
		})
	}
}

func TestDuration_String(t *testing.T) {
	tests := []struct {
		name     string
		duration *Duration
		expected string
	}{
		{
			name: "完整格式",
			duration: &Duration{
				Years:   1,
				Months:  2,
				Weeks:   3,
				Days:    4,
				Hours:   5,
				Minutes: 6,
				Seconds: 7,
			},
			expected: "P1Y2M3W4DT5H6M7S",
		},
		{
			name: "只有日期部分",
			duration: &Duration{
				Years:  1,
				Months: 2,
				Weeks:  3,
				Days:   4,
			},
			expected: "P1Y2M3W4D",
		},
		{
			name: "只有时间部分",
			duration: &Duration{
				Hours:   5,
				Minutes: 6,
				Seconds: 7,
			},
			expected: "PT5H6M7S",
		},
		{
			name:     "零持续时间",
			duration: &Duration{},
			expected: "P",
		},
		{
			name: "只有年",
			duration: &Duration{
				Years: 1,
			},
			expected: "P1Y",
		},

		{
			name: "只有天",
			duration: &Duration{
				Days: 30,
			},
			expected: "P30D",
		},
		{
			name: "只有月",
			duration: &Duration{
				Months: 1,
			},
			expected: "P1M",
		},
		{
			name: "只有周",
			duration: &Duration{
				Weeks: 1,
			},
			expected: "P1W",
		},
		{
			name: "只有秒",
			duration: &Duration{
				Seconds: 1,
			},
			expected: "PT1S",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.duration.String()
			assert.Equal(t, tt.expected, got)
		})
	}
}
