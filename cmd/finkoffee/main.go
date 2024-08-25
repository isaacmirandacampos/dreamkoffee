package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/isaacmirandacampos/finkoffee/internal/infrastructure/database"
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql"
	"github.com/isaacmirandacampos/finkoffee/internal/interface/graphql/model"
	"github.com/isaacmirandacampos/finkoffee/internal/storage/persistence"
	"github.com/isaacmirandacampos/finkoffee/pkg/config"
)

const defaultPort = "8080"

func main() {
	config.Initialize()
	connection, err := database.Initialize()
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	conn := persistence.New(connection)

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{
		Conn:     conn,
		Expenses: []*model.Expense{},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
