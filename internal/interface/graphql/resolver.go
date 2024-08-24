package graphql

import "github.com/isaacmirandacampos/finkofee/internal/interface/graphql/model"

type Resolver struct {
	ListTransactions []*model.Transaction
}
