package ehash

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spaolacci/murmur3"
)

const (
	CHUNK_SIZE = 2 * 1024 * 1024         // 2MB
	SIZE_12M   = 12 * 1024 * 1024        // 12MB
	SIZE_1G    = 1024 * 1024 * 1024      // 1GB
	SIZE_10G   = 10 * 1024 * 1024 * 1024 // 10GB
)

func hashToBase62WithLength(hash uint32, length int) string {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base := uint32(len(chars)) // 62

	// 生成指定长度的62进制字符串
	result := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		result[i] = chars[hash%base]
		hash /= base
	}
	return string(result)
}

// 将32位哈希值转成6位62进制字符串
func hashToBase62(hash uint32) string {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base := uint32(len(chars)) // 62

	// 生成6位的62进制字符串
	result := make([]byte, 6)
	for i := 5; i >= 0; i-- {
		result[i] = chars[hash%base]
		hash /= base
	}
	return string(result)
}

func HashMD532(s string) string {
	sum := md5.Sum([]byte(strings.TrimSpace(s)))
	return hex.EncodeToString(sum[:])
}

func HashMurmurHash340(s string) (string, error) {
	h := murmur3.New128()
	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}
	sum := h.Sum([]byte(s))

	return hex.EncodeToString(sum[:]), nil
}

func HashMurmurHash36(s string) (string, error) {
	num := murmur3.Sum32([]byte(s))
	return hashToBase62(num), nil
}

func HashMurmurHash36WithLength(s string, length int) (string, error) {
	num := murmur3.Sum32([]byte(s))
	return hashToBase62WithLength(num, length), nil
}

func CalcFileSHA256(ctx context.Context, filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// 创建 SHA256 哈希对象
	hash := sha256.New()
	// 读取文件内容并计算哈希值
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	// 获取计算得到的哈希值
	hashValue := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hashString := fmt.Sprintf("%x", hashValue)
	return hashString, nil
}

func CalcContentSHA256(ctx context.Context, content io.Reader) (string, error) {
	// 创建 SHA256 哈希对象
	hash := sha256.New()
	// 读取文件内容并计算哈希值
	if _, err := io.Copy(hash, content); err != nil {
		return "", err
	}
	// 获取计算得到的哈希值
	hashValue := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hashString := fmt.Sprintf("%x", hashValue)
	return hashString, nil
}

func CalcFileSHA1(ctx context.Context, filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// 创建 SHA1 哈希对象
	hash := sha1.New()
	// 读取文件内容并计算哈希值
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	// 获取计算得到的哈希值
	hashValue := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hashString := fmt.Sprintf("%x", hashValue)
	return hashString, nil
}

func CalcContentSHA1(ctx context.Context, content io.Reader) (string, error) {
	// 创建 SHA1 哈希对象
	hash := sha1.New()
	// 读取文件内容并计算哈希值
	if _, err := io.Copy(hash, content); err != nil {
		return "", err
	}
	// 获取计算得到的哈希值
	hashValue := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hashString := fmt.Sprintf("%x", hashValue)
	return hashString, nil
}

func CalcFileResourceHash(filePath string) (string, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	fileSize := fileInfo.Size()
	sizeBytes := []byte(fmt.Sprintf("%016d", fileSize))
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	hash.Write(sizeBytes)
	var chunks [][]byte
	if fileSize <= SIZE_12M {
		chunks, err = readAllChunks(file)
	} else {
		positions := getSamplePositions(fileSize)
		chunks, err = readChunksByPositions(file, positions, fileSize)
	}
	if err != nil {
		return "", err
	}
	for _, chunk := range chunks {
		hash.Write(chunk)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func readAllChunks(file *os.File) ([][]byte, error) {
	var chunks [][]byte
	buf := make([]byte, CHUNK_SIZE)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		chunk := make([]byte, n)
		copy(chunk, buf[:n])
		chunks = append(chunks, chunk)
	}
	return chunks, nil
}

func readChunksByPositions(file *os.File, positions []int64, fileSize int64) ([][]byte, error) {
	var chunks [][]byte
	for _, pos := range positions {
		readPos := pos - CHUNK_SIZE/2
		if pos == 0 {
			readPos = 0
		} else if pos == fileSize {
			readPos = fileSize - CHUNK_SIZE
		}
		if readPos < 0 {
			readPos = 0
		}
		buf := make([]byte, CHUNK_SIZE)
		_, err := file.Seek(readPos, io.SeekStart)
		if err != nil {
			return nil, err
		}
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		chunks = append(chunks, buf[:n])
	}
	return chunks, nil
}

func getSamplePositions(fileSize int64) []int64 {
	var samplePoints int
	switch {
	case fileSize <= SIZE_1G:
		samplePoints = 5
	case fileSize <= SIZE_10G:
		samplePoints = 10
	default:
		samplePoints = 15
	}
	positions := []int64{0}
	remaining := samplePoints - 2
	if remaining > 0 {
		if remaining%2 == 1 {
			half := remaining / 2
			positions = append(positions, fileSize/2)
			for i := 1; i <= half; i++ {
				// 左半区
				left := int64(float64(i) * 50 / float64(half+1) * float64(fileSize) / 100)
				positions = append(positions, left)
				// 右半区
				right := int64((50 + float64(i)*50/float64(half+1)) * float64(fileSize) / 100)
				positions = append(positions, right)
			}
		} else {
			for i := 1; i <= remaining; i++ {
				pos := int64(float64(i) * 100.0 / float64(remaining+1) * float64(fileSize) / 100)
				positions = append(positions, pos)
			}
		}
	}
	positions = append(positions, fileSize)
	return uniqueSortedPositions(positions)
}

func uniqueSortedPositions(positions []int64) []int64 {
	posMap := make(map[int64]bool)
	for _, p := range positions {
		posMap[p] = true
	}
	var result []int64
	for p := range posMap {
		result = append(result, p)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}
