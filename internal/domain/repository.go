package domain

import (
	"context"

	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres/persistence"
)

type Repository interface {
	// expenses
	CreateExpense(ctx context.Context, arg *persistence.CreateExpenseParams) (*persistence.Expense, error)
	DeleteExpense(ctx context.Context, id int32) (*persistence.Expense, error)
	GetExpense(ctx context.Context, id int32) (*persistence.Expense, error)
	GetLastExpense(ctx context.Context) (*persistence.Expense, error)
	ListExpenses(ctx context.Context) ([]*persistence.Expense, error)
	UpdateExpense(ctx context.Context, arg *persistence.UpdateExpenseParams) (*persistence.Expense, error)

	// users
	CreateUser(ctx context.Context, arg *persistence.CreateUserParams) (*persistence.User, error)
	GetLastUser(ctx context.Context) (*persistence.User, error)
	GetUserByEmail(ctx context.Context, email string) (*persistence.User, error)
	GetUser(ctx context.Context, id int32) (*persistence.User, error)
	ExistsAnUserUsingTheSameEmail(ctx context.Context, email string) (bool, error)
}

func NewRepository(repo *persistence.Queries) Repository {
	return repo
}
