package expense

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/finkoffee/internal/utils"
)

func (c *expenseUseCase) CreateExpense(ctx context.Context, input model.NewExpense) (*model.Expense, error) {
	returned, err := c.repo.CreateExpense(ctx, persistence.CreateExpenseParams{
		Description: input.Description,
		Value:       input.Value,
	})
	if err != nil {
		return nil, utils.ErrorHandling(ctx, 400, "bad_request", "Could not create expense", err.Error())
	}
	expense := &model.Expense{
		ID:          int(returned.ID),
		Value:       returned.Value,
		Description: returned.Description,
		CreatedAt:   returned.CreatedAt.String(),
		UpdatedAt:   returned.UpdatedAt.String(),
	}
	return expense, nil
}
