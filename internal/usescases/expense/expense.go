package expense

import (
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
)

type expenseUseCase struct {
	repo *persistence.Queries
}

func NewExpenseUseCase(repo *persistence.Queries) *expenseUseCase {
	return &expenseUseCase{repo: repo}
}
