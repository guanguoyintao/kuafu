package ecopier

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect" // 引入 reflect 包用于 nil 值检查
)

type Cloner[T any] interface {
	Clone() T
}

func DeepCopy[T any](src T) (T, error) {
	if cloner, ok := any(src).(Cloner[T]); ok {
		return cloner.Clone(), nil
	}
	// 在使用 gob 之前，必须处理 nil 输入的情况，因为 gob 会 panic
	val := reflect.ValueOf(src)
	// 对于指针、接口、map、slice 等引用类型，检查它们是否为 nil
	// IsValid() 用于处理完全无类型的 nil 接口
	if !val.IsValid() {
		var zero T
		return zero, nil
	}
	switch val.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface:
		if val.IsNil() {
			var zero T // 对于引用类型，零值就是 nil
			return zero, nil
		}
	}
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	decoder := gob.NewDecoder(&buffer)
	if err := encoder.Encode(src); err != nil {
		var zero T
		return zero, fmt.Errorf("gob encode error: %w", err)
	}
	var dst T
	if err := decoder.Decode(&dst); err != nil {
		var zero T
		return zero, fmt.Errorf("gob decode error: %w", err)
	}
	return dst, nil
}
