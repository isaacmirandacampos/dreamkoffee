package usecases

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
)

type expenseUseCase struct {
	repo *persistence.Queries
}

func NewExpenseUseCase(repo *persistence.Queries) *expenseUseCase {
	return &expenseUseCase{repo: repo}
}

func (uc *expenseUseCase) ListExpenses(ctx context.Context) ([]*model.Expense, error) {
	expenses, err := uc.repo.ListExpenses(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.Expense
	for _, expense := range expenses {
		result = append(result, &model.Expense{
			ID:        int(expense.ID),
			Name:      expense.Name,
			Price:     expense.Price,
			CreatedAt: expense.CreatedAt.String(),
			UpdatedAt: expense.UpdatedAt.String(),
		})
	}

	return result, nil
}
