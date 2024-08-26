package test

import (
	"bytes"
	"log"
	"net/http"
	"testing"

	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	t.Parallel()
	Server, Repo, close := TestWithServerAndDB()
	defer close()
	t.Run("create_a_new_expense", func(t *testing.T) {
		query := `{
			"query": "mutation { createExpense(input: {name: \"Test Expense\", price: \"100.00\"}) { name price } }"
		}`
		log.Printf("HTTP request: %v", Server.URL)

		req, err := http.NewRequest("POST", Server.URL+"/query", bytes.NewBuffer([]byte(query)))
		if err != nil {
			t.Fatalf("Could not create HTTP request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.NoError(t, err)
		if err != nil {
			t.Fatalf("Could not read HTTP response: %v", err)
		}
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		results, err := Repo.ListExpenses(req.Context())
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
