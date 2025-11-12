package earray

// Take 返回切片的前 n 个元素，如果切片长度不足 n，则返回整个切片。
func Take[T any](s []T, n int) []T {
	if n > len(s) {
		return s
	}
	return s[:n]
}
