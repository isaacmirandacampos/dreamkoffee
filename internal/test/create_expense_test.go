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
	Server, DB, close := TestWithServerAndDB()
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
		if err != nil {
			t.Fatalf("Could not perform HTTP request: %v", err)
		}
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		result := &model.Expense{}
		err = DB.QueryRow("SELECT name, price FROM expenses").Scan(&result.Name, &result.Price)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Test Expense", result.Name)
		assert.Equal(t, "\"100\"", utils.MarshalDecimal(result.Price))
	})
}
