package test

import (
	"database/sql"
	"net/http/httptest"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql"
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/finkoffee/internal/test/connection"
	_ "github.com/lib/pq"
)

func TestWithServerAndDB() (Server *httptest.Server, DB *sql.DB, close func()) {
	DB, closeDB := connection.OpenPostgresConnection()
	var err error
	if err != nil {
		panic(err)
	}
	Repo := persistence.New(DB)

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
