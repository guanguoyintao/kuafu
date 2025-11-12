package eid

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Base62 字符集（62 个字符）
const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// GenerateCode 生成唯一码
func GenerateCode(length int) string {
	node, _ := snowflake.NewNode(1) // 机器 ID
	id := node.Generate().Int64()   // 生成唯一 ID
	code := encodeBase62(id)        // 转换为 Base62 编码

	// 如果长度不足，随机填充字符
	for len(code) < length {
		code += string(base62Chars[rand.Intn(len(base62Chars))])
	}

	// 截取指定长度
	return code[:length]
}

// Base62 编码
func encodeBase62(num int64) string {
	if num == 0 {
		return "0"
	}

	var result string
	for num > 0 {
		result = string(base62Chars[num%62]) + result
		num /= 62
	}
	return result
}
