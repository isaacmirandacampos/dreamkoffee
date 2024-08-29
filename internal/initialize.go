package internal

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph"
	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
)

func Initialize(conn *persistence.Queries) (srv *handler.Server) {

	srv = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Conn:     conn,
		Expenses: []*model.Expense{},
	}}))

	return
}
