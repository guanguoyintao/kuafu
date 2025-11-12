package enumber

import (
	"fmt"
	"math"
)

// SafeInt8ToInt16 安全地将 int8 转换为 int16
func SafeInt8ToInt16(value int8) int16 {
	return int16(value)
}

// SafeInt8ToInt32 安全地将 int8 转换为 int32
func SafeInt8ToInt32(value int8) int32 {
	return int32(value)
}

// SafeInt8ToInt64 安全地将 int8 转换为 int64
func SafeInt8ToInt64(value int8) int64 {
	return int64(value)
}

// SafeInt8ToUint8 安全地将 int8 转换为 uint8，避免负数
func SafeInt8ToUint8(value int8) (uint8, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint8", value)
	}
	return uint8(value), nil
}

// SafeInt8ToUint16 安全地将 int8 转换为 uint16，避免负数
func SafeInt8ToUint16(value int8) (uint16, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint16", value)
	}
	return uint16(value), nil
}

// SafeInt8ToUint32 安全地将 int8 转换为 uint32，避免负数
func SafeInt8ToUint32(value int8) (uint32, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint32", value)
	}
	return uint32(value), nil
}

// SafeInt8ToUint64 安全地将 int8 转换为 uint64，避免负数
func SafeInt8ToUint64(value int8) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint64", value)
	}
	return uint64(value), nil
}

// SafeInt16ToInt8 安全地将 int16 转换为 int8，避免溢出
func SafeInt16ToInt8(value int16) (int8, error) {
	if value < math.MinInt8 || value > math.MaxInt8 {
		return 0, fmt.Errorf("value %d overflow, int8 range is [%d, %d]", value, math.MinInt8, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeInt16ToInt32 安全地将 int16 转换为 int32
func SafeInt16ToInt32(value int16) int32 {
	return int32(value)
}

// SafeInt16ToInt64 安全地将 int16 转换为 int64
func SafeInt16ToInt64(value int16) int64 {
	return int64(value)
}

// SafeInt16ToUint8 安全地将 int16 转换为 uint8，避免负数和溢出
func SafeInt16ToUint8(value int16) (uint8, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint8", value)
	}
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("value %d overflow, uint8 max value is %d", value, math.MaxUint8)
	}
	return uint8(value), nil
}

// SafeInt16ToUint16 安全地将 int16 转换为 uint16，避免负数
func SafeInt16ToUint16(value int16) (uint16, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint16", value)
	}
	return uint16(value), nil
}

// SafeInt16ToUint32 安全地将 int16 转换为 uint32，避免负数
func SafeInt16ToUint32(value int16) (uint32, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint32", value)
	}
	return uint32(value), nil
}

// SafeInt16ToUint64 安全地将 int16 转换为 uint64，避免负数
func SafeInt16ToUint64(value int16) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint64", value)
	}
	return uint64(value), nil
}

// SafeInt32ToInt8 安全地将 int32 转换为 int8，避免溢出
func SafeInt32ToInt8(value int32) (int8, error) {
	if value < math.MinInt8 || value > math.MaxInt8 {
		return 0, fmt.Errorf("value %d overflow, int8 range is [%d, %d]", value, math.MinInt8, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeInt32ToInt16 安全地将 int32 转换为 int16，避免溢出
func SafeInt32ToInt16(value int32) (int16, error) {
	if value < math.MinInt16 || value > math.MaxInt16 {
		return 0, fmt.Errorf("value %d overflow, int16 range is [%d, %d]", value, math.MinInt16, math.MaxInt16)
	}
	return int16(value), nil
}

// SafeInt32ToInt64 安全地将 int32 转换为 int64
func SafeInt32ToInt64(value int32) int64 {
	return int64(value)
}

// SafeInt32ToUint8 安全地将 int32 转换为 uint8，避免负数和溢出
func SafeInt32ToUint8(value int32) (uint8, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint8", value)
	}
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("value %d overflow, uint8 max value is %d", value, math.MaxUint8)
	}
	return uint8(value), nil
}

// SafeInt32ToUint16 安全地将 int32 转换为 uint16，避免负数和溢出
func SafeInt32ToUint16(value int32) (uint16, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint16", value)
	}
	if value > math.MaxUint16 {
		return 0, fmt.Errorf("value %d overflow, uint16 max value is %d", value, math.MaxUint16)
	}
	return uint16(value), nil
}

// SafeInt32ToUint32 安全地将 int32 转换为 uint32，避免负数
func SafeInt32ToUint32(value int32) (uint32, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint32", value)
	}
	return uint32(value), nil
}

// SafeInt32ToUint64 安全地将 int32 转换为 uint64，避免负数
func SafeInt32ToUint64(value int32) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint64", value)
	}
	return uint64(value), nil
}

// SafeInt64ToInt8 安全地将 int64 转换为 int8，避免溢出
func SafeInt64ToInt8(value int64) (int8, error) {
	if value < math.MinInt8 || value > math.MaxInt8 {
		return 0, fmt.Errorf("value %d overflow, int8 range is [%d, %d]", value, math.MinInt8, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeInt64ToInt16 安全地将 int64 转换为 int16，避免溢出
func SafeInt64ToInt16(value int64) (int16, error) {
	if value < math.MinInt16 || value > math.MaxInt16 {
		return 0, fmt.Errorf("value %d overflow, int16 range is [%d, %d]", value, math.MinInt16, math.MaxInt16)
	}
	return int16(value), nil
}

// SafeInt64ToInt32 安全地将 int64 转换为 int32，避免溢出
func SafeInt64ToInt32(value int64) (int32, error) {
	if value < math.MinInt32 || value > math.MaxInt32 {
		return 0, fmt.Errorf("value %d overflow, int32 range is [%d, %d]", value, math.MinInt32, math.MaxInt32)
	}
	return int32(value), nil
}

// SafeInt64ToUint8 安全地将 int64 转换为 uint8，避免负数和溢出
func SafeInt64ToUint8(value int64) (uint8, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint8", value)
	}
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("value %d overflow, uint8 max value is %d", value, math.MaxUint8)
	}
	return uint8(value), nil
}

// SafeInt64ToUint16 安全地将 int64 转换为 uint16，避免负数和溢出
func SafeInt64ToUint16(value int64) (uint16, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint16", value)
	}
	if value > math.MaxUint16 {
		return 0, fmt.Errorf("value %d overflow, uint16 max value is %d", value, math.MaxUint16)
	}
	return uint16(value), nil
}

// SafeInt64ToUint32 安全地将 int64 转换为 uint32，避免溢出
func SafeInt64ToUint32(value int64) (uint32, error) {
	// 检查负数
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint32", value)
	}
	// 检查溢出
	if value > math.MaxUint32 {
		return 0, fmt.Errorf("value %d overflow, uint32 max value is %d", value, math.MaxUint32)
	}
	return uint32(value), nil
}

// SafeInt64ToUint64 安全地将 int64 转换为 uint64，避免负数
func SafeInt64ToUint64(value int64) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint64", value)
	}
	return uint64(value), nil
}

// SafeUint8ToInt8 安全地将 uint8 转换为 int8，避免溢出
func SafeUint8ToInt8(value uint8) (int8, error) {
	if value > math.MaxInt8 {
		return 0, fmt.Errorf("value %d overflow, int8 max value is %d", value, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeUint8ToInt16 安全地将 uint8 转换为 int16
func SafeUint8ToInt16(value uint8) int16 {
	return int16(value)
}

// SafeUint8ToInt32 安全地将 uint8 转换为 int32
func SafeUint8ToInt32(value uint8) int32 {
	return int32(value)
}

// SafeUint8ToInt64 安全地将 uint8 转换为 int64
func SafeUint8ToInt64(value uint8) int64 {
	return int64(value)
}

// SafeUint8ToUint16 安全地将 uint8 转换为 uint16
func SafeUint8ToUint16(value uint8) uint16 {
	return uint16(value)
}

// SafeUint8ToUint32 安全地将 uint8 转换为 uint32
func SafeUint8ToUint32(value uint8) uint32 {
	return uint32(value)
}

// SafeUint8ToUint64 安全地将 uint8 转换为 uint64
func SafeUint8ToUint64(value uint8) uint64 {
	return uint64(value)
}

// SafeUint16ToInt8 安全地将 uint16 转换为 int8，避免溢出
func SafeUint16ToInt8(value uint16) (int8, error) {
	if value > math.MaxInt8 {
		return 0, fmt.Errorf("value %d overflow, int8 max value is %d", value, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeUint16ToInt16 安全地将 uint16 转换为 int16，避免溢出
func SafeUint16ToInt16(value uint16) (int16, error) {
	if value > math.MaxInt16 {
		return 0, fmt.Errorf("value %d overflow, int16 max value is %d", value, math.MaxInt16)
	}
	return int16(value), nil
}

// SafeUint16ToInt32 安全地将 uint16 转换为 int32
func SafeUint16ToInt32(value uint16) int32 {
	return int32(value)
}

// SafeUint16ToInt64 安全地将 uint16 转换为 int64
func SafeUint16ToInt64(value uint16) int64 {
	return int64(value)
}

// SafeUint16ToUint8 安全地将 uint16 转换为 uint8，避免溢出
func SafeUint16ToUint8(value uint16) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("value %d overflow, uint8 max value is %d", value, math.MaxUint8)
	}
	return uint8(value), nil
}

// SafeUint16ToUint32 安全地将 uint16 转换为 uint32
func SafeUint16ToUint32(value uint16) uint32 {
	return uint32(value)
}

// SafeUint16ToUint64 安全地将 uint16 转换为 uint64
func SafeUint16ToUint64(value uint16) uint64 {
	return uint64(value)
}

// SafeUint32ToInt8 安全地将 uint32 转换为 int8，避免溢出
func SafeUint32ToInt8(value uint32) (int8, error) {
	if value > math.MaxInt8 {
		return 0, fmt.Errorf("value %d overflow, int8 max value is %d", value, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeUint32ToInt16 安全地将 uint32 转换为 int16，避免溢出
func SafeUint32ToInt16(value uint32) (int16, error) {
	if value > math.MaxInt16 {
		return 0, fmt.Errorf("value %d overflow, int16 max value is %d", value, math.MaxInt16)
	}
	return int16(value), nil
}

// SafeUint32ToInt32 安全地将 uint32 转换为 int32，避免溢出
func SafeUint32ToInt32(value uint32) (int32, error) {
	if value > math.MaxInt32 {
		return 0, fmt.Errorf("value %d overflow, int32 max value is %d", value, math.MaxInt32)
	}
	return int32(value), nil
}

// SafeUint32ToInt64 安全地将 uint32 转换为 int64
func SafeUint32ToInt64(value uint32) int64 {
	return int64(value)
}

// SafeUint32ToUint8 安全地将 uint32 转换为 uint8，避免溢出
func SafeUint32ToUint8(value uint32) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("value %d overflow, uint8 max value is %d", value, math.MaxUint8)
	}
	return uint8(value), nil
}

// SafeUint32ToUint16 安全地将 uint32 转换为 uint16，避免溢出
func SafeUint32ToUint16(value uint32) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, fmt.Errorf("value %d overflow, uint16 max value is %d", value, math.MaxUint16)
	}
	return uint16(value), nil
}

// SafeUint32ToUint64 安全地将 uint32 转换为 uint64
func SafeUint32ToUint64(value uint32) uint64 {
	return uint64(value)
}

// SafeUint64ToInt8 安全地将 uint64 转换为 int8，避免溢出
func SafeUint64ToInt8(value uint64) (int8, error) {
	if value > math.MaxInt8 {
		return 0, fmt.Errorf("value %d overflow, int8 max value is %d", value, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeUint64ToInt16 安全地将 uint64 转换为 int16，避免溢出
func SafeUint64ToInt16(value uint64) (int16, error) {
	if value > math.MaxInt16 {
		return 0, fmt.Errorf("value %d overflow, int16 max value is %d", value, math.MaxInt16)
	}
	return int16(value), nil
}

// SafeUint64ToInt32 安全地将 uint64 转换为 int32，避免溢出
func SafeUint64ToInt32(value uint64) (int32, error) {
	if value > math.MaxInt32 {
		return 0, fmt.Errorf("value %d overflow, int32 max value is %d", value, math.MaxInt32)
	}
	return int32(value), nil
}

// SafeUint64ToInt64 安全地将 uint64 转换为 int64，避免溢出
func SafeUint64ToInt64(value uint64) (int64, error) {
	if value > math.MaxInt64 {
		return 0, fmt.Errorf("value %d overflow, int64 max value is %d", value, math.MaxInt64)
	}
	return int64(value), nil
}

// SafeUint64ToUint8 安全地将 uint64 转换为 uint8，避免溢出
func SafeUint64ToUint8(value uint64) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("value %d overflow, uint8 max value is %d", value, math.MaxUint8)
	}
	return uint8(value), nil
}

// SafeUint64ToUint16 安全地将 uint64 转换为 uint16，避免溢出
func SafeUint64ToUint16(value uint64) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, fmt.Errorf("value %d overflow, uint16 max value is %d", value, math.MaxUint16)
	}
	return uint16(value), nil
}

// SafeUint64ToUint32 安全地将 uint64 转换为 uint32，避免溢出
func SafeUint64ToUint32(value uint64) (uint32, error) {
	if value > math.MaxUint32 {
		return 0, fmt.Errorf("value %d overflow, uint32 max value is %d", value, math.MaxUint32)
	}
	return uint32(value), nil
}

// SafeFloat32ToInt8 安全地将 float32 转换为 int8，避免溢出和精度丢失
func SafeFloat32ToInt8(value float32) (int8, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to int8", value)
	}
	// 检查是否有小数部分
	if value != float32(math.Trunc(float64(value))) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to int8 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, fmt.Errorf("value %f overflow, int8 range is [%d, %d]", value, math.MinInt8, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeFloat32ToInt16 安全地将 float32 转换为 int16，避免溢出和精度丢失
func SafeFloat32ToInt16(value float32) (int16, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to int16", value)
	}
	// 检查是否有小数部分
	if value != float32(math.Trunc(float64(value))) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to int16 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, fmt.Errorf("value %f overflow, int16 range is [%d, %d]", value, math.MinInt16, math.MaxInt16)
	}
	return int16(value), nil
}

// SafeFloat32ToInt32 安全地将 float32 转换为 int32，避免溢出和精度丢失
func SafeFloat32ToInt32(value float32) (int32, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to int32", value)
	}
	// 检查是否有小数部分
	if value != float32(math.Trunc(float64(value))) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to int32 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, fmt.Errorf("value %f overflow, int32 range is [%d, %d]", value, math.MinInt32, math.MaxInt32)
	}
	return int32(value), nil
}

// SafeFloat32ToInt64 安全地将 float32 转换为 int64，避免溢出和精度丢失
func SafeFloat32ToInt64(value float32) (int64, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to int64", value)
	}
	// 检查是否有小数部分
	if value != float32(math.Trunc(float64(value))) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to int64 without precision loss", value)
	}
	// 检查溢出
	if float64(value) > math.MaxInt64 || float64(value) < math.MinInt64 {
		return 0, fmt.Errorf("value %f overflow, int64 range is [%d, %d]", value, math.MinInt64, math.MaxInt64)
	}
	return int64(value), nil
}

// SafeFloat32ToFloat64 安全地将 float32 转换为 float64
func SafeFloat32ToFloat64(value float32) float64 {
	return float64(value)
}

// SafeFloat64ToFloat32 安全地将 float64 转换为 float32，避免溢出
func SafeFloat64ToFloat32(value float64) (float32, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to float32", value)
	}
	// 检查溢出
	if value > math.MaxFloat32 || value < -math.MaxFloat32 {
		return 0, fmt.Errorf("value %f overflow, float32 range is [%f, %f]", value, -math.MaxFloat32, math.MaxFloat32)
	}
	return float32(value), nil
}

// SafeFloat64ToInt8 安全地将 float64 转换为 int8，避免溢出和精度丢失
func SafeFloat64ToInt8(value float64) (int8, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to int8", value)
	}
	// 检查是否有小数部分
	if value != math.Trunc(value) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to int8 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxInt8 || value < math.MinInt8 {
		return 0, fmt.Errorf("value %f overflow, int8 range is [%d, %d]", value, math.MinInt8, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeFloat64ToInt16 安全地将 float64 转换为 int16，避免溢出和精度丢失
func SafeFloat64ToInt16(value float64) (int16, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to int16", value)
	}
	// 检查是否有小数部分
	if value != math.Trunc(value) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to int16 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxInt16 || value < math.MinInt16 {
		return 0, fmt.Errorf("value %f overflow, int16 range is [%d, %d]", value, math.MinInt16, math.MaxInt16)
	}
	return int16(value), nil
}

// SafeFloat64ToInt32 安全地将 float64 转换为 int32，避免溢出和精度丢失
func SafeFloat64ToInt32(value float64) (int32, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to int32", value)
	}
	// 检查是否有小数部分
	if value != math.Trunc(value) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to int32 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, fmt.Errorf("value %f overflow, int32 range is [%d, %d]", value, math.MinInt32, math.MaxInt32)
	}
	return int32(value), nil
}

// SafeFloat64ToInt64 安全地将 float64 转换为 int64，避免溢出和精度丢失
func SafeFloat64ToInt64(value float64) (int64, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to int64", value)
	}
	// 检查是否有小数部分
	if value != math.Trunc(value) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to int64 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxInt64 || value < math.MinInt64 {
		return 0, fmt.Errorf("value %f overflow, int64 range is [%d, %d]", value, math.MinInt64, math.MaxInt64)
	}
	return int64(value), nil
}

// SafeIntToUint 安全地将 int 转换为 uint，避免负数
func SafeIntToUint(value int) (uint, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint", value)
	}
	return uint(value), nil
}

// SafeUintToInt 安全地将 uint 转换为 int，避免溢出
func SafeUintToInt(value uint) (int, error) {
	// 在64位系统上，int和uint都是64位，但int的最大值小于uint的最大值
	if value > uint(math.MaxInt) {
		return 0, fmt.Errorf("value %d overflow, int max value is %d", value, math.MaxInt)
	}
	return int(value), nil
}

// SafeIntToInt8 安全地将 int 转换为 int8，避免溢出
func SafeIntToInt8(value int) (int8, error) {
	if value < math.MinInt8 || value > math.MaxInt8 {
		return 0, fmt.Errorf("value %d overflow, int8 range is [%d, %d]", value, math.MinInt8, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeIntToInt16 安全地将 int 转换为 int16，避免溢出
func SafeIntToInt16(value int) (int16, error) {
	if value < math.MinInt16 || value > math.MaxInt16 {
		return 0, fmt.Errorf("value %d overflow, int16 range is [%d, %d]", value, math.MinInt16, math.MaxInt16)
	}
	return int16(value), nil
}

// SafeIntToInt32 安全地将 int 转换为 int32，避免溢出
func SafeIntToInt32(value int) (int32, error) {
	if value < math.MinInt32 || value > math.MaxInt32 {
		return 0, fmt.Errorf("value %d overflow, int32 range is [%d, %d]", value, math.MinInt32, math.MaxInt32)
	}
	return int32(value), nil
}

// SafeIntToInt64 安全地将 int 转换为 int64
func SafeIntToInt64(value int) int64 {
	return int64(value)
}

// SafeIntToUint8 安全地将 int 转换为 uint8，避免负数和溢出
func SafeIntToUint8(value int) (uint8, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint8", value)
	}
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("value %d overflow, uint8 max value is %d", value, math.MaxUint8)
	}
	return uint8(value), nil
}

// SafeIntToUint16 安全地将 int 转换为 uint16，避免负数和溢出
func SafeIntToUint16(value int) (uint16, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint16", value)
	}
	if value > math.MaxUint16 {
		return 0, fmt.Errorf("value %d overflow, uint16 max value is %d", value, math.MaxUint16)
	}
	return uint16(value), nil
}

// SafeIntToUint32 安全地将 int 转换为 uint32，避免负数和溢出
func SafeIntToUint32(value int) (uint32, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint32", value)
	}
	if value > math.MaxUint32 {
		return 0, fmt.Errorf("value %d overflow, uint32 max value is %d", value, math.MaxUint32)
	}
	return uint32(value), nil
}

// SafeIntToUint64 安全地将 int 转换为 uint64，避免负数
func SafeIntToUint64(value int) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("value %d is negative, cannot convert to uint64", value)
	}
	return uint64(value), nil
}

// SafeIntToFloat32 安全地将 int 转换为 float32
func SafeIntToFloat32(value int) float32 {
	return float32(value)
}

// SafeIntToFloat64 安全地将 int 转换为 float64
func SafeIntToFloat64(value int) float64 {
	return float64(value)
}

// SafeUintToInt8 安全地将 uint 转换为 int8，避免溢出
func SafeUintToInt8(value uint) (int8, error) {
	if value > math.MaxInt8 {
		return 0, fmt.Errorf("value %d overflow, int8 max value is %d", value, math.MaxInt8)
	}
	return int8(value), nil
}

// SafeUintToInt16 安全地将 uint 转换为 int16，避免溢出
func SafeUintToInt16(value uint) (int16, error) {
	if value > math.MaxInt16 {
		return 0, fmt.Errorf("value %d overflow, int16 max value is %d", value, math.MaxInt16)
	}
	return int16(value), nil
}

// SafeUintToInt32 安全地将 uint 转换为 int32，避免溢出
func SafeUintToInt32(value uint) (int32, error) {
	if value > math.MaxInt32 {
		return 0, fmt.Errorf("value %d overflow, int32 max value is %d", value, math.MaxInt32)
	}
	return int32(value), nil
}

// SafeUintToInt64 安全地将 uint 转换为 int64，避免溢出
func SafeUintToInt64(value uint) (int64, error) {
	if value > math.MaxInt64 {
		return 0, fmt.Errorf("value %d overflow, int64 max value is %d", value, math.MaxInt64)
	}
	return int64(value), nil
}

// SafeUintToUint8 安全地将 uint 转换为 uint8，避免溢出
func SafeUintToUint8(value uint) (uint8, error) {
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("value %d overflow, uint8 max value is %d", value, math.MaxUint8)
	}
	return uint8(value), nil
}

// SafeUintToUint16 安全地将 uint 转换为 uint16，避免溢出
func SafeUintToUint16(value uint) (uint16, error) {
	if value > math.MaxUint16 {
		return 0, fmt.Errorf("value %d overflow, uint16 max value is %d", value, math.MaxUint16)
	}
	return uint16(value), nil
}

// SafeUintToUint32 安全地将 uint 转换为 uint32，避免溢出
func SafeUintToUint32(value uint) (uint32, error) {
	if value > math.MaxUint32 {
		return 0, fmt.Errorf("value %d overflow, uint32 max value is %d", value, math.MaxUint32)
	}
	return uint32(value), nil
}

// SafeUintToUint64 安全地将 uint 转换为 uint64
func SafeUintToUint64(value uint) uint64 {
	return uint64(value)
}

// SafeUintToFloat32 安全地将 uint 转换为 float32
func SafeUintToFloat32(value uint) float32 {
	return float32(value)
}

// SafeUintToFloat64 安全地将 uint 转换为 float64
func SafeUintToFloat64(value uint) float64 {
	return float64(value)
}

// SafeFloat32ToUint8 安全地将 float32 转换为 uint8，避免负数、溢出和精度丢失
func SafeFloat32ToUint8(value float32) (uint8, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to uint8", value)
	}
	// 检查是否为负数
	if value < 0 {
		return 0, fmt.Errorf("value %f is negative, cannot convert to uint8", value)
	}
	// 检查是否有小数部分
	if value != float32(math.Trunc(float64(value))) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to uint8 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("value %f overflow, uint8 max value is %d", value, math.MaxUint8)
	}
	return uint8(value), nil
}

// SafeFloat32ToUint16 安全地将 float32 转换为 uint16，避免负数、溢出和精度丢失
func SafeFloat32ToUint16(value float32) (uint16, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to uint16", value)
	}
	// 检查是否为负数
	if value < 0 {
		return 0, fmt.Errorf("value %f is negative, cannot convert to uint16", value)
	}
	// 检查是否有小数部分
	if value != float32(math.Trunc(float64(value))) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to uint16 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxUint16 {
		return 0, fmt.Errorf("value %f overflow, uint16 max value is %d", value, math.MaxUint16)
	}
	return uint16(value), nil
}

// SafeFloat32ToUint32 安全地将 float32 转换为 uint32，避免负数、溢出和精度丢失
func SafeFloat32ToUint32(value float32) (uint32, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to uint32", value)
	}
	// 检查是否为负数
	if value < 0 {
		return 0, fmt.Errorf("value %f is negative, cannot convert to uint32", value)
	}
	// 检查是否有小数部分
	if value != float32(math.Trunc(float64(value))) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to uint32 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxUint32 {
		return 0, fmt.Errorf("value %f overflow, uint32 max value is %d", value, math.MaxUint32)
	}
	return uint32(value), nil
}

// SafeFloat32ToUint64 安全地将 float32 转换为 uint64，避免负数、溢出和精度丢失
func SafeFloat32ToUint64(value float32) (uint64, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to uint64", value)
	}
	// 检查是否为负数
	if value < 0 {
		return 0, fmt.Errorf("value %f is negative, cannot convert to uint64", value)
	}
	// 检查是否有小数部分
	if value != float32(math.Trunc(float64(value))) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to uint64 without precision loss", value)
	}
	// 检查溢出
	if float64(value) > math.MaxUint64 {
		return 0, fmt.Errorf("value %f overflow, uint64 max value is %d", value, uint64(math.MaxUint64))
	}
	return uint64(value), nil
}

// SafeFloat64ToUint8 安全地将 float64 转换为 uint8，避免负数、溢出和精度丢失
func SafeFloat64ToUint8(value float64) (uint8, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to uint8", value)
	}
	// 检查是否为负数
	if value < 0 {
		return 0, fmt.Errorf("value %f is negative, cannot convert to uint8", value)
	}
	// 检查是否有小数部分
	if value != math.Trunc(value) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to uint8 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("value %f overflow, uint8 max value is %d", value, math.MaxUint8)
	}
	return uint8(value), nil
}

// SafeFloat64ToUint16 安全地将 float64 转换为 uint16，避免负数、溢出和精度丢失
func SafeFloat64ToUint16(value float64) (uint16, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to uint16", value)
	}
	// 检查是否为负数
	if value < 0 {
		return 0, fmt.Errorf("value %f is negative, cannot convert to uint16", value)
	}
	// 检查是否有小数部分
	if value != math.Trunc(value) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to uint16 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxUint16 {
		return 0, fmt.Errorf("value %f overflow, uint16 max value is %d", value, math.MaxUint16)
	}
	return uint16(value), nil
}

// SafeFloat64ToUint32 安全地将 float64 转换为 uint32，避免负数、溢出和精度丢失
func SafeFloat64ToUint32(value float64) (uint32, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to uint32", value)
	}
	// 检查是否为负数
	if value < 0 {
		return 0, fmt.Errorf("value %f is negative, cannot convert to uint32", value)
	}
	// 检查是否有小数部分
	if value != math.Trunc(value) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to uint32 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxUint32 {
		return 0, fmt.Errorf("value %f overflow, uint32 max value is %d", value, math.MaxUint32)
	}
	return uint32(value), nil
}

// SafeFloat64ToUint64 安全地将 float64 转换为 uint64，避免负数、溢出和精度丢失
func SafeFloat64ToUint64(value float64) (uint64, error) {
	// 检查是否为无穷大或NaN
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, fmt.Errorf("value %f is infinity or NaN, cannot convert to uint64", value)
	}
	// 检查是否为负数
	if value < 0 {
		return 0, fmt.Errorf("value %f is negative, cannot convert to uint64", value)
	}
	// 检查是否有小数部分
	if value != math.Trunc(value) {
		return 0, fmt.Errorf("value %f has fractional part, cannot convert to uint64 without precision loss", value)
	}
	// 检查溢出
	if value > math.MaxUint64 {
		return 0, fmt.Errorf("value %f overflow, uint64 max value is %d", value, uint64(math.MaxUint64))
	}
	return uint64(value), nil
}
