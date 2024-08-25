package tests

import (
	"database/sql"
	"os"
	"testing"

	"github.com/isaacmirandacampos/finkoffee/internal/tests/connection"
	_ "github.com/lib/pq"
)

var db *sql.DB

func TestMain(m *testing.M) {
	database, close := connection.OpenPostgresConnection()
	defer close()
	db = database
	migration(database)
	code := m.Run()
	os.Exit(code)
}
