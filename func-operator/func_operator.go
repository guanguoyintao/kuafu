package efuncop

// FlatMap 函数式编程中FlatMap 方法
func FlatMap[T any](values []T, f func(T) ([]T, error)) ([]T, error) {
	var result []T
	for _, v := range values {
		m, err := f(v)
		if err != nil {
			return nil, err
		}
		result = append(result, m...)
	}
	return result, nil
}

// Filter 函数式编程中的 `filter` 操作
func Filter[T any](values []T, predicate func(T) (bool, error)) ([]T, error) {
	var filtered []T
	for _, v := range values {
		ok, err := predicate(v)
		if err != nil {
			return nil, err
		}
		if ok {
			filtered = append(filtered, v)
		}
	}
	return filtered, nil
}

// Reduce 实现函数式编程中的 reduce 操作
func Reduce[T1 any, T2 any](values []T1, f func(T2, T1) (T2, error), initial T2) (T2, error) {
	accum := initial
	var empty T2
	var err error
	for _, v := range values {
		accum, err = f(accum, v)
		if err != nil {
			return empty, err
		}
	}
	return accum, nil
}

// Map 函数式编程中Map 方法
func Map[T1 any, T2 any](values []T1, f func(T1) (T2, error)) ([]T2, error) {
	newValues := make([]T2, len(values))
	var err error
	for i, v := range values {
		newValues[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return newValues, nil
}
