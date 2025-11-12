package egorm

import (
	"database/sql/driver"
	"fmt"
	edecimal "github.com/guanguoyintao/kuafu/decimal"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DecimalField 用于 gorm 的 decimal 类型封装
type DecimalField edecimal.Decimal

// Scan 实现 sql.Scanner 接口
func (d *DecimalField) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var err error
	var decimal *edecimal.Decimal

	switch v := value.(type) {
	case string:
		decimal, err = edecimal.NewFromString(v)
	case []byte:
		decimal, err = edecimal.NewFromString(string(v))
	case float64:
		decimal, err = edecimal.NewFromString(fmt.Sprintf("%v", v))
	default:
		return fmt.Errorf("unsupported type for DecimalField: %T", value)
	}

	if err != nil {
		return err
	}

	*d = DecimalField(*decimal)
	return nil
}

// Value 实现 driver.Valuer 接口
func (d DecimalField) Value() (driver.Value, error) {
	if d == (DecimalField{}) {
		return nil, nil
	}
	decimal := edecimal.Decimal(d)
	return decimal.String(), nil
}

// GormDataType 实现 schema.GormDataTypeInterface 接口
func (DecimalField) GormDataType() string {
	return "decimal"
}

// GormDBDataType 实现 schema.GormDBDataTypeInterface 接口
func (DecimalField) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// 可以根据需要调整精度和小数位数
	return "decimal(65,30)"
}

// ToDecimal 类型转换函数
func ToDecimal(d DecimalField) *edecimal.Decimal {
	decimal := edecimal.Decimal(d)
	return &decimal
}

func FromDecimal(d *edecimal.Decimal) DecimalField {
	if d == nil {
		return DecimalField{}
	}
	return DecimalField(*d)
}
