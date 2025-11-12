package eset

import ejson "github.com/guanguoyintao/kuafu/json"

type Set[T comparable] struct {
	set map[string]T
}

// New 创建并返回一个新的空 Set。
func New[T comparable](capacity int) *Set[T] {
	return &Set[T]{
		set: make(map[string]T, capacity),
	}
}

// NewFromSlice 创建并返回一个新的 Set，并使用给定的切片中的元素进行初始化。
func NewFromSlice[T comparable](elements []T) *Set[T] {
	if elements == nil {
		return New[T](0)
	}
	s := New[T](len(elements))
	for _, element := range elements {
		s.Add(element)
	}
	return s
}

// Add 将给定的元素添加到集合中。
func (s *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		key, err := ejson.MarshalString(element)
		if err != nil {
			return
		}
		s.set[key] = element
	}
}

// Remove 从集合中移除给定的元素。
func (s *Set[T]) Remove(elements ...T) {
	for _, element := range elements {
		key, err := ejson.MarshalString(element)
		if err != nil {
			return
		}
		delete(s.set, key)
	}
}

// Contains 检查给定的元素是否存在于集合中。
func (s *Set[T]) Contains(element T) bool {
	key, err := ejson.MarshalString(element)
	if err != nil {
		return false
	}
	_, ok := s.set[key]
	return ok
}

// Len 返回集合中元素的数量。
func (s *Set[T]) Len() int {
	return len(s.set)
}

// IsEmpty 检查集合是否为空。
func (s *Set[T]) IsEmpty() bool {
	return len(s.set) == 0
}

// Clear 移除集合中的所有元素。
func (s *Set[T]) Clear() {
	clear(s.set)
}

// ToSlice 返回一个包含集合中所有元素的切片。切片中元素的顺序不作保证。
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.set))
	for _, element := range s.set {
		slice = append(slice, element)
	}
	return slice
}

// Union 返回一个新的集合，其中包含当前集合和另一个集合中的所有元素。
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	unionSet := New[T](s.Len() + other.Len())
	for _, element := range s.set {
		unionSet.Add(element)
	}
	for _, element := range other.set {
		unionSet.Add(element)
	}
	return unionSet
}

// Intersection 返回一个新的集合，其中包含同时存在于当前集合和另一个集合中的元素。
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersectionSet := New[T](0)
	smaller := s
	larger := other
	if len(other.set) < len(s.set) {
		smaller = other
		larger = s
	}
	for _, element := range smaller.set {
		if larger.Contains(element) {
			intersectionSet.Add(element)
		}
	}
	return intersectionSet
}

// Difference 返回一个新的集合，其中包含存在于当前集合但不存在于另一个集合中的元素。
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	differenceSet := New[T](0)
	for _, element := range s.set {
		if !other.Contains(element) {
			differenceSet.Add(element)
		}
	}
	return differenceSet
}

// IsSubsetOf 检查当前集合是否是另一个集合的子集。
func (s *Set[T]) IsSubsetOf(other *Set[T]) bool {
	if len(s.set) > len(other.set) {
		return false
	}
	for _, element := range s.set {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// IsSupersetOf 检查当前集合是否是另一个集合的超集。
func (s *Set[T]) IsSupersetOf(other *Set[T]) bool {
	return other.IsSubsetOf(s)
}

// Equal 检查当前集合是否与另一个集合相等（包含相同的元素）。
func (s *Set[T]) Equal(other *Set[T]) bool {
	if len(s.set) != len(other.set) {
		return false
	}
	for _, element := range s.set {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}
