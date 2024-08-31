package expense

import (
	"context"

	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/graph/model"
	"github.com/isaacmirandacampos/dreamkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/dreamkoffee/internal/utils"
)

func (c *expenseUseCase) UpdateExpense(ctx context.Context, id *int, input model.UpdateExpense) (*model.Expense, error) {
	if !input.Value.IsPositive() {
		return nil, utils.ErrorHandling(ctx, 400, "value_must_be_positive", "Value must be positive", input.Value)
	}
	_, err := c.repo.GetExpense(ctx, int32(*id))
	if err != nil {
		return nil, utils.ErrorHandling(ctx, 404, "expense_not_found", "Expense not found", err.Error())
	}
	updated, err := c.repo.UpdateExpense(ctx, persistence.UpdateExpenseParams{
		ID:          int32(*id),
		Description: input.Description,
		Value:       input.Value,
	})
	if err != nil {
		return nil, utils.ErrorHandling(ctx, 500, "internal_server_error", "Could not update expense", err.Error())
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
