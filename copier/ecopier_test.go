package ecopier

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- 定义用于测试的结构体 ---

// 嵌套结构体，用于测试指针字段
type Nested struct {
	Value string
}

// 复杂结构体，用于测试 gob 回退路径
// 它没有实现 Cloner 接口
type ComplexStruct struct {
	ID      int
	Name    string
	Tags    []string
	Attrs   map[string]any
	NestedP *Nested // 指针字段
}

// 可克隆结构体，用于测试 Clone() 优先路径
// 它实现了 Cloner 接口
type CloneableStruct struct {
	ComplexStruct     // 嵌套 ComplexStruct 以复用字段
	cloneCount    int // 用于测试 Clone 方法是否真的被调用
}

// 为 *CloneableStruct 实现 Cloner[*CloneableStruct] 接口
func (c *CloneableStruct) Clone() *CloneableStruct {
	if c == nil {
		return nil
	}
	// 手动执行深拷贝
	newC := &CloneableStruct{
		cloneCount: c.cloneCount + 1, // 增加计数器以证明此方法被调用
	}
	newC.ID = c.ID
	newC.Name = c.Name
	// 拷贝切片
	if c.Tags != nil {
		newC.Tags = make([]string, len(c.Tags))
		copy(newC.Tags, c.Tags)
	}
	// 拷贝 map
	if c.Attrs != nil {
		newC.Attrs = make(map[string]any)
		for k, v := range c.Attrs {
			newC.Attrs[k] = v // 假设 map 的值是基本类型
		}
	}
	// 拷贝指针字段
	if c.NestedP != nil {
		newC.NestedP = &Nested{Value: c.NestedP.Value}
	}

	return newC
}

// --- 测试主函数 ---

func TestDeepCopy(t *testing.T) {
	// 定义测试用例
	testCases := []struct {
		name          string                                              // 测试用例名称
		input         any                                                 // 输入数据
		assertionFunc func(t *testing.T, original, copied any, err error) // 断言函数
	}{
		{
			name:  "Gob path with struct value",
			input: ComplexStruct{ID: 1, Name: "test"},
			assertionFunc: func(t *testing.T, original, copied any, err error) {
				require.NoError(t, err)
				assert.Equal(t, original, copied)
			},
		},
		{
			name: "Gob path with struct pointer",
			input: &ComplexStruct{
				ID:   10,
				Name: "ptr test",
				Tags: []string{"a", "b"},
				Attrs: map[string]any{
					"key": "value",
				},
				NestedP: &Nested{Value: "nested"},
			},
			assertionFunc: func(t *testing.T, original, copied any, err error) {
				require.NoError(t, err)

				// 类型转换
				origPtr := original.(*ComplexStruct)
				copyPtr := copied.(*ComplexStruct)

				// 1. 值应该相等
				assert.Equal(t, origPtr, copyPtr)
				// 2. 但指针地址不应该相同，证明是新对象
				assert.NotSame(t, origPtr, copyPtr)
				// 3. 内部引用类型的地址也不应相同
				assert.NotSame(t, &origPtr.Tags[0], &copyPtr.Tags[0])
				assert.NotSame(t, origPtr.NestedP, copyPtr.NestedP)

				// 4. 修改原始对象，不应影响拷贝对象
				origPtr.Tags[0] = "z"
				origPtr.Attrs["key"] = "new_value"
				origPtr.NestedP.Value = "modified"

				assert.Equal(t, "a", copyPtr.Tags[0], "Copy's slice should not be modified")
				assert.Equal(t, "value", copyPtr.Attrs["key"], "Copy's map should not be modified")
				assert.Equal(t, "nested", copyPtr.NestedP.Value, "Copy's nested pointer value should not be modified")
			},
		},
		{
			name: "Clone path with struct pointer",
			input: &CloneableStruct{
				ComplexStruct: ComplexStruct{
					ID:      100,
					Name:    "cloneable",
					Tags:    []string{"c1", "c2"},
					NestedP: &Nested{Value: "clone_nested"},
				},
			},
			assertionFunc: func(t *testing.T, original, copied any, err error) {
				require.NoError(t, err)

				origPtr := original.(*CloneableStruct)
				copyPtr := copied.(*CloneableStruct)

				// 1. 验证 Clone() 方法被调用
				assert.Equal(t, 0, origPtr.cloneCount, "Original clone count should be 0")
				assert.Equal(t, 1, copyPtr.cloneCount, "Copy's clone count should be 1, proving Clone() was called")

				// 2. 基本值应该相等 (除了 cloneCount)
				assert.Equal(t, origPtr.ID, copyPtr.ID)
				assert.Equal(t, origPtr.Name, copyPtr.Name)

				// 3. 指针地址不应相同
				assert.NotSame(t, origPtr, copyPtr)
				assert.NotSame(t, origPtr.NestedP, copyPtr.NestedP)

				// 4. 修改原始对象，不应影响拷贝对象
				origPtr.Tags[0] = "c_mod"
				origPtr.NestedP.Value = "clone_mod"

				assert.Equal(t, "c1", copyPtr.Tags[0])
				assert.Equal(t, "clone_nested", copyPtr.NestedP.Value)
			},
		},
		{
			name:  "Nil pointer input",
			input: (*ComplexStruct)(nil),
			assertionFunc: func(t *testing.T, original, copied any, err error) {
				require.NoError(t, err)
				assert.Nil(t, copied, "Copy of a nil pointer should be nil")
				// 验证类型是否正确
				_, ok := copied.(*ComplexStruct)
				assert.True(t, ok, "The type of the copied nil pointer should be correct")
			},
		},
		{
			name:  "Basic type: string",
			input: "hello world",
			assertionFunc: func(t *testing.T, original, copied any, err error) {
				require.NoError(t, err)
				assert.Equal(t, "hello world", copied)
			},
		},
		{
			name:  "Basic type: slice of structs",
			input: []ComplexStruct{{ID: 1}, {ID: 2}},
			assertionFunc: func(t *testing.T, original, copied any, err error) {
				require.NoError(t, err)
				assert.Equal(t, original, copied)

				origSlice := original.([]ComplexStruct)
				copySlice := copied.([]ComplexStruct)

				// 虽然内容相同，但它们不应该是同一个切片的引用
				// 注意：Go中切片头是值传递，所以&origSlice[0]的地址可能相同
				// 验证深拷贝的最好方法是修改源并检查副本
				origSlice[0].ID = 99
				assert.Equal(t, 1, copySlice[0].ID, "Modifying original slice element should not affect copy")
			},
		},
		{
			name:  "Gob error path with un-encodable type (chan)",
			input: struct{ C chan int }{C: make(chan int)}, // gob 不能编码 channel
			assertionFunc: func(t *testing.T, original, copied any, err error) {
				require.Error(t, err, "Should return an error for un-encodable types")
				assert.Contains(t, err.Error(), "gob encode error", "Error message should indicate gob encoding failed")
				// 验证返回值是零值
				assert.Zero(t, copied)
			},
		},
	}

	// 循环执行所有测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 由于 DeepCopy 是泛型函数，我们不能直接传入 any
			// 需要一个 switch 来为每个具体类型调用它
			switch v := tc.input.(type) {
			case ComplexStruct:
				copied, err := DeepCopy(v)
				tc.assertionFunc(t, v, copied, err)
			case *ComplexStruct:
				copied, err := DeepCopy(v)
				tc.assertionFunc(t, v, copied, err)
			case *CloneableStruct:
				copied, err := DeepCopy(v)
				tc.assertionFunc(t, v, copied, err)
			case string:
				copied, err := DeepCopy(v)
				tc.assertionFunc(t, v, copied, err)
			case []ComplexStruct:
				copied, err := DeepCopy(v)
				tc.assertionFunc(t, v, copied, err)
			case struct{ C chan int }:
				copied, err := DeepCopy(v)
				tc.assertionFunc(t, v, copied, err)
			default:
				// t.Fatalf 会立即停止测试
				t.Fatalf("Unhandled type for test case: %s", tc.name)
			}
		})
	}
}
