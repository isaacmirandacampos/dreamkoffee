package test

import (
	"net/http/httptest"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql"
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/finkoffee/internal/test/connection"
	_ "github.com/lib/pq"
)

func TestWithServerAndDB() (Server *httptest.Server, Repo *persistence.Queries, close func()) {
	db, closeDB := connection.OpenPostgresConnection()
	var err error
	if err != nil {
		panic(err)
	}

	Repo = persistence.New(db)

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &graphql.Resolver{
			Conn:     Repo,
			Expenses: []*model.Expense{},
		},
	}))

	Server = httptest.NewServer(srv)
	close = func() {
		Server.Close()
		closeDB()
	}
	return
}
