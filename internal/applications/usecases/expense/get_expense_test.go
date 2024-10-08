package expense_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/usecases/expense"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/mocks"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres/persistence"
	"github.com/isaacmirandacampos/dreamkoffee/internal/test/helper"
	"github.com/isaacmirandacampos/dreamkoffee/pkg/scalar"
	"github.com/stretchr/testify/assert"
)

func TestGetExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	repository := helper.RepositoryMock(mockRepo)
	useCase := expense.NewExpenseUseCase(&repository)
	ctx := context.Background()
	value, err := scalar.UnmarshalDecimal("5.0")
	assert.NoError(t, err)
	expectedExpense := persistence.Expense{
		ID:          1,
		Description: "Coffee",
		Value:       value,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	id := int(expectedExpense.ID)
	t.Run("Should get a existent expense", func(t *testing.T) {
		mockRepo.EXPECT().GetExpense(ctx, int32(id)).Times(1).Return(expectedExpense, nil)
		result, err := useCase.GetExpense(ctx, &id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, id, result.ID)
		assert.Equal(t, expectedExpense.Description, result.Description)
		assert.Equal(t, expectedExpense.Value, result.Value)
		assert.Equal(t, expectedExpense.CreatedAt.String(), result.CreatedAt)
		assert.Equal(t, expectedExpense.UpdatedAt.String(), result.UpdatedAt)
	})

	t.Run("Shouldn't found a expense", func(t *testing.T) {
		mockRepo.EXPECT().GetExpense(ctx, gomock.Any()).Times(1).Return(persistence.Expense{}, errors.New("Expense not found"))
		result, err := useCase.GetExpense(ctx, &id)
		if err == nil {
			t.Fatalf("Was expected error")
		}
		if result != nil {
			t.Fatalf("Was expected nil in result")
		}
		assert.Equal(t, "input: Expense not found", err.Error())
	})
}
