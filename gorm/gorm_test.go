package egorm

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertDo(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}

	tests := []testCase[any]{
		// 测试 string 类型
		{name: "String NULL", args: args[any]{value: NULLString}, want: ""},
		{name: "String Non-NULL", args: args[any]{value: "hello"}, want: "hello"},

		// 测试 int 类型
		{name: "Int NULL", args: args[any]{value: NULLInt}, want: int(0)},
		{name: "Int Non-NULL", args: args[any]{value: 42}, want: 42},

		// 测试 int8 类型
		{name: "Int8 NULL", args: args[any]{value: int8(NULLInt)}, want: int8(0)},
		{name: "Int8 Non-NULL", args: args[any]{value: int8(8)}, want: int8(8)},

		// 测试 uint 类型
		{name: "Uint NULL", args: args[any]{value: NULLUint}, want: uint(0)},
		{name: "Uint Non-NULL", args: args[any]{value: uint(10)}, want: uint(10)},

		// 测试 uint16 类型
		{name: "Uint16 NULL", args: args[any]{value: uint16(NULLUint)}, want: uint16(0)},
		{name: "Uint16 Non-NULL", args: args[any]{value: uint16(16)}, want: uint16(16)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertDo2Bo(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertDo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertPtrDo(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}

	tests := []testCase[any]{
		// 测试 string 类型
		{name: "String NULL", args: args[any]{value: NULLString}, want: nil},
		{name: "String Non-NULL", args: args[any]{value: "hello"}, want: "hello"},

		// 测试 int 类型
		{name: "Int NULL", args: args[any]{value: NULLInt}, want: nil},
		{name: "Int Non-NULL", args: args[any]{value: 42}, want: 42},

		// 测试 int8 类型
		{name: "Int8 NULL", args: args[any]{value: int8(NULLInt)}, want: nil},
		{name: "Int8 Non-NULL", args: args[any]{value: int8(8)}, want: int8(8)},

		// 测试 uint 类型
		{name: "Uint NULL", args: args[any]{value: NULLUint}, want: nil},
		{name: "Uint Non-NULL", args: args[any]{value: uint(10)}, want: uint(10)},

		// 测试 uint16 类型
		{name: "Uint16 NULL", args: args[any]{value: uint16(NULLUint)}, want: nil},
		{name: "Uint16 Non-NULL", args: args[any]{value: uint16(16)}, want: uint16(16)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertDo2BoPtr(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				if tt.want == nil {
					assert.Nil(t, got, "ConvertPtrDo() = %v, want %v", got, tt.want)
				} else {
					assert.NotNil(t, got)
					assert.Equal(t, *got, tt.want)
				}
			}
		})
	}
}

func TestConvertBo(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[any]{
		// 字符串类型测试用例
		{
			name: "String_NULL",
			args: args[any]{value: ""},
			want: "-1",
		},
		{
			name: "String_Not_NULL",
			args: args[any]{value: "hello"},
			want: "hello",
		},

		// 整型类型测试用例
		{
			name: "Int_NULL",
			args: args[any]{value: int(0)},
			want: int(-1),
		},
		{
			name: "Int_Not_NULL",
			args: args[any]{value: int(42)},
			want: int(42),
		},

		// int8 类型测试用例
		{
			name: "Int8_NULL",
			args: args[any]{value: int8(0)},
			want: int8(-1),
		},
		{
			name: "Int8_Not_NULL",
			args: args[any]{value: int8(42)},
			want: int8(42),
		},

		// int16 类型测试用例
		{
			name: "Int16_NULL",
			args: args[any]{value: int16(0)},
			want: int16(-1),
		},
		{
			name: "Int16_Not_NULL",
			args: args[any]{value: int16(42)},
			want: int16(42),
		},

		// int32 类型测试用例
		{
			name: "Int32_NULL",
			args: args[any]{value: int32(0)},
			want: int32(-1),
		},
		{
			name: "Int32_Not_NULL",
			args: args[any]{value: int32(42)},
			want: int32(42),
		},

		// int64 类型测试用例
		{
			name: "Int64_NULL",
			args: args[any]{value: int64(0)},
			want: int64(-1),
		},
		{
			name: "Int64_Not_NULL",
			args: args[any]{value: int64(42)},
			want: int64(42),
		},

		// uint 类型测试用例
		{
			name: "Uint_NULL",
			args: args[any]{value: uint(0)},
			want: uint(0),
		},
		{
			name: "Uint_Not_NULL",
			args: args[any]{value: uint(42)},
			want: uint(42),
		},

		// uint8 类型测试用例
		{
			name: "Uint8_NULL",
			args: args[any]{value: uint8(0)},
			want: uint8(0),
		},
		{
			name: "Uint8_Not_NULL",
			args: args[any]{value: uint8(42)},
			want: uint8(42),
		},

		// uint16 类型测试用例
		{
			name: "Uint16_NULL",
			args: args[any]{value: uint16(0)},
			want: uint16(0),
		},
		{
			name: "Uint16_Not_NULL",
			args: args[any]{value: uint16(42)},
			want: uint16(42),
		},

		// uint32 类型测试用例
		{
			name: "Uint32_NULL",
			args: args[any]{value: uint32(0)},
			want: uint32(0),
		},
		{
			name: "Uint32_Not_NULL",
			args: args[any]{value: uint32(42)},
			want: uint32(42),
		},

		// uint64 类型测试用例
		{
			name: "Uint64_NULL",
			args: args[any]{value: uint64(0)},
			want: uint64(0),
		},
		{
			name: "Uint64_Not_NULL",
			args: args[any]{value: uint64(42)},
			want: uint64(42),
		},

		// float32 类型测试用例
		{
			name: "Float32_NULL",
			args: args[any]{value: float32(0)},
			want: float32(-1.0),
		},
		{
			name: "Float32_Not_NULL",
			args: args[any]{value: float32(42.0)},
			want: float32(42.0),
		},

		// float64 类型测试用例
		{
			name: "Float64_NULL",
			args: args[any]{value: float64(0)},
			want: float64(-1),
		},
		{
			name: "Float64_Not_NULL",
			args: args[any]{value: float64(42.0)},
			want: float64(42.0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ConvertBo2Do(tt.args.value), "ConvertDo(%v)", tt.args.value)
		})
	}
}

func TestConvertBoPtr2Do(t *testing.T) {
	testCases := []struct {
		name     string
		value    any
		expected any
	}{
		// string 类型指针
		{"string: nil pointer", (*string)(nil), NULLString},
		{"string: empty pointer", func() *string { s := ""; return &s }(), ""},
		{"string: non-empty pointer", func() *string { s := "hello"; return &s }(), "hello"},

		// int 类型指针
		{"int: nil pointer", (*int)(nil), NULLInt},
		{"int: pointer with value 0", func() *int { i := 0; return &i }(), 0},
		{"int: pointer with non-zero value", func() *int { i := 42; return &i }(), 42},

		// int8 类型指针
		{"int8: nil pointer", (*int8)(nil), NULLInt8},
		{"int8: pointer with value 0", func() *int8 { i := int8(0); return &i }(), int8(0)},
		{"int8: pointer with non-zero value", func() *int8 { i := int8(8); return &i }(), int8(8)},

		// float32 类型指针
		{"float32: nil pointer", (*float32)(nil), NULLFloat32},
		{"float32: pointer with value 0.0", func() *float32 { f := float32(0.0); return &f }(), float32(0.0)},
		{"float32: pointer with non-zero value", func() *float32 { f := float32(3.14); return &f }(), float32(3.14)},

		// uint 类型指针
		{"uint: nil pointer", (*uint)(nil), NULLUint},
		{"uint: pointer with value 0", func() *uint { u := uint(0); return &u }(), uint(0)},
		{"uint: pointer with non-zero value", func() *uint { u := uint(100); return &u }(), uint(100)},

		// int64 类型指针
		{"int64: nil pointer", (*int64)(nil), NULLInt64},
		{"int64: pointer with value 0", func() *int64 { i := int64(0); return &i }(), int64(0)},
		{"int64: pointer with non-zero value", func() *int64 { i := int64(12345); return &i }(), int64(12345)},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch v := tc.value.(type) {
			case *string:
				got := ConvertBoPtr2Do(v)
				assert.Equal(t, tc.expected, got)
			case *int:
				got := ConvertBoPtr2Do(v)
				assert.Equal(t, tc.expected, got)
			case *int8:
				got := ConvertBoPtr2Do(v)
				assert.Equal(t, tc.expected, got)
			case *float32:
				got := ConvertBoPtr2Do(v)
				assert.Equal(t, tc.expected, got)
			case *uint:
				got := ConvertBoPtr2Do(v)
				assert.Equal(t, tc.expected, got)
			case *int64:
				got := ConvertBoPtr2Do(v)
				assert.Equal(t, tc.expected, got)
			default:
				t.Fatalf("Unsupported type %T", v)
			}
		})
	}
}
