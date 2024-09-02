package expense_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/handlers/graph/model"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/usecases/expense"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/mocks"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres/persistence"
	"github.com/isaacmirandacampos/dreamkoffee/internal/test/helper"
	"github.com/isaacmirandacampos/dreamkoffee/pkg/scalar"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	repository := helper.RepositoryMock(mockRepo)
	useCase := expense.NewExpenseUseCase(&repository)
	ctx := context.Background()

	t.Run("Should create a expense", func(t *testing.T) {
		value, err := scalar.UnmarshalDecimal("5.0")
		assert.NoError(t, err)
		input := model.NewExpense{
			Description: "Coffee",
			Value:       value,
		}
		expectedExpense := persistence.Expense{
			ID:          1,
			Description: "Coffee",
			Value:       value,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		mockRepo.EXPECT().CreateExpense(ctx, gomock.Any()).Times(1).Return(expectedExpense, nil)

		result, err := useCase.CreateExpense(ctx, input)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int(expectedExpense.ID), result.ID)
		assert.Equal(t, expectedExpense.Description, result.Description)
		assert.Equal(t, expectedExpense.Value, result.Value)
		assert.Equal(t, expectedExpense.CreatedAt.String(), result.CreatedAt)
		assert.Equal(t, expectedExpense.UpdatedAt.String(), result.UpdatedAt)
	})

	t.Run("Shouldn't accept create a expense with negative value", func(t *testing.T) {
		value, err := scalar.UnmarshalDecimal("-5.0")
		assert.NoError(t, err)
		input := model.NewExpense{
			Description: "Coffee",
			Value:       value,
		}
		_, err = useCase.CreateExpense(ctx, input)
		assert.EqualError(t, err, "input: Value must be positive")
	})

	t.Run("Should fail create a expense when run the insert query", func(t *testing.T) {
		value, err := scalar.UnmarshalDecimal("5.0")
		assert.NoError(t, err)
		input := model.NewExpense{
			Description: "Coffee",
			Value:       value,
		}
		mockRepo.EXPECT().CreateExpense(ctx, gomock.Any()).Times(1).Return(persistence.Expense{}, errors.New("Invalid input"))
		result, err := useCase.CreateExpense(ctx, input)
		if err == nil {
			t.Fatalf("Was expected error")
		}
		if result != nil {
			t.Fatalf("Was expected nil in result")
		}
		assert.Equal(t, "input: Could not create expense", err.Error())
	})
}
