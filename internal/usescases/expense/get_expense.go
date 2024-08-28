package expense

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
)

func (uc *expenseUseCase) GetExpenses(ctx context.Context, id *int) (*model.Expense, error) {
	result, err := uc.repo.GetExpense(ctx, int32(*id))
	if err != nil {
		return nil, err
	}

	expense := &model.Expense{
		ID:        int(result.ID),
		Name:      result.Name,
		Price:     result.Price,
		CreatedAt: result.CreatedAt.String(),
		UpdatedAt: result.UpdatedAt.String(),
	}

	return expense, nil
}
