package efile

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

// RemoveFileExtension 去掉文件名后缀的函数
func RemoveFileExtension(fileName string) (string, string) {
	// 找到最后一个点号的位置
	lastDotIndex := strings.LastIndex(fileName, ".")
	// 如果没有点号，直接返回原文件名
	if lastDotIndex == -1 {
		return fileName, ""
	}
	// 去掉后缀
	return fileName[:lastDotIndex], fileName[lastDotIndex+1:]
}

func CreateFileFromReader(reader io.Reader, filePath string, bufferSize int) (string, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	buf := make([]byte, bufferSize)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return "", fmt.Errorf("failed to read data: %w", err)
		}
		if n == 0 {
			break
		}
		if _, err := file.Write(buf[:n]); err != nil {
			return "", fmt.Errorf("failed to write data: %w", err)
		}
	}

	return filePath, nil
}

func DownloadURLFileAsReader(url string) (io.ReadCloser, error) {
	// 创建 Resty 客户端
	client := resty.New()
	client.SetTimeout(3 * time.Minute)
	client.SetDoNotParseResponse(true)
	// 发起 GET 请求
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}
	// 返回响应的 Body 作为 io.ReadCloser
	return resp.RawBody(), nil
}
