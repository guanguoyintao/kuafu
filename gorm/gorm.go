package egorm

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	NULLString  = "-1"
	NULLInt     = int(-1)
	NULLInt8    = int8(-1)
	NULLInt16   = int16(-1)
	NULLInt32   = int32(-1)
	NULLInt64   = int64(-1)
	NULLUint    = uint(0)
	NULLUint8   = uint8(0)
	NULLUint16  = uint16(0)
	NULLUint32  = uint32(0)
	NULLUint64  = uint64(0)
	NULLFloat32 = float32(-1.0)
	NULLFloat64 = float64(-1.0)
)

const (
	BoolFalse int8 = 0
	BoolTrue  int8 = 1
)

func IsTrue(value int8) bool {
	if value == BoolTrue {
		return true
	} else {
		return false
	}
}

func SetBool(value bool) int8 {
	if value {
		return BoolTrue
	} else {
		return BoolFalse
	}
}

func ConvertDo2Bo[T any](value T) T {
	// 使用类型断言判断类型并进行比较
	switch v := any(value).(type) {
	case string:
		if v == NULLString {
			return any("").(T)
		}
	case int:
		if v == NULLInt {
			return any(0).(T)
		}
	case int8:
		if v == NULLInt8 {
			return any(int8(0)).(T)
		}
	case int16:
		if v == NULLInt16 {
			return any(int16(0)).(T)
		}
	case int32:
		if v == NULLInt32 {
			return any(int32(0)).(T)
		}
	case int64:
		if v == NULLInt64 {
			return any(int64(0)).(T)
		}
	case uint:
		if v == NULLUint {
			return any(uint(0)).(T)
		}
	case uint8:
		if v == NULLUint8 {
			return any(uint8(0)).(T)
		}
	case uint16:
		if v == NULLUint16 {
			return any(uint16(0)).(T)
		}
	case uint32:
		if v == NULLUint32 {
			return any(uint32(0)).(T)
		}
	case uint64:
		if v == NULLUint64 {
			return any(uint64(0)).(T)
		}
	case float32:
		if v == NULLFloat32 {
			return any(float32(0)).(T)
		}
	case float64:
		if v == NULLFloat64 {
			return any(float64(0)).(T)
		}
	default:
		// 默认情况下可以返回原值，或者返回一个零值
		return value
	}
	// 若没有匹配到 NULL 值，直接返回输入的值
	return value
}

func ConvertDo2BoPtr[T any](value T) *T {
	// 使用类型断言判断类型并进行比较
	switch v := any(value).(type) {
	case string:
		if v == NULLString {
			return nil
		}
	case int:
		if v == NULLInt {
			return nil
		}
	case int8:
		if v == NULLInt8 {
			return nil
		}
	case int16:
		if v == NULLInt16 {
			return nil
		}
	case int32:
		if v == NULLInt32 {
			return nil
		}
	case int64:
		if v == NULLInt64 {
			return nil
		}
	case uint:
		if v == NULLUint {
			return nil
		}
	case uint8:
		if v == NULLUint8 {
			return nil
		}
	case uint16:
		if v == NULLUint16 {
			return nil
		}
	case uint32:
		if v == NULLUint32 {
			return nil
		}
	case uint64:
		if v == NULLUint64 {
			return nil
		}
	case float32:
		if v == NULLFloat32 {
			return nil
		}
	case float64:
		if v == NULLFloat64 {
			return nil
		}
	default:
		// 默认情况下可以返回原值，或者返回一个零值
		return &value
	}
	// 若没有匹配到 NULL 值，直接返回输入的值
	return &value
}

func ConvertBo2Do[T any](value T) T {
	// 使用类型断言判断类型并进行比较
	switch v := any(value).(type) {
	case string:
		if v == "" {
			return any(NULLString).(T)
		}
	case int:
		if v == 0 {
			return any(NULLInt).(T)
		}
	case int8:
		if v == 0 {
			return any(NULLInt8).(T)
		}
	case int16:
		if v == 0 {
			return any(NULLInt16).(T)
		}
	case int32:
		if v == 0 {
			return any(NULLInt32).(T)
		}
	case int64:
		if v == 0 {
			return any(NULLInt64).(T)
		}
	case uint:
		if v == 0 {
			return any(NULLUint).(T)
		}
	case uint8:
		if v == 0 {
			return any(NULLUint8).(T)
		}
	case uint16:
		if v == 0 {
			return any(NULLUint16).(T)
		}
	case uint32:
		if v == 0 {
			return any(NULLUint32).(T)
		}
	case uint64:
		if v == 0 {
			return any(NULLUint64).(T)
		}
	case float32:
		if v == 0 {
			return any(NULLFloat32).(T)
		}
	case float64:
		if v == 0 {
			return any(NULLFloat64).(T)
		}
	default:
		// 默认情况下可以返回原值，或者返回一个零值
		return value
	}
	// 若没有匹配到零值，直接返回输入的值
	return value
}

func ConvertBoPtr2Do[T any](value *T) T {
	// 使用类型断言判断类型并进行比较
	switch v := any(value).(type) {
	case *string:
		if v == nil {
			return any(NULLString).(T)
		}
	case *int:
		if v == nil {
			return any(NULLInt).(T)
		}
	case *int8:
		if v == nil {
			return any(NULLInt8).(T)
		}
	case *int16:
		if v == nil {
			return any(NULLInt16).(T)
		}
	case *int32:
		if v == nil {
			return any(NULLInt32).(T)
		}
	case *int64:
		if v == nil {
			return any(NULLInt64).(T)
		}
	case *uint:
		if v == nil {
			return any(NULLUint).(T)
		}
	case *uint8:
		if v == nil {
			return any(NULLUint8).(T)
		}
	case *uint16:
		if v == nil {
			return any(NULLUint16).(T)
		}
	case *uint32:
		if v == nil {
			return any(NULLUint32).(T)
		}
	case *uint64:
		if v == nil {
			return any(NULLUint64).(T)
		}
	case *float32:
		if v == nil {
			return any(NULLFloat32).(T)
		}
	case *float64:
		if v == nil {
			return any(NULLFloat64).(T)
		}
	default:
		// 默认情况下可以返回原值，或者返回一个零值
		return *value
	}
	// 若没有匹配到零值，直接返回输入的值
	return *value
}

// todo: gorm业务空值类型映射

// ExtractDuplicateEntry 从错误信息中提取冲突的唯一键值
func ExtractDuplicateEntry(mySQLError *mysql.MySQLError) (*DuplicateEntry, error) {
	if mySQLError.Number == 1062 {
		// 使用正则表达式匹配单引号内的冲突值
		re := regexp.MustCompile(`Duplicate entry '([^']+)'`)
		matches := re.FindStringSubmatch(mySQLError.Message)
		if len(matches) < 2 {
			return nil, fmt.Errorf("no duplicate entry found in the error message")
		}
		// 获取匹配的唯一键值
		duplicateEntry := matches[1]
		// 分割键值为多个部分
		parts := strings.Split(duplicateEntry, "-")
		// 使用正则表达式匹配 for key '...' 部分
		reKey := regexp.MustCompile(`for key '([^']+)'`)
		keyMatches := reKey.FindStringSubmatch(mySQLError.Message)
		if len(keyMatches) < 2 {
			return nil, fmt.Errorf("no key found in the error message")
		}
		// 获取匹配的唯一键名（例如：affiliate_invited_user.uk_device）
		duplicateKey := keyMatches[1]
		// 提取 . 后面的部分（即：uk_device）
		duplicateKeys := strings.Split(duplicateKey, ".")
		if len(parts) > 1 {
			// 返回 . 后面的部分，即最后一部分
			return &DuplicateEntry{
				Entry: duplicateEntry,
				Parts: parts,
				Key:   duplicateKeys[len(duplicateKeys)-1],
			}, nil
		}

		return &DuplicateEntry{
			Entry: duplicateEntry,
			Parts: parts,
			Key:   duplicateKeys[0],
		}, nil
	}
	return nil, nil
}
