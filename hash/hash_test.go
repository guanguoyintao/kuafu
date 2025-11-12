package ehash

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMD532(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sort",
			args: args{
				s: "test",
			},
		},
		{
			name: "long",
			args: args{
				s: "awfbhukabfvliawbgvliabhwngilahwbjksbvjzxbvalwiughiwlqhgnliwqghilwqahbglisaebnghliaehuglaiehgjawnlawhgfkawuhgilfawhglisauhglaiwsghlishfgliawhugfialwughawilhgfbaiuw",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HashMD532(tt.args.s)
			fmt.Println(got)
			assert.Equal(t, 32, len(got))
		})
	}
}

func TestHashMurmurHash340(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sort",
			args: args{
				s: "test",
			},
		},
		{
			name: "long",
			args: args{
				s: "awfbhukabfvliawbgvliabhwngilahwbjksbvjzxbvalwiughiwlqhgnliwqghilwqahbglisaebnghliaehuglaiehgjawnlawhgfkawuhgilfawhglisauhglaiwsghlishfgliawhugfialwughawilhgfbaiuw",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashMurmurHash340(tt.args.s)
			if err != nil {
				t.Error(err)
			}
			fmt.Println(got)
			assert.Equal(t, 40, len(got))
		})
	}
}

func TestCalcFileSHA256(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "ExistingFile",
			args: args{
				filePath: "existing.txt", // 指定一个已存在的文件路径
			},
			want:    "3bace75732dae79185fed42047d71dfb32a4cb90605d87d4cf03e2d14b09f3d7",
			wantErr: assert.NoError,
		},
		{
			name: "EmptyFile",
			args: args{
				filePath: "empty.txt", // 指定一个空文件路径
			},
			want:    "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			wantErr: assert.NoError,
		},
		{
			name: "NonExistingFile",
			args: args{
				filePath: "non_existing.txt", // 指定一个不存在的文件路径
			},
			want:    "",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcFileSHA256(context.Background(), tt.args.filePath)
			if !tt.wantErr(t, err, fmt.Sprintf("CalculateFileSHA256(%v)", tt.args.filePath)) {
				return
			}
			assert.Equalf(t, tt.want, got, "CalculateFileSHA256(%v)", tt.args.filePath)
			if err == nil {
				assert.Equalf(t, 64, len(got), "CalculateFileSHA256(%v)", tt.args.filePath)
			}
		})
	}
}

func TestHashMurmurHash36(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr assert.ErrorAssertionFunc
	}{
		// 普通字符串
		{
			name: "Standard input",
			args: args{
				s: "awghfkigawlfiahf",
			},
			wantLen: 6,
			wantErr: assert.NoError,
		},
		// 空字符串
		{
			name: "Empty string",
			args: args{
				s: "",
			},
			wantLen: 6,
			wantErr: assert.NoError,
		},
		// 单字符字符串
		{
			name: "Single character",
			args: args{
				s: "a",
			},
			wantLen: 6,
			wantErr: assert.NoError,
		},
		// 重复字符
		{
			name: "Repeating characters",
			args: args{
				s: "aaaaaaa",
			},
			wantLen: 6,
			wantErr: assert.NoError,
		},
		// 特殊字符
		{
			name: "Special characters",
			args: args{
				s: "!@#$%^&*()",
			},
			wantLen: 6,
			wantErr: assert.NoError,
		},
		// 数字字符串
		{
			name: "Numeric string",
			args: args{
				s: "1234567890",
			},
			wantLen: 6,
			wantErr: assert.NoError,
		},
		// Unicode 字符串
		{
			name: "Unicode characters",
			args: args{
				s: "你好，世界",
			},
			wantLen: 6,
			wantErr: assert.NoError,
		},
		// 长字符串
		{
			name: "Long string",
			args: args{
				s: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
			},
			wantLen: 6,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashMurmurHash36(tt.args.s)
			// 检查是否有错误
			if tt.wantErr != nil {
				tt.wantErr(t, err)
			}
			// 打印生成的哈希字符串并检查长度
			fmt.Println("Generated hash:", got)
			assert.Equal(t, tt.wantLen, len(got), "Expected length does not match")
		})
	}
}

func TestHashMurmurHash36WithLength(t *testing.T) {
	type args struct {
		s      string
		lenght int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr assert.ErrorAssertionFunc
	}{
		// 普通字符串
		{
			name: "Standard input",
			args: args{
				s:      "awghfkigawlfiahf",
				lenght: 5,
			},
			wantLen: 5,
			wantErr: assert.NoError,
		},
		// 空字符串
		{
			name: "Empty string",
			args: args{
				s:      "",
				lenght: 5,
			},
			wantLen: 5,
			wantErr: assert.NoError,
		},
		// 单字符字符串
		{
			name: "Single character",
			args: args{
				s:      "a",
				lenght: 5,
			},
			wantLen: 5,
			wantErr: assert.NoError,
		},
		// 重复字符
		{
			name: "Repeating characters",
			args: args{
				s:      "aaaaaaa",
				lenght: 5,
			},
			wantLen: 5,
			wantErr: assert.NoError,
		},
		// 特殊字符
		{
			name: "Special characters",
			args: args{
				s:      "!@#$%^&*()",
				lenght: 5,
			},
			wantLen: 5,
			wantErr: assert.NoError,
		},
		// 数字字符串
		{
			name: "Numeric string",
			args: args{
				s:      "1234567890",
				lenght: 5,
			},
			wantLen: 5,
			wantErr: assert.NoError,
		},
		// Unicode 字符串
		{
			name: "Unicode characters",
			args: args{
				s:      "你好，世界",
				lenght: 5,
			},
			wantLen: 5,
			wantErr: assert.NoError,
		},
		// 长字符串
		{
			name: "Long string",
			args: args{
				s:      "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
				lenght: 5,
			},
			wantLen: 5,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashMurmurHash36WithLength(tt.args.s, tt.args.lenght)
			// 检查是否有错误
			if tt.wantErr != nil {
				tt.wantErr(t, err)
			}
			// 打印生成的哈希字符串并检查长度
			fmt.Println("Generated hash:", got)
			assert.Equal(t, tt.wantLen, len(got), "Expected length does not match")
		})
	}
}
