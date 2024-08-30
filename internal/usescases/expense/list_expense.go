package expense

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph/model"
	"github.com/isaacmirandacampos/finkoffee/internal/utils"
)

func (uc *expenseUseCase) ListExpenses(ctx context.Context) ([]*model.Expense, error) {
	results, err := uc.repo.ListExpenses(ctx)
	if err != nil {
		return nil, utils.ErrorHandling(ctx, 500, "internal_server_error", "Could not list expenses", err.Error())
	}
	expenses := make([]*model.Expense, 0, len(results))

	for _, expense := range results {
		expenses = append(expenses, &model.Expense{
			ID:          int(expense.ID),
			Description: expense.Description,
			Value:       expense.Value,
			CreatedAt:   expense.CreatedAt.String(),
			UpdatedAt:   expense.UpdatedAt.String(),
		})
	}

	return expenses, nil
}
