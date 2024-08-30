package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/isaacmirandacampos/finkoffee/configs"
	"github.com/isaacmirandacampos/finkoffee/internal/applications"
	"github.com/isaacmirandacampos/finkoffee/internal/domain"
	"github.com/isaacmirandacampos/finkoffee/internal/infrastructure/database"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
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
