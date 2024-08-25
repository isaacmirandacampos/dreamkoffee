package tests

import (
	"context"
	"testing"

	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/finkoffee/internal/usescases/expense"
	"github.com/isaacmirandacampos/finkoffee/internal/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	repo := persistence.New(db)
	useCase := expense.NewExpenseUseCase(repo)
	ctx := context.Background()

	price, err := utils.UnmarshalDecimal(100)
	assert.NoError(t, err)
	input := model.NewExpense{
		Name:  "Test Expense",
		Price: price,
	}

	result, err := useCase.CreateExpense(ctx, input)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Expense", result.Name)
	assert.IsType(t, decimal.Decimal{}, result.Price)
}
