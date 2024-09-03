package applications

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/directives"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/handlers/graph"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/handlers/graph/model"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/middleware"
	"github.com/isaacmirandacampos/dreamkoffee/internal/domain"
)

func Initialize(repo *domain.Repository) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.Auth, middleware.WithResponseWriter)
	c := graph.Config{Resolvers: &graph.Resolver{
		Repo:     repo,
		Expenses: []*model.Expense{},
		Users:    []*model.User{},
	},
		Directives: graph.DirectiveRoot{
			Auth: directives.Auth,
		},
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	return router
}
