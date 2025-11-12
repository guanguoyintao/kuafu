package eua

import (
	"fmt"
	"testing"
)

func TestHeaderGenerator(t *testing.T) {
	type args struct {
		headers []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				headers: []string{USER_AGENT},
			},
		},
		{
			name: "2",
			args: args{
				headers: []string{USER_AGENT},
			},
		},
		{
			name: "3",
			args: args{
				headers: []string{USER_AGENT},
			},
		},
		{
			name: "4",
			args: args{
				headers: []string{USER_AGENT},
			},
		},
		{
			name: "5",
			args: args{
				headers: []string{USER_AGENT},
			},
		},
		{
			name: "6",
			args: args{
				headers: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HeaderGenerator(tt.args.headers)
			fmt.Printf("testing--------result-----%s\n", tt.name)
			fmt.Println(got)
		})
	}
}
