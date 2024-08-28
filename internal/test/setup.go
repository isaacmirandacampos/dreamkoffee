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

type Persistence struct {
	Repo *persistence.Queries
	DB   *sql.DB
}

func TestWithServerAndDB() (Server *httptest.Server, database *Persistence, close func()) {
	db, closeDB := connection.OpenPostgresConnection()
	repo := persistence.New(db)
	database = &Persistence{
		Repo: repo,
		DB:   db,
	}
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &graphql.Resolver{
			Conn:     repo,
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
