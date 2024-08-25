package graphql

import (
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
)

type Resolver struct {
	Conn     *persistence.Queries
	Expenses []*model.Expense
}
