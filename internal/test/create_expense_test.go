package test

import (
	"bytes"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	t.Run("create_a_new_expense", func(t *testing.T){
		query := `{
			"query": "mutation { createExpense(input: {name: \"Test Expense\", price: 100}) { name price } }"
		}`
		log.Printf("HTTP request: %v", Server.URL)

		req, err := http.NewRequest("POST", Server.URL+"/query", bytes.NewBuffer([]byte(query)))
		if err != nil {
			t.Fatalf("Could not create HTTP request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Could not perform HTTP request: %v", err)
		}
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
