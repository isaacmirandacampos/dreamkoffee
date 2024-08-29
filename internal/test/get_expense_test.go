package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/finkoffee/internal/test/helper"
	"github.com/isaacmirandacampos/finkoffee/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetExpense(t *testing.T) {
	t.Parallel()
	Server, database, close := TestWithServerAndDB()
	defer close()
	t.Run("get_a_existent_expense", func(t *testing.T) {
		price, err := utils.UnmarshalDecimal(100)
		if err != nil {
			t.Fatalf("Could not unmarshal decimal: %v", err)
		}
		ctx := context.Background()
		result, err := database.Repo.CreateExpense(ctx, persistence.CreateExpenseParams{
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
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var response struct {
			Data struct {
				GetExpense persistence.Expense `json:"getExpense"`
			} `json:"data"`
		}
		err = helper.TransformBody(resp.Body, &response)
		assert.NoError(t, err)
		expense := response.Data.GetExpense
		assert.Equal(t, result.ID, expense.ID)
		assert.Equal(t, "Test Expense", expense.Description)
	})
}
