package expense

import (
	"context"

	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/error_handling"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/handlers/graph/model"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/middleware"
	"github.com/isaacmirandacampos/dreamkoffee/internal/domain/entity"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres/persistence"
)

func (c *expenseUseCase) CreateExpense(ctx context.Context, input model.NewExpense) (*model.Expense, error) {
	userid, err := middleware.GetUserID(ctx)
	if err != nil {
		return nil, error_handling.Graphql(ctx, 401, "unauthorized", "Unauthorized", err)
	}
	expenseEntity := entity.NewExpense(
		&entity.Expense{
			Description: input.Description,
			Value:       input.Value,
		},
	)
	err = expenseEntity.ValueIsValid()
	if err != nil {
		return nil, error_handling.Graphql(ctx, 400, "value_must_be_positive", "Value must be positive", err.Error())
	}

	returned, err := c.repo.CreateExpense(ctx, &persistence.CreateExpenseParams{
		Description: expenseEntity.Description,
		Value:       expenseEntity.Value,
		UserID:      userid,
	})
	if err != nil {
		return nil, error_handling.Graphql(ctx, 500, "INTERNAL_SERVER_ERROR", "Could not create expense", err.Error())
	}
	expense := &model.Expense{
		ID:          int(returned.ID),
		Value:       returned.Value,
		Description: returned.Description,
		CreatedAt:   returned.CreatedAt.String(),
		UpdatedAt:   returned.UpdatedAt.String(),
	}
	return expense, nil
}
