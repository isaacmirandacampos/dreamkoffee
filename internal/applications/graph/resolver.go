package graph

import (
	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Conn     *persistence.Queries
	Expenses []*model.Expense
}
