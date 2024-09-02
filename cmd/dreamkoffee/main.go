package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"

	"github.com/isaacmirandacampos/dreamkoffee/configs"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications"
	"github.com/isaacmirandacampos/dreamkoffee/internal/applications/middleware"
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

	router := mux.NewRouter()
	router.Use(middleware.Auth, middleware.WithResponseWriter)

	srv := applications.Initialize(&repository)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}
