package earray

// LeftJoinResult 定义左侧连接结果的结构体
type LeftJoinResult[T1 any, T2 any] struct {
	Left  T1
	Right []T2 // 右侧可以有多个对应项
}

// RightJoinResult 定义右连接结果的结构体
type RightJoinResult[T1 any, T2 any] struct {
	Left  []T1 // 左侧可以有多个对应项
	Right T2   // 右侧
}
