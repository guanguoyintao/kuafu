package efuncop

import (
	gerrors "errors"
	"reflect"
	"strings"
	"testing"
	"time"
)

type Person struct {
	Name string
	Age  int
}

// Directory 结构体
type Directory struct {
	ID        uint64
	UID       uint64
	ParentID  uint64
	Name      string
	Children  []*Directory
	Leaf      []*Directory
	Ancestors []*Directory
	UpdatedAt time.Time
	CreatedAt time.Time
	IsRoot    bool
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		values   interface{} // 输入的值
		f        interface{} // 传入的函数
		expected interface{} // 期望的输出
	}{
		{
			name:     "Square each integer",
			values:   []int{1, 2, 3, 4},
			f:        func(x int) int { return x * x },
			expected: []int{1, 4, 9, 16},
		},
		{
			name:     "Uppercase each string",
			values:   []string{"hello", "world"},
			f:        func(s string) string { return strings.ToUpper(s) },
			expected: []string{"HELLO", "WORLD"},
		},
		{
			name: "Double age in Person struct",
			values: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
			},
			f: func(p Person) Person {
				return Person{Name: p.Name, Age: p.Age * 2}
			},
			expected: []Person{
				{Name: "Alice", Age: 50},
				{Name: "Bob", Age: 60},
			},
		},
		{
			name:     "Halve each float",
			values:   []float64{2.0, 4.0, 6.0},
			f:        func(x float64) float64 { return x / 2 },
			expected: []float64{1.0, 2.0, 3.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.values.(type) {
			case []int:
				got := Map(v, tt.f.(func(int) int))
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Map() = %v, want %v", got, tt.expected)
				}
			case []string:
				got := Map(v, tt.f.(func(string) string))
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Map() = %v, want %v", got, tt.expected)
				}
			case []float64:
				got := Map(v, tt.f.(func(float64) float64))
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Map() = %v, want %v", got, tt.expected)
				}
			case []Person:
				got := Map(v, tt.f.(func(Person) Person))
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Map() = %v, want %v", got, tt.expected)
				}
			default:
				t.Errorf("Unsupported type for test values")
			}
		})
	}
}

func TestReduce(t *testing.T) {
	// 创建一个根目录结构用于路径拼接测试
	root := &Directory{Name: "root"}
	subfolder := &Directory{Name: "subfolder", Ancestors: []*Directory{root}}
	file := &Directory{Name: "file", Ancestors: []*Directory{root, subfolder}}

	// 定义表驱动测试用例
	tests := []struct {
		name     string
		values   interface{}
		f        interface{}
		initial  interface{}
		expected interface{}
	}{
		{
			name:   "Path concatenation",
			values: file.Ancestors,
			f: func(path string, dir *Directory) string {
				return path + "/" + dir.Name
			},
			initial:  "",
			expected: "/root/subfolder",
		},
		{
			name:   "Sum of integers",
			values: []int{1, 2, 3, 4, 5},
			f: func(accum int, v int) int {
				return accum + v
			},
			initial:  0,
			expected: 15,
		},
		{
			name:   "String concatenation",
			values: []string{"Go", "is", "awesome"},
			f: func(accum string, s string) string {
				return accum + " " + s
			},
			initial:  "",
			expected: " Go is awesome",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.values.(type) {
			case []*Directory:
				result := Reduce(v, tt.f.(func(string, *Directory) string), tt.initial.(string))
				if result != tt.expected.(string) {
					t.Errorf("Expected %v, but got %v", tt.expected, result)
				}
			case []int:
				result := Reduce(v, tt.f.(func(int, int) int), tt.initial.(int))
				if result != tt.expected.(int) {
					t.Errorf("Expected %v, but got %v", tt.expected, result)
				}
			case []string:
				result := Reduce(v, tt.f.(func(string, string) string), tt.initial.(string))
				if result != tt.expected.(string) {
					t.Errorf("Expected %v, but got %v", tt.expected, result)
				}
			default:
				t.Errorf("Unsupported test type %T", v)
			}
		})
	}
}

// Path 拼接函数，返回完整路径
func (d *Directory) Path() (string, error) {
	if d.IsRoot {
		return "/", nil
	}
	if d.Ancestors == nil {
		return "", gerrors.New("directory not found ancestors")
	}
	p := Reduce(d.Ancestors, func(path string, dir *Directory) string {
		return path + "/" + dir.Name
	}, "") + "/" + d.Name
	// hack: 去掉改逻辑
	if strings.HasPrefix(p, "//") {
		p = p[1:]
	}
	return p, nil
}

// 测试用例
var testCases = []struct {
	name     string
	dir      *Directory
	expected string
	err      error
}{
	{
		name: "Root Directory",
		dir: &Directory{
			IsRoot: true,
		},
		expected: "/",
		err:      nil,
	},
	{
		name: "Directory with Ancestors",
		dir: &Directory{
			Name: "dir3",
			Ancestors: []*Directory{
				{Name: "dir1"},
				{Name: "dir2"},
			},
		},
		expected: "/dir1/dir2/dir3",
		err:      nil,
	},
	{
		name: "Directory with No Ancestors",
		dir: &Directory{
			Name:      "dir1",
			Ancestors: nil,
		},
		expected: "",
		err:      gerrors.New("directory not found ancestors"),
	},
}

func TestDirectoryPath(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path, err := tc.dir.Path()
			if err != nil && tc.err == nil {
				t.Errorf("unexpected error: %v", err)
			}
			if err == nil && tc.err != nil {
				t.Errorf("expected error: %v, got nil", tc.err)
			}
			if path != tc.expected {
				t.Errorf("expected path: %s, got: %s", tc.expected, path)
			}
		})
	}
}
