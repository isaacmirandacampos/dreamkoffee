package e2e_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres/persistence"
	"github.com/isaacmirandacampos/dreamkoffee/internal/test"
	"github.com/isaacmirandacampos/dreamkoffee/internal/test/helper"
	"github.com/isaacmirandacampos/dreamkoffee/pkg/scalar"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	t.Parallel()
	Server, database, close := test.TestWithServerAndDB()
	defer close()
	t.Run("create_a_new_expense", func(t *testing.T) {
		user, err := database.Repo.CreateUser(context.Background(), &persistence.CreateUserParams{
			FullName: "Test User",
			Email:    "teste@user.com",
		})
		assert.NoError(t, err)
		input := map[string]interface{}{
			"value":       "100.00",
			"description": "Test Expense",
			"userId":      user.ID,
		}
		query := helper.FabricateMutation("createExpense", input, []string{"description", "value"})
		resp, close, err := helper.HttpRequest(query, Server.URL, "POST")
		assert.NoError(t, err)
		defer close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		ctx := context.Background()

		expenses, err := database.Repo.GetLastExpense(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, expenses)
		assert.Equal(t, "Test Expense", expenses.Description)
		assert.Equal(t, "\"100\"", scalar.MarshalDecimal(expenses.Value))
	})
}
