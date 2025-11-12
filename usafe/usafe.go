package eusafe

import "unsafe"

// GetSize 用于计算变量的字节大小
func GetSize(v any) int {
	return int(unsafe.Sizeof(v))
}
