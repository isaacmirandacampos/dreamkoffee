package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
)

func Initialize() (*sql.DB, error) {
	drive := os.Getenv("DRIVER")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	conn, err := sql.Open(drive, dsn)

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err.Error())
	}

	return conn, nil
}
