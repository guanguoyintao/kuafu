package enumber

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSafeInt8ToInt16 测试 int8 到 int16 的转换
func TestSafeInt8ToInt16(t *testing.T) {
	assert.Equal(t, int16(127), SafeInt8ToInt16(127))
	assert.Equal(t, int16(-128), SafeInt8ToInt16(-128))
	assert.Equal(t, int16(0), SafeInt8ToInt16(0))
}

// TestSafeInt8ToInt32 测试 int8 到 int32 的转换
func TestSafeInt8ToInt32(t *testing.T) {
	assert.Equal(t, int32(127), SafeInt8ToInt32(127))
	assert.Equal(t, int32(-128), SafeInt8ToInt32(-128))
	assert.Equal(t, int32(0), SafeInt8ToInt32(0))
}

// TestSafeInt8ToInt64 测试 int8 到 int64 的转换
func TestSafeInt8ToInt64(t *testing.T) {
	assert.Equal(t, int64(127), SafeInt8ToInt64(127))
	assert.Equal(t, int64(-128), SafeInt8ToInt64(-128))
	assert.Equal(t, int64(0), SafeInt8ToInt64(0))
}

// TestSafeInt8ToUint8 测试 int8 到 uint8 的安全转换
func TestSafeInt8ToUint8(t *testing.T) {
	// 正常情况
	result, err := SafeInt8ToUint8(100)
	assert.NoError(t, err)
	assert.Equal(t, uint8(100), result)
	// 边界值
	result, err = SafeInt8ToUint8(127)
	assert.NoError(t, err)
	assert.Equal(t, uint8(127), result)
	result, err = SafeInt8ToUint8(0)
	assert.NoError(t, err)
	assert.Equal(t, uint8(0), result)
	// 负数情况
	_, err = SafeInt8ToUint8(-1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
}

// TestSafeInt32ToInt8 测试 int32 到 int8 的安全转换
func TestSafeInt32ToInt8(t *testing.T) {
	// 正常情况
	result, err := SafeInt32ToInt8(100)
	assert.NoError(t, err)
	assert.Equal(t, int8(100), result)
	// 边界值
	result, err = SafeInt32ToInt8(127)
	assert.NoError(t, err)
	assert.Equal(t, int8(127), result)
	result, err = SafeInt32ToInt8(-128)
	assert.NoError(t, err)
	assert.Equal(t, int8(-128), result)
	// 溢出情况
	_, err = SafeInt32ToInt8(128)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
	_, err = SafeInt32ToInt8(-129)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeInt32ToUint32 测试 int32 到 uint32 的安全转换
func TestSafeInt32ToUint32(t *testing.T) {
	// 正常情况
	result, err := SafeInt32ToUint32(100)
	assert.NoError(t, err)
	assert.Equal(t, uint32(100), result)
	// 边界值
	result, err = SafeInt32ToUint32(0)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), result)
	result, err = SafeInt32ToUint32(math.MaxInt32)
	assert.NoError(t, err)
	assert.Equal(t, uint32(math.MaxInt32), result)
	// 负数情况
	_, err = SafeInt32ToUint32(-1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
}

// TestSafeInt64ToInt32 测试 int64 到 int32 的安全转换
func TestSafeInt64ToInt32(t *testing.T) {
	// 正常情况
	result, err := SafeInt64ToInt32(100)
	assert.NoError(t, err)
	assert.Equal(t, int32(100), result)
	// 边界值
	result, err = SafeInt64ToInt32(math.MaxInt32)
	assert.NoError(t, err)
	assert.Equal(t, int32(math.MaxInt32), result)
	result, err = SafeInt64ToInt32(math.MinInt32)
	assert.NoError(t, err)
	assert.Equal(t, int32(math.MinInt32), result)
	// 溢出情况
	_, err = SafeInt64ToInt32(int64(math.MaxInt32) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
	_, err = SafeInt64ToInt32(int64(math.MinInt32) - 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeInt64ToUint32 测试 int64 到 uint32 的安全转换
func TestSafeInt64ToUint32(t *testing.T) {
	// 正常情况
	result, err := SafeInt64ToUint32(100)
	assert.NoError(t, err)
	assert.Equal(t, uint32(100), result)
	// 边界值
	result, err = SafeInt64ToUint32(0)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), result)
	result, err = SafeInt64ToUint32(math.MaxUint32)
	assert.NoError(t, err)
	assert.Equal(t, uint32(math.MaxUint32), result)
	// 负数情况
	_, err = SafeInt64ToUint32(-1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 溢出情况
	_, err = SafeInt64ToUint32(int64(math.MaxUint32) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeInt64ToUint64 测试 int64 到 uint64 的安全转换
func TestSafeInt64ToUint64(t *testing.T) {
	// 正常情况
	result, err := SafeInt64ToUint64(100)
	assert.NoError(t, err)
	assert.Equal(t, uint64(100), result)
	// 边界值
	result, err = SafeInt64ToUint64(0)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), result)
	result, err = SafeInt64ToUint64(math.MaxInt64)
	assert.NoError(t, err)
	assert.Equal(t, uint64(math.MaxInt64), result)
	// 负数情况
	_, err = SafeInt64ToUint64(-1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
}

// TestSafeUint32ToInt32 测试 uint32 到 int32 的安全转换
func TestSafeUint32ToInt32(t *testing.T) {
	// 正常情况
	result, err := SafeUint32ToInt32(100)
	assert.NoError(t, err)
	assert.Equal(t, int32(100), result)
	// 边界值
	result, err = SafeUint32ToInt32(0)
	assert.NoError(t, err)
	assert.Equal(t, int32(0), result)
	result, err = SafeUint32ToInt32(math.MaxInt32)
	assert.NoError(t, err)
	assert.Equal(t, int32(math.MaxInt32), result)
	// 溢出情况
	_, err = SafeUint32ToInt32(uint32(math.MaxInt32) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeUint64ToInt64 测试 uint64 到 int64 的安全转换
func TestSafeUint64ToInt64(t *testing.T) {
	// 正常情况
	result, err := SafeUint64ToInt64(100)
	assert.NoError(t, err)
	assert.Equal(t, int64(100), result)
	// 边界值
	result, err = SafeUint64ToInt64(0)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), result)
	result, err = SafeUint64ToInt64(math.MaxInt64)
	assert.NoError(t, err)
	assert.Equal(t, int64(math.MaxInt64), result)
	// 溢出情况
	_, err = SafeUint64ToInt64(uint64(math.MaxInt64) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeUint64ToUint32 测试 uint64 到 uint32 的安全转换
func TestSafeUint64ToUint32(t *testing.T) {
	// 正常情况
	result, err := SafeUint64ToUint32(100)
	assert.NoError(t, err)
	assert.Equal(t, uint32(100), result)
	// 边界值
	result, err = SafeUint64ToUint32(0)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), result)
	result, err = SafeUint64ToUint32(math.MaxUint32)
	assert.NoError(t, err)
	assert.Equal(t, uint32(math.MaxUint32), result)
	// 溢出情况
	_, err = SafeUint64ToUint32(uint64(math.MaxUint32) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeFloat64ToFloat32 测试 float64 到 float32 的安全转换
func TestSafeFloat64ToFloat32(t *testing.T) {
	// 正常情况
	result, err := SafeFloat64ToFloat32(3.14)
	assert.NoError(t, err)
	assert.Equal(t, float32(3.14), result)
	// 边界值
	result, err = SafeFloat64ToFloat32(0.0)
	assert.NoError(t, err)
	assert.Equal(t, float32(0.0), result)
	result, err = SafeFloat64ToFloat32(math.MaxFloat32)
	assert.NoError(t, err)
	assert.Equal(t, float32(math.MaxFloat32), result)
	// 无穷大情况
	_, err = SafeFloat64ToFloat32(math.Inf(1))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "infinity")
	// NaN情况
	_, err = SafeFloat64ToFloat32(math.NaN())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "NaN")
	// 溢出情况
	_, err = SafeFloat64ToFloat32(math.MaxFloat64)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeFloat64ToInt64 测试 float64 到 int64 的安全转换
func TestSafeFloat64ToInt64(t *testing.T) {
	// 正常情况
	result, err := SafeFloat64ToInt64(100.0)
	assert.NoError(t, err)
	assert.Equal(t, int64(100), result)
	// 边界值
	result, err = SafeFloat64ToInt64(0.0)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), result)
	result, err = SafeFloat64ToInt64(float64(math.MaxInt64))
	assert.NoError(t, err)
	assert.Equal(t, int64(math.MaxInt64), result)
	// 有小数部分的情况
	_, err = SafeFloat64ToInt64(3.14)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fractional part")
	// 无穷大情况
	_, err = SafeFloat64ToInt64(math.Inf(1))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "infinity")
	// NaN情况
	_, err = SafeFloat64ToInt64(math.NaN())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "NaN")
	// 溢出情况
	_, err = SafeFloat64ToInt64(math.MaxFloat64)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeIntToUint 测试 int 到 uint 的安全转换
func TestSafeIntToUint(t *testing.T) {
	// 正常情况
	result, err := SafeIntToUint(100)
	assert.NoError(t, err)
	assert.Equal(t, uint(100), result)
	// 边界值
	result, err = SafeIntToUint(0)
	assert.NoError(t, err)
	assert.Equal(t, uint(0), result)
	// 负数情况
	_, err = SafeIntToUint(-1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
}

// TestSafeUintToInt 测试 uint 到 int 的安全转换
func TestSafeUintToInt(t *testing.T) {
	// 正常情况
	result, err := SafeUintToInt(100)
	assert.NoError(t, err)
	assert.Equal(t, int(100), result)
	// 边界值
	result, err = SafeUintToInt(0)
	assert.NoError(t, err)
	assert.Equal(t, int(0), result)
	result, err = SafeUintToInt(uint(math.MaxInt))
	assert.NoError(t, err)
	assert.Equal(t, int(math.MaxInt), result)
	// 溢出情况
	_, err = SafeUintToInt(uint(math.MaxInt) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeUint8ToInt8 测试 uint8 到 int8 的安全转换
func TestSafeUint8ToInt8(t *testing.T) {
	// 正常情况
	result, err := SafeUint8ToInt8(100)
	assert.NoError(t, err)
	assert.Equal(t, int8(100), result)
	// 边界值
	result, err = SafeUint8ToInt8(0)
	assert.NoError(t, err)
	assert.Equal(t, int8(0), result)
	result, err = SafeUint8ToInt8(math.MaxInt8)
	assert.NoError(t, err)
	assert.Equal(t, int8(math.MaxInt8), result)
	// 溢出情况
	_, err = SafeUint8ToInt8(uint8(math.MaxInt8) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeFloat32ToInt32 测试 float32 到 int32 的安全转换
func TestSafeFloat32ToInt32(t *testing.T) {
	// 正常情况
	result, err := SafeFloat32ToInt32(100.0)
	assert.NoError(t, err)
	assert.Equal(t, int32(100), result)
	// 边界值
	result, err = SafeFloat32ToInt32(0.0)
	assert.NoError(t, err)
	assert.Equal(t, int32(0), result)
	// 有小数部分的情况
	_, err = SafeFloat32ToInt32(3.14)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fractional part")
	// 无穷大情况
	_, err = SafeFloat32ToInt32(float32(math.Inf(1)))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "infinity")
	// NaN情况
	_, err = SafeFloat32ToInt32(float32(math.NaN()))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "NaN")
}

// TestSafeFloat32ToFloat64 测试 float32 到 float64 的转换
func TestSafeFloat32ToFloat64(t *testing.T) {
	// 使用 float32 精确表示的值进行测试
	var f32 float32 = 3.14
	result := SafeFloat32ToFloat64(f32)
	assert.Equal(t, float64(f32), result)
	assert.Equal(t, float64(0.0), SafeFloat32ToFloat64(0.0))
	f32 = -3.14
	result = SafeFloat32ToFloat64(f32)
	assert.Equal(t, float64(f32), result)
}

// TestSafeIntToInt8 测试 int 到 int8 的安全转换
func TestSafeIntToInt8(t *testing.T) {
	// 正常情况
	result, err := SafeIntToInt8(100)
	assert.NoError(t, err)
	assert.Equal(t, int8(100), result)
	// 边界值
	result, err = SafeIntToInt8(127)
	assert.NoError(t, err)
	assert.Equal(t, int8(127), result)
	result, err = SafeIntToInt8(-128)
	assert.NoError(t, err)
	assert.Equal(t, int8(-128), result)
	// 溢出情况
	_, err = SafeIntToInt8(128)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
	_, err = SafeIntToInt8(-129)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeIntToInt16 测试 int 到 int16 的安全转换
func TestSafeIntToInt16(t *testing.T) {
	// 正常情况
	result, err := SafeIntToInt16(1000)
	assert.NoError(t, err)
	assert.Equal(t, int16(1000), result)
	// 边界值
	result, err = SafeIntToInt16(math.MaxInt16)
	assert.NoError(t, err)
	assert.Equal(t, int16(math.MaxInt16), result)
	result, err = SafeIntToInt16(math.MinInt16)
	assert.NoError(t, err)
	assert.Equal(t, int16(math.MinInt16), result)
	// 溢出情况
	_, err = SafeIntToInt16(int(math.MaxInt16) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
	_, err = SafeIntToInt16(int(math.MinInt16) - 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeIntToInt32 测试 int 到 int32 的安全转换
func TestSafeIntToInt32(t *testing.T) {
	// 正常情况
	result, err := SafeIntToInt32(100000)
	assert.NoError(t, err)
	assert.Equal(t, int32(100000), result)
	// 边界值
	result, err = SafeIntToInt32(math.MaxInt32)
	assert.NoError(t, err)
	assert.Equal(t, int32(math.MaxInt32), result)
	result, err = SafeIntToInt32(math.MinInt32)
	assert.NoError(t, err)
	assert.Equal(t, int32(math.MinInt32), result)
	// 在64位系统上测试溢出情况
	if math.MaxInt > math.MaxInt32 {
		_, err = SafeIntToInt32(int(math.MaxInt32) + 1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "overflow")
		_, err = SafeIntToInt32(int(math.MinInt32) - 1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "overflow")
	}
}

// TestSafeIntToInt64 测试 int 到 int64 的安全转换
func TestSafeIntToInt64(t *testing.T) {
	// 正常情况
	result := SafeIntToInt64(100000)
	assert.Equal(t, int64(100000), result)
	// 边界值
	result = SafeIntToInt64(math.MaxInt)
	assert.Equal(t, int64(math.MaxInt), result)
	result = SafeIntToInt64(math.MinInt)
	assert.Equal(t, int64(math.MinInt), result)
}

// TestSafeIntToUint8 测试 int 到 uint8 的安全转换
func TestSafeIntToUint8(t *testing.T) {
	// 正常情况
	result, err := SafeIntToUint8(100)
	assert.NoError(t, err)
	assert.Equal(t, uint8(100), result)
	// 边界值
	result, err = SafeIntToUint8(0)
	assert.NoError(t, err)
	assert.Equal(t, uint8(0), result)
	result, err = SafeIntToUint8(math.MaxUint8)
	assert.NoError(t, err)
	assert.Equal(t, uint8(math.MaxUint8), result)
	// 负数情况
	_, err = SafeIntToUint8(-1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 溢出情况
	_, err = SafeIntToUint8(int(math.MaxUint8) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeIntToUint16 测试 int 到 uint16 的安全转换
func TestSafeIntToUint16(t *testing.T) {
	// 正常情况
	result, err := SafeIntToUint16(1000)
	assert.NoError(t, err)
	assert.Equal(t, uint16(1000), result)
	// 边界值
	result, err = SafeIntToUint16(0)
	assert.NoError(t, err)
	assert.Equal(t, uint16(0), result)
	result, err = SafeIntToUint16(math.MaxUint16)
	assert.NoError(t, err)
	assert.Equal(t, uint16(math.MaxUint16), result)
	// 负数情况
	_, err = SafeIntToUint16(-1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 溢出情况
	_, err = SafeIntToUint16(int(math.MaxUint16) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeIntToUint32 测试 int 到 uint32 的安全转换
func TestSafeIntToUint32(t *testing.T) {
	// 正常情况
	result, err := SafeIntToUint32(100000)
	assert.NoError(t, err)
	assert.Equal(t, uint32(100000), result)
	// 边界值
	result, err = SafeIntToUint32(0)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), result)
	result, err = SafeIntToUint32(math.MaxUint32)
	assert.NoError(t, err)
	assert.Equal(t, uint32(math.MaxUint32), result)
	// 负数情况
	_, err = SafeIntToUint32(-1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 在64位系统上测试溢出情况
	if math.MaxInt > math.MaxUint32 {
		_, err = SafeIntToUint32(int(math.MaxUint32) + 1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "overflow")
	}
}

// TestSafeIntToUint64 测试 int 到 uint64 的安全转换
func TestSafeIntToUint64(t *testing.T) {
	// 正常情况
	result, err := SafeIntToUint64(100000)
	assert.NoError(t, err)
	assert.Equal(t, uint64(100000), result)
	// 边界值
	result, err = SafeIntToUint64(0)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), result)
	result, err = SafeIntToUint64(math.MaxInt)
	assert.NoError(t, err)
	assert.Equal(t, uint64(math.MaxInt), result)
	// 负数情况
	_, err = SafeIntToUint64(-1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
}

// TestSafeIntToFloat32 测试 int 到 float32 的转换
func TestSafeIntToFloat32(t *testing.T) {
	// 正常情况
	result := SafeIntToFloat32(100)
	assert.Equal(t, float32(100), result)
	// 边界值
	result = SafeIntToFloat32(0)
	assert.Equal(t, float32(0), result)
	result = SafeIntToFloat32(-100)
	assert.Equal(t, float32(-100), result)
	result = SafeIntToFloat32(math.MaxInt)
	assert.Equal(t, float32(math.MaxInt), result)
}

// TestSafeIntToFloat64 测试 int 到 float64 的转换
func TestSafeIntToFloat64(t *testing.T) {
	// 正常情况
	result := SafeIntToFloat64(100)
	assert.Equal(t, float64(100), result)
	// 边界值
	result = SafeIntToFloat64(0)
	assert.Equal(t, float64(0), result)
	result = SafeIntToFloat64(-100)
	assert.Equal(t, float64(-100), result)
	result = SafeIntToFloat64(math.MaxInt)
	assert.Equal(t, float64(math.MaxInt), result)
}

// TestSafeUintToInt8 测试 uint 到 int8 的安全转换
func TestSafeUintToInt8(t *testing.T) {
	// 正常情况
	result, err := SafeUintToInt8(100)
	assert.NoError(t, err)
	assert.Equal(t, int8(100), result)
	// 边界值
	result, err = SafeUintToInt8(0)
	assert.NoError(t, err)
	assert.Equal(t, int8(0), result)
	result, err = SafeUintToInt8(math.MaxInt8)
	assert.NoError(t, err)
	assert.Equal(t, int8(math.MaxInt8), result)
	// 溢出情况
	_, err = SafeUintToInt8(uint(math.MaxInt8) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeUintToInt16 测试 uint 到 int16 的安全转换
func TestSafeUintToInt16(t *testing.T) {
	// 正常情况
	result, err := SafeUintToInt16(1000)
	assert.NoError(t, err)
	assert.Equal(t, int16(1000), result)
	// 边界值
	result, err = SafeUintToInt16(0)
	assert.NoError(t, err)
	assert.Equal(t, int16(0), result)
	result, err = SafeUintToInt16(math.MaxInt16)
	assert.NoError(t, err)
	assert.Equal(t, int16(math.MaxInt16), result)
	// 溢出情况
	_, err = SafeUintToInt16(uint(math.MaxInt16) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeUintToInt32 测试 uint 到 int32 的安全转换
func TestSafeUintToInt32(t *testing.T) {
	// 正常情况
	result, err := SafeUintToInt32(100000)
	assert.NoError(t, err)
	assert.Equal(t, int32(100000), result)
	// 边界值
	result, err = SafeUintToInt32(0)
	assert.NoError(t, err)
	assert.Equal(t, int32(0), result)
	result, err = SafeUintToInt32(math.MaxInt32)
	assert.NoError(t, err)
	assert.Equal(t, int32(math.MaxInt32), result)
	// 溢出情况
	_, err = SafeUintToInt32(uint(math.MaxInt32) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeUintToInt64 测试 uint 到 int64 的安全转换
func TestSafeUintToInt64(t *testing.T) {
	// 正常情况
	result, err := SafeUintToInt64(100000)
	assert.NoError(t, err)
	assert.Equal(t, int64(100000), result)
	// 边界值
	result, err = SafeUintToInt64(0)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), result)
	result, err = SafeUintToInt64(math.MaxInt64)
	assert.NoError(t, err)
	assert.Equal(t, int64(math.MaxInt64), result)
	// 溢出情况
	_, err = SafeUintToInt64(uint(math.MaxInt64) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeUintToUint8 测试 uint 到 uint8 的安全转换
func TestSafeUintToUint8(t *testing.T) {
	// 正常情况
	result, err := SafeUintToUint8(100)
	assert.NoError(t, err)
	assert.Equal(t, uint8(100), result)
	// 边界值
	result, err = SafeUintToUint8(0)
	assert.NoError(t, err)
	assert.Equal(t, uint8(0), result)
	result, err = SafeUintToUint8(math.MaxUint8)
	assert.NoError(t, err)
	assert.Equal(t, uint8(math.MaxUint8), result)
	// 溢出情况
	_, err = SafeUintToUint8(uint(math.MaxUint8) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeUintToUint16 测试 uint 到 uint16 的安全转换
func TestSafeUintToUint16(t *testing.T) {
	// 正常情况
	result, err := SafeUintToUint16(1000)
	assert.NoError(t, err)
	assert.Equal(t, uint16(1000), result)
	// 边界值
	result, err = SafeUintToUint16(0)
	assert.NoError(t, err)
	assert.Equal(t, uint16(0), result)
	result, err = SafeUintToUint16(math.MaxUint16)
	assert.NoError(t, err)
	assert.Equal(t, uint16(math.MaxUint16), result)
	// 溢出情况
	_, err = SafeUintToUint16(uint(math.MaxUint16) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeUintToUint32 测试 uint 到 uint32 的安全转换
func TestSafeUintToUint32(t *testing.T) {
	// 正常情况
	result, err := SafeUintToUint32(100000)
	assert.NoError(t, err)
	assert.Equal(t, uint32(100000), result)
	// 边界值
	result, err = SafeUintToUint32(0)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), result)
	result, err = SafeUintToUint32(math.MaxUint32)
	assert.NoError(t, err)
	assert.Equal(t, uint32(math.MaxUint32), result)
	// 在64位系统上测试溢出情况
	if math.MaxUint > math.MaxUint32 {
		_, err = SafeUintToUint32(uint(math.MaxUint32) + 1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "overflow")
	}
}

// TestSafeUintToUint64 测试 uint 到 uint64 的安全转换
func TestSafeUintToUint64(t *testing.T) {
	// 正常情况
	result := SafeUintToUint64(100000)
	assert.Equal(t, uint64(100000), result)
	// 边界值
	result = SafeUintToUint64(0)
	assert.Equal(t, uint64(0), result)
	result = SafeUintToUint64(math.MaxUint)
	assert.Equal(t, uint64(math.MaxUint), result)
}

// TestSafeUintToFloat32 测试 uint 到 float32 的转换
func TestSafeUintToFloat32(t *testing.T) {
	// 正常情况
	result := SafeUintToFloat32(100)
	assert.Equal(t, float32(100), result)
	// 边界值
	result = SafeUintToFloat32(0)
	assert.Equal(t, float32(0), result)
	result = SafeUintToFloat32(math.MaxUint)
	assert.Equal(t, float32(math.MaxUint), result)
}

// TestSafeUintToFloat64 测试 uint 到 float64 的转换
func TestSafeUintToFloat64(t *testing.T) {
	// 正常情况
	result := SafeUintToFloat64(100)
	assert.Equal(t, float64(100), result)
	// 边界值
	result = SafeUintToFloat64(0)
	assert.Equal(t, float64(0), result)
	result = SafeUintToFloat64(math.MaxUint)
	assert.Equal(t, float64(math.MaxUint), result)
}

// TestSafeFloat32ToUint8 测试 float32 到 uint8 的安全转换
func TestSafeFloat32ToUint8(t *testing.T) {
	// 正常情况
	result, err := SafeFloat32ToUint8(100.0)
	assert.NoError(t, err)
	assert.Equal(t, uint8(100), result)
	// 边界值
	result, err = SafeFloat32ToUint8(0.0)
	assert.NoError(t, err)
	assert.Equal(t, uint8(0), result)
	result, err = SafeFloat32ToUint8(float32(math.MaxUint8))
	assert.NoError(t, err)
	assert.Equal(t, uint8(math.MaxUint8), result)
	// 负数情况
	_, err = SafeFloat32ToUint8(-1.0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 有小数部分的情况
	_, err = SafeFloat32ToUint8(3.14)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fractional part")
	// 无穷大情况
	_, err = SafeFloat32ToUint8(float32(math.Inf(1)))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "infinity")
	// NaN情况
	_, err = SafeFloat32ToUint8(float32(math.NaN()))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "NaN")
	// 溢出情况
	_, err = SafeFloat32ToUint8(float32(math.MaxUint8) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeFloat32ToUint16 测试 float32 到 uint16 的安全转换
func TestSafeFloat32ToUint16(t *testing.T) {
	// 正常情况
	result, err := SafeFloat32ToUint16(1000.0)
	assert.NoError(t, err)
	assert.Equal(t, uint16(1000), result)
	// 边界值
	result, err = SafeFloat32ToUint16(0.0)
	assert.NoError(t, err)
	assert.Equal(t, uint16(0), result)
	result, err = SafeFloat32ToUint16(float32(math.MaxUint16))
	assert.NoError(t, err)
	assert.Equal(t, uint16(math.MaxUint16), result)
	// 负数情况
	_, err = SafeFloat32ToUint16(-1.0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 有小数部分的情况
	_, err = SafeFloat32ToUint16(3.14)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fractional part")
	// 溢出情况
	_, err = SafeFloat32ToUint16(float32(math.MaxUint16) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeFloat32ToUint32 测试 float32 到 uint32 的安全转换
func TestSafeFloat32ToUint32(t *testing.T) {
	// 正常情况
	result, err := SafeFloat32ToUint32(100000.0)
	assert.NoError(t, err)
	assert.Equal(t, uint32(100000), result)
	// 边界值
	result, err = SafeFloat32ToUint32(0.0)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), result)
	// 负数情况
	_, err = SafeFloat32ToUint32(-1.0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 有小数部分的情况
	_, err = SafeFloat32ToUint32(3.14)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fractional part")
}

// TestSafeFloat32ToUint64 测试 float32 到 uint64 的安全转换
func TestSafeFloat32ToUint64(t *testing.T) {
	// 正常情况
	result, err := SafeFloat32ToUint64(100000.0)
	assert.NoError(t, err)
	assert.Equal(t, uint64(100000), result)
	// 边界值
	result, err = SafeFloat32ToUint64(0.0)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), result)
	// 负数情况
	_, err = SafeFloat32ToUint64(-1.0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 有小数部分的情况
	_, err = SafeFloat32ToUint64(3.14)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fractional part")
}

// TestSafeFloat64ToUint8 测试 float64 到 uint8 的安全转换
func TestSafeFloat64ToUint8(t *testing.T) {
	// 正常情况
	result, err := SafeFloat64ToUint8(100.0)
	assert.NoError(t, err)
	assert.Equal(t, uint8(100), result)
	// 边界值
	result, err = SafeFloat64ToUint8(0.0)
	assert.NoError(t, err)
	assert.Equal(t, uint8(0), result)
	result, err = SafeFloat64ToUint8(float64(math.MaxUint8))
	assert.NoError(t, err)
	assert.Equal(t, uint8(math.MaxUint8), result)
	// 负数情况
	_, err = SafeFloat64ToUint8(-1.0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 有小数部分的情况
	_, err = SafeFloat64ToUint8(3.14)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fractional part")
	// 无穷大情况
	_, err = SafeFloat64ToUint8(math.Inf(1))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "infinity")
	// NaN情况
	_, err = SafeFloat64ToUint8(math.NaN())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "NaN")
	// 溢出情况
	_, err = SafeFloat64ToUint8(float64(math.MaxUint8) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeFloat64ToUint16 测试 float64 到 uint16 的安全转换
func TestSafeFloat64ToUint16(t *testing.T) {
	// 正常情况
	result, err := SafeFloat64ToUint16(1000.0)
	assert.NoError(t, err)
	assert.Equal(t, uint16(1000), result)
	// 边界值
	result, err = SafeFloat64ToUint16(0.0)
	assert.NoError(t, err)
	assert.Equal(t, uint16(0), result)
	result, err = SafeFloat64ToUint16(float64(math.MaxUint16))
	assert.NoError(t, err)
	assert.Equal(t, uint16(math.MaxUint16), result)
	// 负数情况
	_, err = SafeFloat64ToUint16(-1.0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 有小数部分的情况
	_, err = SafeFloat64ToUint16(3.14)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fractional part")
	// 溢出情况
	_, err = SafeFloat64ToUint16(float64(math.MaxUint16) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeFloat64ToUint32 测试 float64 到 uint32 的安全转换
func TestSafeFloat64ToUint32(t *testing.T) {
	// 正常情况
	result, err := SafeFloat64ToUint32(100000.0)
	assert.NoError(t, err)
	assert.Equal(t, uint32(100000), result)
	// 边界值
	result, err = SafeFloat64ToUint32(0.0)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), result)
	result, err = SafeFloat64ToUint32(float64(math.MaxUint32))
	assert.NoError(t, err)
	assert.Equal(t, uint32(math.MaxUint32), result)
	// 负数情况
	_, err = SafeFloat64ToUint32(-1.0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 有小数部分的情况
	_, err = SafeFloat64ToUint32(3.14)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fractional part")
	// 溢出情况
	_, err = SafeFloat64ToUint32(float64(math.MaxUint32) + 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overflow")
}

// TestSafeFloat64ToUint64 测试 float64 到 uint64 的安全转换
func TestSafeFloat64ToUint64(t *testing.T) {
	// 正常情况
	result, err := SafeFloat64ToUint64(100000.0)
	assert.NoError(t, err)
	assert.Equal(t, uint64(100000), result)
	// 边界值
	result, err = SafeFloat64ToUint64(0.0)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), result)
	// 负数情况
	_, err = SafeFloat64ToUint64(-1.0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "negative")
	// 有小数部分的情况
	_, err = SafeFloat64ToUint64(3.14)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fractional part")
}