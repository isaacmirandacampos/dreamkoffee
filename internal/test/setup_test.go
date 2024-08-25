package test

import (
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql"
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
	_ "github.com/lib/pq"
)

var (
	Server *httptest.Server
)

func TestMain(m *testing.M) {
	conn, closeDB := OpenPostgresConnection()
	Repo := persistence.New(conn)
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &graphql.Resolver{
			Conn:     Repo,
			Expenses: []*model.Expense{},
		},
	}))
	Server = httptest.NewServer(srv)
	defer Server.Close()
	defer closeDB()
	m.Run()
}
