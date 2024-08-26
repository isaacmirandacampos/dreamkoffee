package utils_test

import (
	"testing"

	"github.com/isaacmirandacampos/finkoffee/internal/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalDecimal(t *testing.T) {
	t.Run("convert_string_to_decimal", func(t *testing.T){
		d, err := utils.UnmarshalDecimal("100")
		assert.Nil(t, err)
		assert.Equal(t, decimal.NewFromInt(100), d)
	})
	t.Run("convert_float_to_decimal", func(t *testing.T){
		d, err := utils.UnmarshalDecimal(100.0)
		assert.Nil(t, err)
		assert.Equal(t, decimal.NewFromFloat(100.0), d)
	})
	t.Run("convert_int_to_decimal", func(t *testing.T){
		d, err := utils.UnmarshalDecimal(100)
		assert.Nil(t, err)
		assert.Equal(t, decimal.NewFromInt(100), d)
	})
	t.Run("convert_invalid_type_to_decimal", func(t *testing.T){
		_, err := utils.UnmarshalDecimal(true)
		assert.NotNil(t, err)
	})
}

func TestMarshalDecimal(t *testing.T) {
	t.Run("convert_decimal_to_string", func(t *testing.T){
		d := decimal.NewFromInt(100)
		m := utils.MarshalDecimal(d)
		assert.Equal(t, "\"100\"", m)
	})
}
