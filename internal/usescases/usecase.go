package usecases

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
)

type ExpenseUseCase interface {
	ListExpenses(ctx context.Context) ([]*model.Expense, error)
}
