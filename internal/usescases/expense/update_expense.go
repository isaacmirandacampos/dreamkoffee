package expense

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
)

func (c *expenseUseCase) UpdateExpense(ctx context.Context, id *int, input model.UpdateExpense) (*model.Expense, error) {
	updated, err := c.repo.UpdateExpense(ctx, persistence.UpdateExpenseParams{
		ID:          int32(*id),
		Description: input.Description,
		Value:       input.Value,
	})
	if err != nil {
		return nil, err
	}
	expense := &model.Expense{
		ID:          int(updated.ID),
		Value:       updated.Value,
		Description: updated.Description,
		CreatedAt:   updated.CreatedAt.String(),
		UpdatedAt:   updated.UpdatedAt.String(),
	}
	return expense, nil
}
