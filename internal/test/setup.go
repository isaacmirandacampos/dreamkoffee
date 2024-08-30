package test

import (
	"database/sql"
	"net/http/httptest"

	"github.com/isaacmirandacampos/finkoffee/internal/applications"
	"github.com/isaacmirandacampos/finkoffee/internal/domain"
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
	repository := domain.NewRepository(repo)
	srv := applications.Initialize(&repository)

	Server = httptest.NewServer(srv)
	close = func() {
		Server.Close()
		closeDB()
	}
	return
}
