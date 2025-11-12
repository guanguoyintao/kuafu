package evideo

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	ebuffer "github.com/guanguoyintao/kuafu/buffer"

	eimage "github.com/guanguoyintao/kuafu/image"
	"github.com/guanguoyintao/kuafu/kratos-x/kxlogging"
)

func ExtractDuration(ctx context.Context, videoURL string) (time.Duration, error) {
	cmd := exec.Command("ffprobe", "-i", videoURL, "-show_entries", "format=duration", "-v", "quiet", "-of", "csv=p=0")
	buf := ebuffer.NewPoolBuffer(5 << 20)
	stderr := ebuffer.NewPoolBuffer(5 << 20)
	defer buf.Close()
	defer stderr.Close()
	cmd.Stdout = buf
	cmd.Stderr = stderr
	kxlogging.GetGlobalLogger().WithContext(ctx).Debugf("ffmpeg extract duration cmd: %s", cmd.String())
	err := cmd.Run()
	if err != nil {
		return 0, fmt.Errorf("ffprobe command failed: %w", err)
	}
	durationStr := strings.TrimSpace(buf.String())
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration: %w", err)
	}
	return time.Duration(duration * float64(time.Second)), nil
}

func ExtractBitrate(ctx context.Context, videoURL string) (int, error) {
	cmd := exec.CommandContext(ctx, "ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=bit_rate", "-of", "default=noprint_wrappers=1:nokey=1", videoURL)
	buf := ebuffer.NewPoolBuffer(5 << 20)
	stderr := ebuffer.NewPoolBuffer(5 << 20)
	defer buf.Close()
	defer stderr.Close()
	cmd.Stdout = buf
	cmd.Stderr = stderr
	kxlogging.GetGlobalLogger().WithContext(ctx).Debugf("ffmpeg extract bitrate cmd: %s", cmd.String())
	err := cmd.Run()
	if err != nil {
		//Check for specific errors like ffprobe not found or network issues.  More robust error handling could be added here.
		return 0, fmt.Errorf("ffprobe command failed: %w", err)
	}
	re := regexp.MustCompile(`\d+`)
	bitrateStr := re.FindString(buf.String())
	if bitrateStr == "" {
		return 0, fmt.Errorf("bitrate not found for URL: %s", videoURL)
	}
	bitrate, err := strconv.Atoi(bitrateStr)
	if err != nil {
		return 0, fmt.Errorf("convert bitrate to int failed: %w", err)
	}
	return bitrate, nil
}

func GenThumb(ctx context.Context, videoURL string, second float64, scale int) (io.ReadCloser, int, error) {
	width, height, err := eimage.ExtractDimensions(ctx, videoURL)
	if err != nil {
		return nil, 0, err
	}
	// 按长边缩放，短边padding
	// scale=width:height
	args := []string{
		"-i", videoURL,
		"-ss", strconv.FormatFloat(second, 'g', -1, 64),
		"-q:v", "2", // 设置图像质量 (可选)
		"-vframes", "1", // 只获取一帧
		"-f", "image2pipe",
		"-vcodec", "mjpeg", //  输出 mjpeg 格式
		"-strict", "-2", // 放宽了编码器的限制
	}
	var filters []string
	// 标准的全范围颜色空间格式,更好地保持原始颜色的精确度
	filters = append(filters, "format=yuvj420p")
	// 添加关键帧选择过滤器
	filters = append(filters, "select='eq(pict_type\\,I)'")
	// 如果原始图片大小比视频帧大就resize图片作为缩略图
	if scale < width || scale < height {
		var thumbVF string
		if width > height {
			thumbVF = fmt.Sprintf("scale=%d:-1:flags=lanczos", scale)
		} else {
			thumbVF = fmt.Sprintf("scale=-1:%d:flags=lanczos", scale)
		}
		filters = append(filters, thumbVF)
	}
	// 合并所有过滤器
	if len(filters) > 0 {
		args = append(args, "-vf", strings.Join(filters, ","))
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

func ExtractFrame(ctx context.Context, videoURL string, second int64) (io.ReadCloser, error) {
	cmd := exec.Command("ffmpeg",
		"-ss", strconv.FormatInt(second, 10),
		"-i", videoURL,
		"-vframes", "1", // 只获取一帧
		"-q:v", "2", // 设置图像质量 (可选)
		"-f", "image2pipe",
		"-vcodec", "mjpeg", //  输出 mjpeg 格式
		"-",
	)

	buf := ebuffer.NewPoolBuffer(5 << 30)
	cmd.Stdout = buf
	kxlogging.GetGlobalLogger().WithContext(ctx).Debugf("ffmpeg extract frame cmd: %s", cmd.String())
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to execute ffmpeg command: %w, output: %s", err, buf.String())
	}
	return buf, nil
}

// GenPreview generates a video preview from a URL and saves it to a temporary file.
func GenPreview(ctx context.Context, videoURL string, previewDuration time.Duration) (io.ReadCloser, int64, error) {
	tempFileName := fmt.Sprintf("preview-%s.mp4", uuid.New())
	toDuration := fmt.Sprintf("%02d:%02d", int(previewDuration.Minutes())%60, int(previewDuration.Seconds())%60)
	args := []string{
		"-to", toDuration,
		"-i", videoURL, //  替换为实际的视频文件路径或 URL
		"-c:a", "copy",
		//"-c:a", "aac",
		//"-ab", "32k",
		//"-ac", "1",
		"-c:v", "copy",
		//"-c:v", "libx264",
		//"-keyint_min", "150",
		//"-sc_threshold", "0",
		//"-r", "20",
		//"-vf", "scale=320:-1",
		//"-preset", "faster",
		"-movflags", "faststart",
	}
	//args = append(args, "-x264opts", "bframes=10:b-adapt=0")
	args = append(args, tempFileName)
	cmd := exec.Command("ffmpeg", args...)
	kxlogging.GetGlobalLogger().WithContext(ctx).Debugf("ffmpeg gen preview cmd: %s", cmd.String())
	//ffmpeg -to 00:30 -i '' -c:a aac -ab 32k -ac 1 -c:v  libx264 -x264opts  "bframes=10:b-adapt=0" -keyint_min 150 -sc_threshold 0 -r 20 -vf scale=320:-1 -preset faster bframe320.mp4
	output, err := cmd.CombinedOutput()
	if err != nil {
		os.Remove(tempFileName)
		return nil, 0, fmt.Errorf("running ffmpeg: %w, output: %s", err, string(output))
	}
	buf, err := ebuffer.NewTempFile(tempFileName)
	if err != nil {
		return nil, 0, fmt.Errorf("creating temp file: %w", err)
	}
	size, err := buf.Len()
	if err != nil {
		return nil, 0, err
	}
	return buf, size, nil
}
