package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/test/helper"
	"github.com/isaacmirandacampos/finkoffee/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	t.Parallel()
	Server, database, close := TestWithServerAndDB()
	defer close()
	t.Run("create_a_new_expense", func(t *testing.T) {
		input := map[string]string{
			"price": "100.00",
			"name":  "Test Expense",
		}
		query := helper.QueryMutation("createExpense", input, []string{"name", "price"})
		resp, close, err := helper.HttpRequest(query, Server.URL, "POST")
		assert.NoError(t, err)
		defer close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		ctx := context.Background()
		results, err := database.Repo.ListExpenses(ctx)
		expenses := make([]*model.Expense, 0, len(results))
		assert.NoError(t, err)

		for _, r := range results {
			expenses = append(expenses, &model.Expense{
				ID:        int(r.ID),
				Name:      r.Name,
				Price:     r.Price,
				CreatedAt: r.CreatedAt.String(),
				UpdatedAt: r.UpdatedAt.String(),
			})
		}

		assert.NotNil(t, expenses)
		assert.Equal(t, 1, len(expenses))
		result := expenses[0]
		assert.Equal(t, "Test Expense", result.Name)
		assert.Equal(t, "\"100\"", utils.MarshalDecimal(result.Price))
	})
}
