package expense

import (
	"context"

	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/error_handling"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/handlers/graph/model"
	"github.com/isaacmirandacampos/dreamkoffee/internal/domain/entity"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres/persistence"
)

func (c *expenseUseCase) UpdateExpense(ctx context.Context, id *int, input model.UpdateExpense) (*model.Expense, error) {
	expenseEntity := entity.NewExpense(
		&entity.Expense{
			ID:          int32(*id),
			Description: input.Description,
			Value:       input.Value,
		},
	)
	err := expenseEntity.ValueIsValid()
	if err != nil {
		return nil, error_handling.Graphql(ctx, 400, "value_must_be_positive", "Value must be positive", err.Error())
	}

	_, err = c.repo.GetExpense(ctx, expenseEntity.ID)
	if err != nil {
		return nil, error_handling.Graphql(ctx, 404, "expense_not_found", "Expense not found", err.Error())
	}
	updated, err := c.repo.UpdateExpense(ctx, &persistence.UpdateExpenseParams{
		ID:          expenseEntity.ID,
		Description: expenseEntity.Description,
		Value:       expenseEntity.Value,
	})
	if err != nil {
		return nil, error_handling.Graphql(ctx, 500, "internal_server_error", "Could not update expense", err.Error())
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
