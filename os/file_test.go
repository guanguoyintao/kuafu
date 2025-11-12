package eos

import (
	"io"
	"os"
	"testing"
)

func TestLineCounter(t *testing.T) {
	type args struct {
		r io.Reader
	}
	fd, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "3 line txt",
			args: args{
				r: fd,
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LineCounter(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("LineCounter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LineCounter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
