package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/isaacmirandacampos/finkofee/internal/interface/graphql"
	"github.com/isaacmirandacampos/finkofee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkofee/pkg/config"
)

const defaultPort = "8080"

func main() {
	config.Initialize()
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{
		ListTransactions: []*model.Transaction{},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
