package connection

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Import the file source driver
)

func migration(database *sql.DB) {
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create postgres driver: %s", err)
	}

	cwd, _ := os.Getwd()
	migrationPath := filepath.Join(cwd, "../../infrastructure/database/postgres/migration")
	migrationSource := "file://" + migrationPath

	migrator, err := migrate.NewWithDatabaseInstance(migrationSource, "postgres", driver)
	if err != nil {
		log.Fatalf("Could not create migrator: %s", err)
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run migrations: %s", err)
	}
}
