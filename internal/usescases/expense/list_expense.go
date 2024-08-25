package expense

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
)

func (uc *expenseUseCase) ListExpenses(ctx context.Context) ([]*model.Expense, error) {
	results, err := uc.repo.ListExpenses(ctx)
	if err != nil {
		return nil, err
	}
	expenses := make([]*model.Expense, 0, len(results))

	for _, expense := range results {
		expenses = append(expenses, &model.Expense{
			ID:        int(expense.ID),
			Name:      expense.Name,
			Price:     expense.Price,
			CreatedAt: expense.CreatedAt.String(),
			UpdatedAt: expense.UpdatedAt.String(),
		})
	}

	return expenses, nil
}
