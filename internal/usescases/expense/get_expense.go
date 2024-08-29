package expense

import (
	"context"

	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph/model"
	"github.com/isaacmirandacampos/finkoffee/internal/utils"
)

func (uc *expenseUseCase) GetExpense(ctx context.Context, id *int) (*model.Expense, error) {
	result, err := uc.repo.GetExpense(ctx, int32(*id))
	if err != nil {
		utils.ErrorHandling(ctx, 404, "expense_not_found", "Expense not found")
		return nil, nil
	}

	expense := &model.Expense{
		ID:          int(result.ID),
		Description: result.Description,
		Value:       result.Value,
		CreatedAt:   result.CreatedAt.String(),
		UpdatedAt:   result.UpdatedAt.String(),
	}

	return expense, nil
}
