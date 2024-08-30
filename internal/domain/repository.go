package domain

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
)

type Repository interface {
	CreateExpense(ctx context.Context, arg persistence.CreateExpenseParams) (persistence.Expense, error)
	DeleteExpense(ctx context.Context, id int32) (persistence.Expense, error)
	GetExpense(ctx context.Context, id int32) (persistence.Expense, error)
	GetLastExpense(ctx context.Context) (persistence.Expense, error)
	ListExpenses(ctx context.Context) ([]persistence.Expense, error)
	UpdateExpense(ctx context.Context, arg persistence.UpdateExpenseParams) (persistence.Expense, error)
}

func NewRepository(repo *persistence.Queries) Repository {
	return repo
}
