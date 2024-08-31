package entity_test

import (
	"testing"

	"github.com/isaacmirandacampos/dreamkoffee/internal/domain/entity"
	"github.com/isaacmirandacampos/dreamkoffee/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestExpenseEntity(t *testing.T) {
	t.Run("Should validate the ValueIsValid method", func(t *testing.T) {
		t.Run("Should value be valid", func(t *testing.T) {
			value, err := utils.UnmarshalDecimal("10.0")
			assert.NoError(t, err)
			expense := entity.New(
				&entity.Expense{
					Description: "Test Description",
					Value:       value,
				},
			)
			assert.NotNil(t, expense)
			assert.Equal(t, "Test Description", expense.Description)
			assert.Equal(t, value, expense.Value)
			err = expense.ValueIsValid()
			assert.NoError(t, err)
		})

		t.Run("Should error when value be negative", func(t *testing.T) {
			value, err := utils.UnmarshalDecimal("-10")
			assert.NoError(t, err)
			expense := entity.New(
				&entity.Expense{
					Description: "Test Description",
					Value:       value,
				},
			)
			assert.NotNil(t, expense)
			assert.Equal(t, "Test Description", expense.Description)
			assert.Equal(t, value, expense.Value)
			err = expense.ValueIsValid()
			assert.EqualError(t, err, "value cannot be negative")
		})

		t.Run("Should error when value be zero", func(t *testing.T) {
			value, err := utils.UnmarshalDecimal("0")
			assert.NoError(t, err)
			expense := entity.New(
				&entity.Expense{
					Description: "Test Description",
					Value:       value,
				},
			)
			assert.NotNil(t, expense)
			assert.Equal(t, "Test Description", expense.Description)
			assert.Equal(t, value, expense.Value)
			err = expense.ValueIsValid()
			assert.EqualError(t, err, "value cannot be zero")
		})
	})
}
