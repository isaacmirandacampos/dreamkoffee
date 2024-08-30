package expense_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph/model"
	"github.com/isaacmirandacampos/finkoffee/internal/domain"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/finkoffee/internal/test/mocks"
	"github.com/isaacmirandacampos/finkoffee/internal/usescases/expense"
	"github.com/isaacmirandacampos/finkoffee/internal/utils"
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
		mockRepo.EXPECT().CreateExpense(ctx, gomock.Any()).Times(1).Return(persistence.Expense{}, utils.ErrorHandling(ctx, 400, "bad_request", "Could not create expense"))
		result, err := useCase.CreateExpense(ctx, input)
		if result != nil {
			t.Fatalf("Expected nil")
		}
		if err == nil {
			t.Fatalf("Expected error")
		}
	})
}
