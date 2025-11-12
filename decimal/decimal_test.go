package edecimal

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"正整数", "123", "123", false},
		{"负整数", "-123", "-123", false},
		{"正小数", "123.456", "123.456", false},
		{"负小数", "-123.456", "-123.456", false},
		{"零", "0", "0", false},
		{"零小数", "0.0", "0.0", false},
		{"前导零", "00123.456", "123.456", false},
		{"多个小数点", "123.456.789", "", true},
		{"空字符串", "", "", true},
		{"非法字符", "12a34", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromString(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got.String())
			}
		})
	}
}

func TestDecimalOperations(t *testing.T) {
	tests := []struct {
		name      string
		a         string
		b         string
		wantAdd   string
		wantSub   string
		wantMul   string
		wantDiv   string
		precision int
	}{
		{
			name:      "正数运算",
			a:         "123.45",
			b:         "67.89",
			wantAdd:   "191.34",
			wantSub:   "55.56",
			wantMul:   "8381.0205",
			wantDiv:   "1.8183826778612461334511710119310649580203",
			precision: 40,
		},
		{
			name:      "负数运算",
			a:         "-123.45",
			b:         "67.89",
			wantAdd:   "-55.56",
			wantSub:   "-191.34",
			wantMul:   "-8381.0205",
			wantDiv:   "-1.8183826778612461334511710119310649580203",
			precision: 40,
		},
		{
			name:      "零相关运算",
			a:         "123.45",
			b:         "0",
			wantAdd:   "123.45",
			wantSub:   "123.45",
			wantMul:   "0.00",
			wantDiv:   "", // 除以零应该报错
			precision: 40,
		},
		{
			name:      "相同数字运算",
			a:         "123.45",
			b:         "123.45",
			wantAdd:   "246.90",
			wantSub:   "0.00",
			wantMul:   "15239.9025",
			wantDiv:   "1.0000000000000000000000000000000000000000",
			precision: 40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err1 := NewFromString(tt.a)
			b, err2 := NewFromString(tt.b)
			assert.NoError(t, err1)
			assert.NoError(t, err2)

			// 测试加法
			got := a.Add(b)
			assert.Equal(t, tt.wantAdd, got.String())

			// 测试减法
			got = a.Sub(b)
			assert.Equal(t, tt.wantSub, got.String())

			// 测试乘法
			got = a.Mul(b)
			assert.Equal(t, tt.wantMul, got.String())

			// 测试除法
			if tt.wantDiv != "" {
				got, err := a.Div(b, tt.precision)
				assert.NoError(t, err)
				assert.Equal(t, tt.wantDiv, got.String())
			} else {
				_, err := a.Div(b, tt.precision)
				assert.Error(t, err)
			}
		})
	}
}

func TestDecimalString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"正常数字", "123.456", "123.456"},
		{"负数", "-123.456", "-123.456"},
		{"整数", "123", "123"},
		{"零", "0", "0"},
		{"小于1的数", "0.123", "0.123"},
		{"末尾有零", "123.4500", "123.4500"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := NewFromString(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, d.String())
		})
	}
}

func TestDecimal_Float64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{
			name:     "整数",
			input:    "123",
			expected: 123.0,
		},
		{
			name:     "带小数",
			input:    "123.456",
			expected: 123.456,
		},
		{
			name:     "负数",
			input:    "-123.456",
			expected: -123.456,
		},
		{
			name:     "零",
			input:    "0",
			expected: 0,
		},
		{
			name:     "小数点后多位",
			input:    "0.123456789012345",
			expected: 0.123456789012345,
		},
		{
			name:     "大数",
			input:    "999999999999.999999999999",
			expected: 999999999999.999999999999,
		},
		{
			name:     "科学计数法范围内的小数",
			input:    "0.000000000000123",
			expected: 0.000000000000123,
		},
		{
			name:     "多个零的小数",
			input:    "1.000000",
			expected: 1.0,
		},
		{
			name:     "特殊值 - 0.1",
			input:    "0.1",
			expected: 0.1,
		},
		{
			name:     "特殊值 - 0.7",
			input:    "0.7",
			expected: 0.7,
		},
		{
			name:     "特殊值 - -0.1",
			input:    "-0.1",
			expected: -0.1,
		},
		{
			name:     "特殊值 - -0.7",
			input:    "-0.7",
			expected: -0.7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := NewFromString(tt.input)
			assert.NoError(t, err)

			result := d.Float64()
			assert.Equal(t, tt.expected, result)

			// 额外验证字符串表示是否一致
			resultStr := strconv.FormatFloat(result, 'f', -1, 64)
			expectedStr := strconv.FormatFloat(tt.expected, 'f', -1, 64)
			assert.Equal(t, expectedStr, resultStr)
		})
	}
}

// 测试空值和 nil 情况
func TestDecimal_Float64_Nil(t *testing.T) {
	var d *Decimal
	assert.Equal(t, float64(0), d.Float64())

	d = &Decimal{}
	assert.Equal(t, float64(0), d.Float64())
}
