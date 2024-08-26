package utils

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func MarshalDecimal(d decimal.Decimal) string {
	return fmt.Sprintf("\"%s\"", d.String())
}

func UnmarshalDecimal(v interface{}) (decimal.Decimal, error) {
	switch v := v.(type) {
	case string:
		return decimal.NewFromString(v)
	case float64:
		return decimal.NewFromFloat(v), nil
	case int:
		return decimal.NewFromInt(int64(v)), nil
	default:
		return decimal.Decimal{}, fmt.Errorf("unexpected type %T for Decimal", v)
	}
}
