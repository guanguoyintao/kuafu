package ucenter

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashUserID(t *testing.T) {
	type args struct {
		UserID uint64
		ctx    context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "6",
			args: args{
				UserID: 1,
				ctx:    context.Background(),
			},
			want:    "aBZ41jn1DYer",
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				UserID: 2,
				ctx:    context.Background(),
			},
			want:    "OWKZAGlXxjrB",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashUserID(tt.args.ctx)
			if err != nil {
				assert.NoError(t, err)
			}
			userHashID, err := got.Encode(tt.args.UserID)
			if err != nil {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, userHashID)
		})
	}
}
