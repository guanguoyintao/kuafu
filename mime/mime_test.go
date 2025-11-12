package emime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileExtensionFromMIMEType(t *testing.T) {
	tests := []struct {
		name          string
		mimeType      MIMEType
		wantExtension string
	}{
		{
			name:          "空MIME类型",
			mimeType:      "",
			wantExtension: "",
		},
		{
			name:          "JPEG图片",
			mimeType:      "image/jpeg",
			wantExtension: "jpeg",
		},
		{
			name:          "PNG图片",
			mimeType:      "image/png",
			wantExtension: "png",
		},
		{
			name:          "GIF图片",
			mimeType:      "image/gif",
			wantExtension: "gif",
		},
		{
			name:          "PDF文档",
			mimeType:      "application/pdf",
			wantExtension: "pdf",
		},
		{
			name:          "Word文档",
			mimeType:      "application/msword",
			wantExtension: "doc",
		},
		{
			name:          "Excel文档",
			mimeType:      "application/vnd.ms-excel",
			wantExtension: "xls",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExtension := GetFileExtensionFromMIMEType(tt.mimeType)
			assert.Equal(t, tt.wantExtension, gotExtension,
				"MIME type %s should return extension %s, but got %s",
				tt.mimeType, tt.wantExtension, gotExtension)
		})
	}
}

func TestIsVideoMIMEType(t *testing.T) {
	tests := []struct {
		name     string
		mimeType MIMEType
		want     bool
	}{
		{
			name:     "MP4视频",
			mimeType: "video/mp4",
			want:     true,
		},
		{
			name:     "QuickTime视频",
			mimeType: "video/quicktime",
			want:     true,
		},
		{
			name:     "MPEG视频",
			mimeType: "video/mpeg",
			want:     true,
		},
		{
			name:     "WebM视频",
			mimeType: "video/webm",
			want:     true,
		},
		{
			name:     "图片类型",
			mimeType: "image/jpeg",
			want:     false,
		},
		{
			name:     "音频类型",
			mimeType: "audio/mpeg",
			want:     false,
		},
		{
			name:     "空类型",
			mimeType: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsVideoMIMEType(tt.mimeType)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsAudioMIMEType(t *testing.T) {
	tests := []struct {
		name     string
		mimeType MIMEType
		want     bool
	}{
		{
			name:     "MP3音频",
			mimeType: "audio/mpeg",
			want:     true,
		},
		{
			name:     "WAV音频",
			mimeType: "audio/x-wav",
			want:     true,
		},
		{
			name:     "AAC音频",
			mimeType: "audio/aac",
			want:     true,
		},
		{
			name:     "OGG音频",
			mimeType: "audio/ogg",
			want:     true,
		},
		{
			name:     "视频类型",
			mimeType: "video/mp4",
			want:     false,
		},
		{
			name:     "图片类型",
			mimeType: "image/png",
			want:     false,
		},
		{
			name:     "空类型",
			mimeType: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAudioMIMEType(tt.mimeType)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetMIMETypeFromFileExtension(t *testing.T) {
	tests := []struct {
		name           string
		extension      string
		expectedMIME   MIMEType
		expectedExists bool
	}{
		// 视频文件测试用例
		{
			name:           "M2TS视频",
			extension:      ".m2ts",
			expectedMIME:   "video/mp2t",
			expectedExists: true,
		},
		{
			name:           "MP4视频",
			extension:      ".mp4",
			expectedMIME:   "video/mp4",
			expectedExists: true,
		},
		{
			name:           "MOV视频",
			extension:      ".mov",
			expectedMIME:   "video/quicktime",
			expectedExists: true,
		},
		{
			name:           "MPEG视频",
			extension:      ".mpeg",
			expectedMIME:   "video/mpeg",
			expectedExists: true,
		},
		{
			name:           "MKV视频",
			extension:      ".mkv",
			expectedMIME:   "video/x-matroska",
			expectedExists: true,
		},
		{
			name:           "WMV视频",
			extension:      ".wmv",
			expectedMIME:   "video/x-ms-wmv",
			expectedExists: true,
		},
		{
			name:           "M3U8视频",
			extension:      ".m3u8",
			expectedMIME:   "application/vnd.apple.mpegurl",
			expectedExists: true,
		},
		{
			name:           "3G2视频",
			extension:      ".3g2",
			expectedMIME:   "video/3gpp2",
			expectedExists: true,
		},
		{
			name:           "3GPP视频",
			extension:      ".3gpp",
			expectedMIME:   "video/3gpp",
			expectedExists: true,
		},
		{
			name:           "DVB视频",
			extension:      ".dvb",
			expectedMIME:   "video/vnd.dvb.file",
			expectedExists: true,
		},
		{
			name:           "FLI视频",
			extension:      ".fli",
			expectedMIME:   "video/x-fli",
			expectedExists: true,
		},
		{
			name:           "FLV视频",
			extension:      ".flv",
			expectedMIME:   "video/x-flv",
			expectedExists: true,
		},
		{
			name:           "H261视频",
			extension:      ".h261",
			expectedMIME:   "video/h261",
			expectedExists: true,
		},
		{
			name:           "H263视频",
			extension:      ".h263",
			expectedMIME:   "video/h263",
			expectedExists: true,
		},
		{
			name:           "H264视频",
			extension:      ".h264",
			expectedMIME:   "video/h264",
			expectedExists: true,
		},
		{
			name:           "JPGV视频",
			extension:      ".jpgv",
			expectedMIME:   "video/jpeg",
			expectedExists: true,
		},
		{
			name:           "MJ2视频",
			extension:      ".mj2",
			expectedMIME:   "video/mj2",
			expectedExists: true,
		},
		{
			name:           "MNG视频",
			extension:      ".mng",
			expectedMIME:   "video/x-mng",
			expectedExists: true,
		},
		{
			name:           "OGV视频",
			extension:      ".ogv",
			expectedMIME:   "video/ogg",
			expectedExists: true,
		},
		{
			name:           "WebM视频",
			extension:      ".webm",
			expectedMIME:   "video/webm",
			expectedExists: true,
		},

		// 音频文件测试用例
		{
			name:           "ADPCM音频",
			extension:      ".adp",
			expectedMIME:   "audio/adpcm",
			expectedExists: true,
		},
		{
			name:           "WAV音频",
			extension:      ".wav",
			expectedMIME:   "audio/wav",
			expectedExists: true,
		},
		{
			name:           "AAC音频",
			extension:      ".aac",
			expectedMIME:   "audio/aac",
			expectedExists: true,
		},
		{
			name:           "FLAC音频",
			extension:      ".flac",
			expectedMIME:   "audio/flac",
			expectedExists: true,
		},
		{
			name:           "OGG音频",
			extension:      ".ogg",
			expectedMIME:   "audio/ogg",
			expectedExists: true,
		},
		{
			name:           "AMR音频",
			extension:      ".amr",
			expectedMIME:   "audio/amr",
			expectedExists: true,
		},
		{
			name:           "AC3音频",
			extension:      ".ac3",
			expectedMIME:   "audio/ac3",
			expectedExists: true,
		},
		{
			name:           "MP3音频",
			extension:      ".mp3",
			expectedMIME:   "audio/mpeg",
			expectedExists: true,
		},
		{
			name:           "MKA音频",
			extension:      ".mka",
			expectedMIME:   "audio/x-matroska",
			expectedExists: true,
		},
		{
			name:           "M4A音频",
			extension:      ".m4a",
			expectedMIME:   "audio/mp4",
			expectedExists: true,
		},
		{
			name:           "AIFF音频",
			extension:      ".aiff",
			expectedMIME:   "audio/x-aiff",
			expectedExists: true,
		},
		{
			name:           "CAF音频",
			extension:      ".caf",
			expectedMIME:   "audio/x-caf",
			expectedExists: true,
		},
		{
			name:           "DTS音频",
			extension:      ".dts",
			expectedMIME:   "audio/vnd.dts",
			expectedExists: true,
		},
		{
			name:           "DTSHD音频",
			extension:      ".dtshd",
			expectedMIME:   "audio/vnd.dts.hd",
			expectedExists: true,
		},
		{
			name:           "MIDI音频",
			extension:      ".midi",
			expectedMIME:   "audio/midi",
			expectedExists: true,
		},
		{
			name:           "WMA音频",
			extension:      ".wma",
			expectedMIME:   "audio/x-ms-wma",
			expectedExists: true,
		},

		// 图片文件测试用例
		{
			name:           "3DS图片",
			extension:      ".3ds",
			expectedMIME:   "image/x-3ds",
			expectedExists: true,
		},
		{
			name:           "APNG图片",
			extension:      ".apng",
			expectedMIME:   "image/apng",
			expectedExists: true,
		},
		{
			name:           "BMP图片",
			extension:      ".bmp",
			expectedMIME:   "image/x-ms-bmp",
			expectedExists: true,
		},
		{
			name:           "CGM图片",
			extension:      ".cgm",
			expectedMIME:   "image/cgm",
			expectedExists: true,
		},
		{
			name:           "DNG图片",
			extension:      ".dng",
			expectedMIME:   "image/x-adobe-dng",
			expectedExists: true,
		},
		{
			name:           "GIF图片",
			extension:      ".gif",
			expectedMIME:   "image/gif",
			expectedExists: true,
		},
		{
			name:           "HEIC图片",
			extension:      ".heic",
			expectedMIME:   "image/heic",
			expectedExists: true,
		},
		{
			name:           "ICO图片",
			extension:      ".ico",
			expectedMIME:   "image/x-icon",
			expectedExists: true,
		},
		{
			name:           "JPEG图片",
			extension:      ".jpeg",
			expectedMIME:   "image/jpeg",
			expectedExists: true,
		},
		{
			name:           "JPG图片",
			extension:      ".jpg",
			expectedMIME:   "image/jpeg",
			expectedExists: true,
		},
		{
			name:           "PNG图片",
			extension:      ".png",
			expectedMIME:   "image/png",
			expectedExists: true,
		},
		{
			name:           "SVG图片",
			extension:      ".svg",
			expectedMIME:   "image/svg+xml",
			expectedExists: true,
		},
		{
			name:           "TIFF图片",
			extension:      ".tiff",
			expectedMIME:   "image/tiff",
			expectedExists: true,
		},
		{
			name:           "WEBP图片",
			extension:      ".webp",
			expectedMIME:   "image/webp",
			expectedExists: true,
		},
		{
			name:           "CR2图片",
			extension:      ".cr2",
			expectedMIME:   "image/x-canon-cr2",
			expectedExists: true,
		},
		{
			name:           "NEF图片",
			extension:      ".nef",
			expectedMIME:   "image/x-nikon-nef",
			expectedExists: true,
		},

		// 特殊情况测试
		{
			name:           "未知扩展名",
			extension:      ".unknown",
			expectedMIME:   "application/unknown",
			expectedExists: false,
		},
		{
			name:           "空扩展名",
			extension:      "",
			expectedMIME:   "application/unknown",
			expectedExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mimeType, exists := GetMIMETypeFromFileExtension(tt.extension)
			assert.Equal(t, tt.expectedMIME, mimeType, "MIME类型不匹配")
			assert.Equal(t, tt.expectedExists, exists, "存在性检查结果不匹配")
		})
	}
}
