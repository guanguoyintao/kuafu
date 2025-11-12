package evideo

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/guanguoyintao/kuafu/kratos-x/kxlogging"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func TestGenThumb(t *testing.T) {
	kxlogging.InitGlobalLoggerHelper(log.DefaultLogger)
	type args struct {
		ctx      context.Context
		videoURL string
		second   float64
		scale    int
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
		want1   int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "short video",
			args: args{
				ctx:      context.Background(),
				videoURL: "http://192.168.0.39:9000/dev/SaveTik.co_7447065335099952418-hd_1957692_b96a96b52e1b32da6d148f8fcd2a5a52_1733991118.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=NlHxoQThXOJ6BYtz6rRR%2F20241216%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20241216T031813Z&X-Amz-Expires=604800&X-Amz-SignedHeaders=host&X-Amz-Signature=62f3c3e2963371efb68179619a0704a7406c75d5a89ad7f8ff1a9d09bd862cda",
				second:   0,
				scale:    300,
			},
			want:    nil,
			want1:   0,
			wantErr: nil,
		},
		{
			name: "mov video",
			args: args{
				ctx:      context.Background(),
				videoURL: "http://192.168.0.39:9000/dev/IMG_0078_722605_eb3db36fe7817139b8e0037eed14943a_1734320540.MOV?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=NlHxoQThXOJ6BYtz6rRR%2F20241216%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20241216T062926Z&X-Amz-Expires=604800&X-Amz-SignedHeaders=host&X-Amz-Signature=5291a958679bb02930595ae3eede3a1e0312ec6924c3e11db986b65c897d86f3",
				second:   0,
				scale:    300,
			},
			want:    nil,
			want1:   0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GenThumb(tt.args.ctx, tt.args.videoURL, tt.args.second, tt.args.scale)
			if !tt.wantErr(t, err, fmt.Sprintf("GenThumb(%v, %v, %v, %v)", tt.args.ctx, tt.args.videoURL, tt.args.second, tt.args.scale)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GenThumb(%v, %v, %v, %v)", tt.args.ctx, tt.args.videoURL, tt.args.second, tt.args.scale)
			assert.Equalf(t, tt.want1, got1, "GenThumb(%v, %v, %v, %v)", tt.args.ctx, tt.args.videoURL, tt.args.second, tt.args.scale)
		})
	}
}

func TestGenPreview(t *testing.T) {
	type args struct {
		ctx             context.Context
		videoURL        string
		previewDuration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    io.ReadCloser
		want1   int64
		want2   func()
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			args: args{
				ctx:             context.Background(),
				videoURL:        "http://192.168.0.39:9000/dev/IMG_0078_722605_eb3db36fe7817139b8e0037eed14943a_1734320540.MOV?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=NlHxoQThXOJ6BYtz6rRR%2F20241216%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20241216T062926Z&X-Amz-Expires=604800&X-Amz-SignedHeaders=host&X-Amz-Signature=5291a958679bb02930595ae3eede3a1e0312ec6924c3e11db986b65c897d86f3",
				previewDuration: 0,
			},
			want:    nil,
			want1:   0,
			want2:   nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := GenPreview(tt.args.ctx, tt.args.videoURL, tt.args.previewDuration)
			if !tt.wantErr(t, err, fmt.Sprintf("GenPreview(%v, %v, %v)", tt.args.ctx, tt.args.videoURL, tt.args.previewDuration)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GenPreview(%v, %v, %v)", tt.args.ctx, tt.args.videoURL, tt.args.previewDuration)
			assert.Equalf(t, tt.want1, got1, "GenPreview(%v, %v, %v)", tt.args.ctx, tt.args.videoURL, tt.args.previewDuration)
			assert.Equalf(t, tt.want2, got2, "GenPreview(%v, %v, %v)", tt.args.ctx, tt.args.videoURL, tt.args.previewDuration)
		})
	}
}
