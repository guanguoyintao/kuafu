package ucenter

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var uid = uint64(5124125125125)
var secret = "s"

func TestGetUidFromToken(t *testing.T) {
	ctx, err := GenContext(context.Background(), uid, secret)
	if err != nil {
		assert.NoError(t, err)
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "uid",
			args:    args{ctx: ctx},
			want:    uid,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUidFromToken(tt.args.ctx)
			spew.Dump(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUidFromToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUidFromToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
