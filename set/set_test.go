package eset

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New[int](10)
	assert.NotNil(t, s)
	assert.Empty(t, s.ToSlice())
}

func TestNewFromSlice(t *testing.T) {
	t.Run("with non-empty slice", func(t *testing.T) {
		elements := []string{"apple", "banana", "cherry"}
		s := NewFromSlice(elements)
		assert.NotNil(t, s)
		assert.Len(t, s.ToSlice(), len(elements))
		assert.True(t, s.Contains("apple"))
		assert.True(t, s.Contains("banana"))
		assert.True(t, s.Contains("cherry"))
	})

	t.Run("with empty slice", func(t *testing.T) {
		emptySlice := []int{}
		s := NewFromSlice(emptySlice)
		assert.NotNil(t, s)
		assert.Empty(t, s.ToSlice())
	})

	t.Run("with nil slice", func(t *testing.T) {
		s := NewFromSlice[float64](nil)
		assert.NotNil(t, s)
		assert.Empty(t, s.ToSlice())
	})
}

func TestSet_Add(t *testing.T) {
	s := New[int](0)
	s.Add(1)
	s.Add(1)
	s.Add(2)
	assert.Len(t, s.ToSlice(), 2)
	assert.True(t, s.Contains(1))

	s.Add(1) // Adding the same element again
	assert.Len(t, s.ToSlice(), 2)
}

func TestSet_Remove(t *testing.T) {
	s := NewFromSlice([]int{1, 2, 3})
	s.Remove(2)
	assert.Len(t, s.ToSlice(), 2)
	assert.False(t, s.Contains(2))

	s.Remove(4) // Removing a non-existent element
	assert.Len(t, s.ToSlice(), 2)
}

func TestSet_Contains(t *testing.T) {
	s := NewFromSlice([]string{"hello", "world"})
	assert.True(t, s.Contains("hello"))
	assert.True(t, s.Contains("world"))
	assert.False(t, s.Contains("go"))
}

func TestSet_Len(t *testing.T) {
	s1 := New[bool](0)
	assert.Len(t, s1.ToSlice(), 0)

	s2 := NewFromSlice([]float64{1.1, 2.2, 3.3, 4.4})
	assert.Len(t, s2.ToSlice(), 4)
}

func TestSet_IsEmpty(t *testing.T) {
	s1 := New[rune](5)
	assert.True(t, s1.IsEmpty())

	s2 := NewFromSlice([]rune{'a', 'b'})
	assert.False(t, s2.IsEmpty())
}

func TestSet_Clear(t *testing.T) {
	s := NewFromSlice([]int{10, 20, 30})
	s.Clear()
	assert.Empty(t, s.ToSlice())
	assert.True(t, s.IsEmpty())
}

func TestSet_ToSlice(t *testing.T) {
	elements := []int{5, 2, 8, 1}
	s := NewFromSlice(elements)
	slice := s.ToSlice()
	assert.Len(t, slice, len(elements))
	elementMap := make(map[int]bool)
	for _, elem := range elements {
		elementMap[elem] = true
	}
	for _, slElem := range slice {
		assert.True(t, elementMap[slElem], "unexpected element in slice")
	}

	emptySet := New[string](0)
	emptySlice := emptySet.ToSlice()
	assert.Empty(t, emptySlice)
}

func TestSet_Union(t *testing.T) {
	testCases := []struct {
		name string
		s1   *Set[int]
		s2   *Set[int]
		want []int
	}{
		{"basic union", NewFromSlice([]int{1, 2, 3}), NewFromSlice([]int{3, 4, 5}), []int{1, 2, 3, 4, 5}},
		{"union with empty set", NewFromSlice([]int{1, 2}), New[int](0), []int{1, 2}},
		{"empty set union with set", New[int](0), NewFromSlice([]int{3, 4}), []int{3, 4}},
		{"union of empty sets", New[int](0), New[int](0), []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			unionSet := tc.s1.Union(tc.s2)
			got := unionSet.ToSlice()
			sort.Ints(got)
			sort.Ints(tc.want)
			assert.ElementsMatch(t, got, tc.want)
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	testCases := []struct {
		name string
		s1   *Set[string]
		s2   *Set[string]
		want []string
	}{
		{"basic intersection", NewFromSlice([]string{"a", "b", "c"}), NewFromSlice([]string{"b", "c", "d"}), []string{"b", "c"}},
		{"no intersection", NewFromSlice([]string{"a", "b"}), NewFromSlice([]string{"c", "d"}), []string{}},
		{"intersection with empty set", NewFromSlice([]string{"a", "b"}), New[string](0), []string{}},
		{"empty set intersection with set", New[string](0), NewFromSlice([]string{"c", "d"}), []string{}},
		{"intersection of empty sets", New[string](0), New[string](0), []string{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			intersectionSet := tc.s1.Intersection(tc.s2)
			got := intersectionSet.ToSlice()
			sort.Strings(got)
			sort.Strings(tc.want)
			assert.ElementsMatch(t, got, tc.want)
		})
	}
}

func TestSet_Difference(t *testing.T) {
	testCases := []struct {
		name string
		s1   *Set[int]
		s2   *Set[int]
		want []int
	}{
		{"basic difference", NewFromSlice([]int{10, 20, 30, 40}), NewFromSlice([]int{30, 40, 50, 60}), []int{10, 20}},
		{"difference with empty set", NewFromSlice([]int{1, 2}), New[int](0), []int{1, 2}},
		{"empty set difference with set", New[int](0), NewFromSlice([]int{3, 4}), []int{}},
		{"difference of empty sets", New[int](0), New[int](0), []int{}},
		{"disjoint sets difference", NewFromSlice([]int{1, 2}), NewFromSlice([]int{3, 4}), []int{1, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			differenceSet := tc.s1.Difference(tc.s2)
			got := differenceSet.ToSlice()
			sort.Ints(got)
			sort.Ints(tc.want)
			assert.ElementsMatch(t, got, tc.want)
		})
	}
}

func TestSet_IsSubsetOf(t *testing.T) {
	testCases := []struct {
		name string
		s1   *Set[string]
		s2   *Set[string]
		want bool
	}{
		{"is subset", NewFromSlice([]string{"x", "y"}), NewFromSlice([]string{"w", "x", "y", "z"}), true},
		{"is not subset (longer)", NewFromSlice([]string{"w", "x", "y", "z"}), NewFromSlice([]string{"x", "y"}), false},
		{"equal sets", NewFromSlice([]string{"a", "b"}), NewFromSlice([]string{"b", "a"}), true},
		{"empty set is subset", New[string](0), NewFromSlice([]string{"a", "b"}), true},
		{"set is not subset of empty set", NewFromSlice([]string{"a", "b"}), New[string](0), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s1.IsSubsetOf(tc.s2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSet_IsSupersetOf(t *testing.T) {
	testCases := []struct {
		name string
		s1   *Set[float64]
		s2   *Set[float64]
		want bool
	}{
		{"is superset", NewFromSlice([]float64{3.14, 2.71, 1.618}), NewFromSlice([]float64{3.14, 2.71}), true},
		{"is not superset (shorter)", NewFromSlice([]float64{3.14, 2.71}), NewFromSlice([]float64{3.14, 2.71, 1.618}), false},
		{"equal sets", NewFromSlice([]float64{1.0, 2.0}), NewFromSlice([]float64{2.0, 1.0}), true},
		{"set is superset of empty set", NewFromSlice([]float64{1.0, 2.0}), New[float64](0), true},
		{"empty set is not superset of set", New[float64](0), NewFromSlice([]float64{1.0, 2.0}), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s1.IsSupersetOf(tc.s2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSet_Equal(t *testing.T) {
	testCases := []struct {
		name string
		s1   *Set[int]
		s2   *Set[int]
		want bool
	}{
		{"equal sets", NewFromSlice([]int{1, 2, 3}), NewFromSlice([]int{3, 2, 1}), true},
		{"not equal (different element)", NewFromSlice([]int{1, 2, 3}), NewFromSlice([]int{1, 2, 4}), false},
		{"not equal (different length)", NewFromSlice([]int{1, 2}), NewFromSlice([]int{1, 2, 3}), false},
		{"empty sets are equal", New[int](0), New[int](0), true},
		{"empty and non-empty are not equal", New[int](0), NewFromSlice([]int{1}), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s1.Equal(tc.s2)
			assert.Equal(t, tc.want, got)
		})
	}
}
