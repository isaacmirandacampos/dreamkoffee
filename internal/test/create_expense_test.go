package test

import (
	"context"
	"net/http"
	"testing"

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
		query := helper.FabricateMutation("createExpense", input, []string{"name", "price"})
		resp, close, err := helper.HttpRequest(query, Server.URL, "POST")
		assert.NoError(t, err)
		defer close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		ctx := context.Background()

		expenses, err := database.Repo.GetLastExpense(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, expenses)
		assert.Equal(t, "Test Expense", expenses.Name)
		assert.Equal(t, "\"100\"", utils.MarshalDecimal(expenses.Price))
	})
}
