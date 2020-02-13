package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/ericatcorsha/rpc-compare/servers/graphql"
)

func main() {
	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = ":7080"
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}})))

	log.Printf("connect to http://%s/ for GraphQL playground", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}
