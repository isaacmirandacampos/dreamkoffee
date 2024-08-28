package test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
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
			Name:  "Test Expense",
			Price: price,
		})
		if err != nil {
			t.Fatalf("Could not create expense: %v", err)
		}

		query := `{
			"query": "query { getExpense(id: 1) { id name } }"
		}`

		req, err := http.NewRequest("POST", Server.URL+"/query", bytes.NewBuffer([]byte(query)))
		if err != nil {
			t.Fatalf("Could not create HTTP request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Could not execute HTTP request: %v", err)
		}
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Could not read HTTP response: %v", err)
		}

		var response struct {
			Data struct {
				GetExpense persistence.Expense `json:"getExpense"`
			} `json:"data"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatalf("Could not unmarshal JSON response: %v", err)
		}
		expense := response.Data.GetExpense
		assert.Equal(t, result.ID, expense.ID)
		assert.Equal(t, "Test Expense", expense.Name)
	})
}
