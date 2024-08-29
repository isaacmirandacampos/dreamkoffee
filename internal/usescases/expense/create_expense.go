package expense

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/applications/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
)

func (c *expenseUseCase) CreateExpense(ctx context.Context, input model.NewExpense) (*model.Expense, error) {
	returned, err := c.repo.CreateExpense(ctx, persistence.CreateExpenseParams{
		Name:  input.Name,
		Price: input.Price,
	})
	if err != nil {
		return nil, err
	}
	expense := &model.Expense{
		ID:        int(returned.ID),
		Price:     returned.Price,
		Name:      returned.Name,
		CreatedAt: returned.CreatedAt.String(),
		UpdatedAt: returned.UpdatedAt.String(),
	}
	return expense, nil
}
