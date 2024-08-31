package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/isaacmirandacampos/dreamkoffee/configs"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications"
	"github.com/isaacmirandacampos/dreamkoffee/internal/domain"
	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database"
	"github.com/isaacmirandacampos/dreamkoffee/internal/storage/persistence"
)

const defaultPort = "8080"

func main() {
	configs.Initialize()
	connection, err := database.Initialize()
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	repo := persistence.New(connection)
	repository := domain.NewRepository(repo)
	srv := applications.Initialize(&repository)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
