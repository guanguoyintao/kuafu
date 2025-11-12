package exp

import "testing"

func TestIn(t *testing.T) {
	tests := []struct {
		name     string
		val      interface{}
		slice    interface{}
		expected bool
	}{
		{
			name:     "Value exists in slice of integers",
			val:      3,
			slice:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Value does not exist in slice of integers",
			val:      6,
			slice:    []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "Value exists in slice of strings",
			val:      "apple",
			slice:    []string{"banana", "apple", "orange"},
			expected: true,
		},
		{
			name:     "Value does not exist in slice of strings",
			val:      "grape",
			slice:    []string{"banana", "apple", "orange"},
			expected: false,
		},
		{
			name:     "Empty slice",
			val:      1,
			slice:    []int{},
			expected: false,
		},
		{
			name:     "Value exists in a slice with a single element",
			val:      "test",
			slice:    []string{"test"},
			expected: true,
		},
		{
			name:     "Value does not exist in a slice with a single element",
			val:      "test",
			slice:    []string{"hello"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert val and slice to the proper types before passing to the In function
			var result bool
			switch v := tt.slice.(type) {
			case []int:
				result = In(tt.val.(int), v)
			case []string:
				result = In(tt.val.(string), v)
			default:
				t.Fatalf("Unsupported type in test case: %T", v)
			}
			if result != tt.expected {
				t.Errorf("In(%v, %v) = %v; expected %v", tt.val, tt.slice, result, tt.expected)
			}
		})
	}
}
