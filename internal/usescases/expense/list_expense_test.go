package expense_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/isaacmirandacampos/dreamkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/dreamkoffee/internal/test/helper"
	"github.com/isaacmirandacampos/dreamkoffee/internal/test/mocks"
	"github.com/isaacmirandacampos/dreamkoffee/internal/usescases/expense"
	"github.com/isaacmirandacampos/dreamkoffee/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestListExpenses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	repository := helper.RepositoryMock(mockRepo)
	useCase := expense.NewExpenseUseCase(repository)
	ctx := context.Background()
	value, err := utils.UnmarshalDecimal("5.0")
	assert.NoError(t, err)

	t.Run("Should list expenses", func(t *testing.T) {
		expectedExpenses := []persistence.Expense{
			{
				ID:          1,
				Description: "Coffee",
				Value:       value,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}
		mockRepo.EXPECT().ListExpenses(ctx).Times(1).Return(expectedExpenses, nil)
		result, err := useCase.ListExpenses(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int(expectedExpenses[0].ID), result[0].ID)
		assert.Equal(t, expectedExpenses[0].Description, result[0].Description)
		assert.Equal(t, expectedExpenses[0].Value, result[0].Value)
		assert.Equal(t, expectedExpenses[0].CreatedAt.String(), result[0].CreatedAt)
		assert.Equal(t, expectedExpenses[0].UpdatedAt.String(), result[0].UpdatedAt)
	})

	t.Run("Should only return a empty array", func(t *testing.T) {
		mockRepo.EXPECT().ListExpenses(ctx).Times(1).Return([]persistence.Expense{}, nil)
		result, err := useCase.ListExpenses(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result)
	})
}
