package test

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql"
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/finkoffee/internal/test/connection"
	_ "github.com/lib/pq"
)

var (
	Server *httptest.Server
	DB *sql.DB
)

func TestMain(m *testing.M) {
	var closeDB func()
	DB, closeDB = connection.OpenPostgresConnection()
	Repo := persistence.New(DB)
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
