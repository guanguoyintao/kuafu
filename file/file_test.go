package efile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFileExtension(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want struct {
			fileName string
			ext      string
		}
	}{
		{
			name: "minio with ext",
			args: args{
				fileName: "minio/dev/ebadf8e6f4bd447e874c0d9b12c6603c.jpeg",
			},
			want: struct {
				fileName string
				ext      string
			}{
				fileName: "minio/dev/ebadf8e6f4bd447e874c0d9b12c6603c",
				ext:      "jpeg",
			},
		},
		{
			name: "minio with sort ext",
			args: args{
				fileName: "minio/dev/ebadf8e6f4bd447e874c0d9b12c6603c.j",
			},
			want: struct {
				fileName string
				ext      string
			}{
				fileName: "minio/dev/ebadf8e6f4bd447e874c0d9b12c6603c",
				ext:      "j",
			},
		},
		{
			name: "minio without ext",
			args: args{
				fileName: "minio/dev/ebadf8e6f4bd447e874c0d9b12c6603c",
			},
			want: struct {
				fileName string
				ext      string
			}{
				fileName: "minio/dev/ebadf8e6f4bd447e874c0d9b12c6603c",
				ext:      "",
			},
		},

		{
			name: "url with ext",
			args: args{
				fileName: "http://192.168.0.39:9000/dev/1099997_829085_9c14e4bb4c3a2d0d9ec7b977e6f6be74_1732339374.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=NlHxoQThXOJ6BYtz6rRR%2F20241202%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20241202T060644Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=243441ec961a16c32266b590b4967e1e666d71ce0aa887df9f2049e421a5ef03",
			},
			want: struct {
				fileName string
				ext      string
			}{
				fileName: "minio/dev/ebadf8e6f4bd447e874c0d9b12c6603c",
				ext:      "png",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileName, ext := RemoveFileExtension(tt.args.fileName)
			assert.Equal(t, tt.want.fileName, fileName)
			assert.Equal(t, tt.want.ext, ext)
		})
	}
}

func TestRemoveFileExtension1(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "a.txt",
			args: args{
				fileName: "a.txt",
			},
			want:  "a",
			want1: "txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RemoveFileExtension(tt.args.fileName)
			assert.Equalf(t, tt.want, got, "RemoveFileExtension(%v)", tt.args.fileName)
			assert.Equalf(t, tt.want1, got1, "RemoveFileExtension(%v)", tt.args.fileName)
		})
	}
}
