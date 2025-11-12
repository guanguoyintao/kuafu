package earray

// LeftJoin 函数实现左连接
func LeftJoin[T1 any, T2 any](leftArray []T1, rightArray []T2, leftIDFunc func(T1) string, rightIDFunc func(T2) string) []LeftJoinResult[T1, T2] {
	result := make([]LeftJoinResult[T1, T2], 0, len(leftArray)+len(rightArray))
	// 初始化结果，包含所有左侧元素
	for _, left := range leftArray {
		result = append(result, LeftJoinResult[T1, T2]{
			Left:  left,
			Right: []T2{}, // 初始化右侧数组为空
		})
	}
	// 关联右侧元素
	for _, right := range rightArray {
		rightID := rightIDFunc(right)
		// 查找对应的左侧元素并更新右侧数组
		for i := range result {
			if leftIDFunc(result[i].Left) == rightID {
				result[i].Right = append(result[i].Right, right)
			}
		}
	}

	return result
}

// RightJoin 函数实现右连接
func RightJoin[T1 any, T2 any](leftArray []T1, rightArray []T2, leftIDFunc func(T1) string, rightIDFunc func(T2) string) []RightJoinResult[T1, T2] {
	result := make([]RightJoinResult[T1, T2], 0, len(leftArray)+len(rightArray))
	// 使用一个 map 存储左侧元素，方便快速查找
	leftMap := make(map[string][]T1) // 左侧可以有多个对应项
	for _, left := range leftArray {
		leftID := leftIDFunc(left)
		leftMap[leftID] = append(leftMap[leftID], left) // 允许左侧有多个对应项
	}
	// 遍历右侧元素并查找对应的左侧元素
	for _, right := range rightArray {
		rightID := rightIDFunc(right)
		leftValues := leftMap[rightID] // 获取对应的左侧元素
		// 将左侧和右侧对应项添加到结果中
		result = append(result, RightJoinResult[T1, T2]{Left: leftValues, Right: right})
	}

	return result
}

// Distinct 函数实现去重
func Distinct[T any, K comparable](slice []T, keyFunc func(T) K) []T {
	uniqueMap := make(map[K]T)
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		key := keyFunc(item) // 获取唯一标识
		if _, exists := uniqueMap[key]; !exists {
			uniqueMap[key] = item
			result = append(result, item) // 添加到结果切片
		}
	}

	return result
}
