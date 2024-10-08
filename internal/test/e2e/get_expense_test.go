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

func TestGetExpense(t *testing.T) {
	t.Parallel()
	Server, database, close := test.TestWithServerAndDB()
	defer close()
	type dataResponse struct {
		GetExpense persistence.Expense `json:"getExpense"`
	}
	t.Run("get_a_existent_expense", func(t *testing.T) {
		price, err := scalar.UnmarshalDecimal(100)
		if err != nil {
			t.Fatalf("Could not unmarshal decimal: %v", err)
		}
		ctx := context.Background()
		result, err := database.Repo.CreateExpense(ctx, &persistence.CreateExpenseParams{
			Description: "Test Expense",
			Value:       price,
		})
		if err != nil {
			t.Fatalf("Could not create expense: %v", err)
		}

		query := `{
			"query": "query { getExpense(id: 1) { id description } }"
		}`

		resp, close, err := helper.HttpRequest(query, Server.URL, "POST")
		assert.NoError(t, err)
		defer close()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var response struct {
			Errors helper.ResponseError `json:"errors"`
			Data   *dataResponse        `json:"data"`
		}
		err = helper.TransformBody(resp.Body, &response)
		assert.NoError(t, err)
		expense := response.Data.GetExpense
		assert.Equal(t, result.ID, expense.ID)
		assert.Equal(t, "Test Expense", expense.Description)
		assert.Empty(t, response.Errors)
	})

	t.Run("get_a_non_existent_expense", func(t *testing.T) {
		query := `{
			"query": "query { getExpense(id: 100) { id description } }"
		}`

		resp, close, err := helper.HttpRequest(query, Server.URL, "POST")
		assert.NoError(t, err)
		defer close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var response struct {
			Errors helper.ResponseError `json:"errors"`
			Data   *dataResponse        `json:"data"`
		}
		err = helper.TransformBody(resp.Body, &response)
		assert.NoError(t, err)
		assert.Equal(t, "Expense not found", response.Errors[0].Message)
		assert.Equal(t, "EXPENSE_NOT_FOUND", response.Errors[0].Extensions.Error)
		assert.Equal(t, 404, response.Errors[0].Extensions.StatusCode)
	})
}
