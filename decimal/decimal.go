package edecimal

import (
	"fmt"
	"math/big"
	"strings"
)

// Decimal 表示一个高精度小数
type Decimal struct {
	value    *big.Int // 存储去除小数点后的数值
	scale    int      // 小数位数
	negative bool     // 是否为负数
}

// NewFromString 从字符串创建Decimal
func NewFromString(s string) (*Decimal, error) {
	if s == "" {
		return nil, fmt.Errorf("empty string")
	}
	// 处理负号
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	// 分离整数和小数部分
	parts := strings.Split(s, ".")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid decimal format")
	}

	// 计算scale并构造纯数字字符串
	scale := 0
	numStr := parts[0]
	if len(parts) == 2 {
		scale = len(parts[1])
		numStr += parts[1]
	}

	// 转换为big.Int
	value := new(big.Int)
	_, ok := value.SetString(numStr, 10)
	if !ok {
		return nil, fmt.Errorf("invalid number format")
	}

	return &Decimal{
		value:    value,
		scale:    scale,
		negative: negative,
	}, nil
}

// Float64 将 Decimal 转换为 float64，保持精确度
func (d *Decimal) Float64() float64 {
	if d == nil || d.value == nil {
		return 0
	}

	// 获取整数部分
	intPart := new(big.Int).Set(d.value)
	scale := d.scale

	// 如果有小数位，需要进行除法运算
	if scale > 0 {
		divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(scale)), nil)
		// 使用 big.Rat 进行精确除法
		rat := new(big.Rat).SetFrac(intPart, divisor)
		result, _ := rat.Float64()

		if d.negative {
			return -result
		}
		return result
	}

	// 如果没有小数位，直接转换
	result, _ := new(big.Rat).SetInt(intPart).Float64()
	if d.negative {
		return -result
	}
	return result
}

// String 将Decimal转换为字符串
func (d *Decimal) String() string {
	if d == nil || d.value == nil {
		return "0"
	}

	// 获取绝对值字符串
	s := d.value.String()

	// 补齐前导零
	if len(s) <= d.scale {
		s = strings.Repeat("0", d.scale-len(s)+1) + s
	}

	// 插入小数点
	if d.scale > 0 {
		pos := len(s) - d.scale
		s = s[:pos] + "." + s[pos:]
	}

	// 添加负号
	if d.negative && s != "0" {
		s = "-" + s
	}

	return s
}

// normalize 标准化两个Decimal的scale
func (d *Decimal) normalize(d1, d2 *Decimal) (*big.Int, *big.Int, int) {
	scale := max(d1.scale, d2.scale)
	v1 := new(big.Int).Set(d1.value)
	v2 := new(big.Int).Set(d2.value)

	// 将scale较小的数扩大到相同scale
	if d1.scale < scale {
		v1.Mul(v1, d.pow10(scale-d1.scale))
	}
	if d2.scale < scale {
		v2.Mul(v2, d.pow10(scale-d2.scale))
	}

	return v1, v2, scale
}

// Add 加法运算
func (d *Decimal) Add(other *Decimal) *Decimal {
	v1, v2, scale := d.normalize(d, other)

	result := new(big.Int)
	if d.negative == other.negative {
		result.Add(v1, v2)
	} else {
		// 不同号时进行减法
		if v1.CmpAbs(v2) >= 0 {
			result.Sub(v1, v2)
		} else {
			result.Sub(v2, v1)
			d.negative = !d.negative
		}
	}

	return &Decimal{
		value:    result,
		scale:    scale,
		negative: d.negative && result.Sign() != 0,
	}
}

// Sub 减法运算
func (d *Decimal) Sub(other *Decimal) *Decimal {
	// 转换为加上相反数
	neg := &Decimal{
		value:    other.value,
		scale:    other.scale,
		negative: !other.negative,
	}
	return d.Add(neg)
}

// Mul 乘法运算
func (d *Decimal) Mul(other *Decimal) *Decimal {
	result := new(big.Int).Mul(d.value, other.value)
	return &Decimal{
		value:    result,
		scale:    d.scale + other.scale,
		negative: d.negative != other.negative,
	}
}

// Div 执行高精度除法运算
func (d *Decimal) Div(other *Decimal, precision int) (*Decimal, error) {
	if other.value.Sign() == 0 {
		return nil, fmt.Errorf("division by zero")
	}

	// 计算目标精度的缩放因子
	scaleFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil)

	// 扩展被除数以达到所需精度
	numerator := new(big.Int).Mul(d.value, scaleFactor)

	// 执行除法运算
	quotient := new(big.Int).Quo(numerator, other.value)

	// 计算最终的小数位数
	finalScale := d.scale - other.scale + precision

	return &Decimal{
		value:    quotient,
		scale:    finalScale,
		negative: d.negative != other.negative,
	}, nil
}

// pow10 返回10的n次方
func (d *Decimal) pow10(n int) *big.Int {
	result := new(big.Int).SetInt64(1)
	if n > 0 {
		result.Exp(big.NewInt(10), big.NewInt(int64(n)), nil)
	}
	return result
}

func (d *Decimal) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
