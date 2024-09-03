package main

import (
	"log"
	"net/http"

	"github.com/isaacmirandacampos/dreamkoffee/configs"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications"
	"github.com/isaacmirandacampos/dreamkoffee/internal/domain"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres/persistence"
)

const defaultPort = "8080"

func main() {
	configs.Initialize()
	connection, err := postgres.Initialize()
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	repo := persistence.New(connection)
	repository := domain.NewRepository(repo)
	router := applications.Initialize(&repository)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}
