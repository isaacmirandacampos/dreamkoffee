package expense

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph/model"
)

func (uc *expenseUseCase) GetExpense(ctx context.Context, id *int) (*model.Expense, error) {
	result, err := uc.repo.GetExpense(ctx, int32(*id))
	if err != nil {
		return nil, err
	}

	if result.ID == 0 {
		err = fmt.Errorf("Expense not found")
		graphql.AddError(ctx, err)
		return nil, err
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
