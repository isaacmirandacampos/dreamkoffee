package applications

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph"
	"github.com/isaacmirandacampos/finkoffee/internal/applications/graph/model"
	"github.com/isaacmirandacampos/finkoffee/internal/domain"
)

func Initialize(repo *domain.Repository) (srv *handler.Server) {

	srv = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Repo:     repo,
		Expenses: []*model.Expense{},
	}}))

	return
}
