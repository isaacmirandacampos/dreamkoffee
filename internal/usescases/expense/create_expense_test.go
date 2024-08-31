package expense_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/graph/model"
	"github.com/isaacmirandacampos/dreamkoffee/internal/domain"
	"github.com/isaacmirandacampos/dreamkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/dreamkoffee/internal/test/mocks"
	"github.com/isaacmirandacampos/dreamkoffee/internal/usescases/expense"
	"github.com/isaacmirandacampos/dreamkoffee/internal/utils"
	"github.com/stretchr/testify/assert"
)

func RepositoryMock(repo *mocks.MockRepository) domain.Repository {
	return repo
}

func TestCreateExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	repository := RepositoryMock(mockRepo)
	useCase := expense.NewExpenseUseCase(repository)
	ctx := context.Background()
	value, err := utils.UnmarshalDecimal("5.0")
	assert.NoError(t, err)
	input := model.NewExpense{
		Description: "Coffee",
		Value:       value,
	}
	t.Run("CreateExpense", func(t *testing.T) {
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

	t.Run("CreateExpenseError", func(t *testing.T) {
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
