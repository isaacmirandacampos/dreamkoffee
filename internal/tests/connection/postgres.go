package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ory/dockertest"
)

func OpenPostgresConnection() (db *sql.DB, close func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
		return
	}
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not ping docker: %s", err)
		return
	}

	resource, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "postgres",
			Tag:        "latest",
			Env: []string{
				"POSTGRES_PASSWORD=" + "password_test",
				"POSTGRES_USERNAME=" + "user_test",
				"POSTGRES_DATABASE=" + "dbname_test",
				"POSTGRES_DB=" + "dbname_test",
				"POSTGRES_USER=" + "user_test",
				"POSTGRES_PASSWORD=" + "password_test",
			},
		},
	)

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
		return
	}
	dsn := fmt.Sprintf("host=127.0.0.1 port=%s user=%s password=%s dbname=%s sslmode=disable",
		resource.GetPort("5432/tcp"),
		"user_test",
		"password_test",
		"dbname_test",
	)
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	close = func() {
		err = pool.Purge(resource)
		if err != nil {
			log.Fatalf("Could not purge resource: %s", err)
			return
		}
	}
	return db, close
}
