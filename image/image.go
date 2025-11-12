package eimage

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"

	ebuffer "github.com/guanguoyintao/kuafu/buffer"
	"github.com/guanguoyintao/kuafu/kratos-x/kxlogging"
)

func GenThumb(ctx context.Context, imageURI string, scale int) (io.ReadCloser, int, error) {
	width, height, err := ExtractDimensions(ctx, imageURI)
	if err != nil {
		return nil, 0, err
	}
	// 如果原始图片大小比缩略图小就返回原图
	if width <= scale && height <= scale {
		// todo: 低成本获取原图
		return nil, 0, nil
	}
	// 按长边缩放，短边padding
	// scale=width:height
	var vf string
	if width > height {
		vf = fmt.Sprintf("scale=%d:-1:flags=bicubic", scale)
	} else {
		vf = fmt.Sprintf("scale=-1:%d:flags=bicubic", scale)
	}
	args := []string{
		"-i", imageURI,
		"-vf", vf,
		"-frames:v", "1",
		"-f", "image2pipe",
		"-vcodec", "mjpeg", //  输出 mjpeg 格式
		"-q:v", "1", // 设置图像质量 (可选)
	}
	args = append(args, "-")
	cmd := exec.Command("ffmpeg", args...)
	kxlogging.GetGlobalLogger().WithContext(ctx).Debugf("ffmpeg gen thumb cmd: %s", cmd.String())
	buf := ebuffer.NewPoolBuffer(5 << 30)
	stderr := ebuffer.NewPoolBuffer(5 << 20)
	cmd.Stdout = buf
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		return nil, 0, fmt.Errorf("ffmpeg command failed: %w, stderr: %s", err, stderr.String()) //包含stderr信息
	}
	return buf, buf.Len(), nil
}

func ExtractDimensions(ctx context.Context, videoURL string) (int, int, error) {
	cmd := exec.CommandContext(ctx, "ffprobe",
		"-v", "error",
		"-select_streams", "v:0",
		"-show_entries", "stream=width,height",
		"-of", "csv=p=0",
		videoURL)
	buf := ebuffer.NewPoolBuffer(5 << 20)
	stderr := ebuffer.NewPoolBuffer(5 << 20)
	defer buf.Close()
	defer stderr.Close()
	cmd.Stdout = buf
	cmd.Stderr = stderr
	kxlogging.GetGlobalLogger().WithContext(ctx).Debugf("ffmpeg extract dimensions cmd: %s", cmd.String())
	err := cmd.Run()
	output := strings.TrimSpace(buf.String())
	parts := strings.Split(output, ",")
	if len(parts) < 2 {
		return 0, 0, fmt.Errorf("unexpected ffprobe output: %s", output)
	}
	// 转换为整数
	var width, height int
	_, err = fmt.Sscanf(parts[0], "%d", &width)
	if err != nil {
		return 0, 0, err
	}
	_, err = fmt.Sscanf(parts[1], "%d", &height)
	if err != nil {
		return 0, 0, err
	}
	return width, height, nil
}
