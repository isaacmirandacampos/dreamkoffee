package expense

import "github.com/isaacmirandacampos/dreamkoffee/internal/domain"

type expenseUseCase struct {
	repo domain.Repository
}

func NewExpenseUseCase(repo *domain.Repository) *expenseUseCase {
	return &expenseUseCase{repo: *repo}
}
