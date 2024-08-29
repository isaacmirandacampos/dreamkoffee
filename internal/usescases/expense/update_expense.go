package expense

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/applications/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
)

func (c *expenseUseCase) UpdateExpense(ctx context.Context, id *int, input model.UpdateExpense) (*model.Expense, error) {
	updated, err := c.repo.UpdateExpense(ctx, persistence.UpdateExpenseParams{
		ID:    int32(*id),
		Name:  input.Name,
		Price: input.Price,
	})
	if err != nil {
		return nil, err
	}
	expense := &model.Expense{
		ID:        int(updated.ID),
		Price:     updated.Price,
		Name:      updated.Name,
		CreatedAt: updated.CreatedAt.String(),
		UpdatedAt: updated.UpdatedAt.String(),
	}
	return expense, nil
}
