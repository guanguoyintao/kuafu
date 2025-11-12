package erand

import (
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomRange 生成指定范围随机数，左右闭，即[min,max]
func RandomRange(min, max int) int {
	return seededRand.Intn(max+1-min) + min
}

// RandomRangeFloat64 生成指定范围随机小数，左右闭，即[min,max]
func RandomRangeFloat64(min, max float64) float64 {
	return min + seededRand.Float64()*(max-min)
}

// GenRandomString 生成指定长度的随机字符串
func GenRandomString(length int) string {
	if length <= 0 {
		length = 0
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GenRandomAlphaNumeric 生成指定长度的随机码，一定包含字母和数字
func GenRandomAlphaNumeric(length int) string {
	if length < 2 {
		return GenRandomString(length)
	}
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"
	const charset = letters + digits
	// 确保至少包含一个字母和一个数字
	result := make([]byte, length)
	result[0] = letters[seededRand.Intn(len(letters))]
	result[1] = digits[seededRand.Intn(len(digits))]
	// 用完整字符集填充剩余字符
	for i := 2; i < length; i++ {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	// 打乱结果以确保随机性
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}
