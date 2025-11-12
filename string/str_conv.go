package estrings

import (
	"errors"
	"math"
	"strings"
)

// const chars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
// 基于标准62个字符shuffle，减少反解风险
const chars string = "t2RiOnNoHdzjCpkAMvShZD8eXTuK14BV7J5fagsqYrW0Gb6LlwQxc3EyFU9ImP"

// ConvBase10To62 10进制转62进制
func ConvBase10To62(num uint64) string {
	bytes := []byte{}
	if num == 0 {
		return string(chars[0])
	}
	for num > 0 {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

// ConvBase62To10 62进制转10进制
func ConvBase62To10(str string) (uint64, error) {
	var num uint64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(chars, str[i])
		if pos < 0 {
			return 0, errors.New("invalid base62 string")
		}
		num += uint64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num, nil
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
