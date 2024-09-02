package user

import "github.com/isaacmirandacampos/dreamkoffee/internal/domain"

type userUseCase struct {
	repo domain.Repository
}

func NewExpenseUseCase(repo *domain.Repository) *userUseCase {
	return &userUseCase{repo: *repo}
}
