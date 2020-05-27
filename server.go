package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/adelowo/url-shortner/datastore/postgres"
	"github.com/adelowo/url-shortner/graph"
	"github.com/adelowo/url-shortner/graph/generated"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func getDSN() string {
	return os.Getenv("POSTGRESQL_DSN")
}

func main() {

	var port string
	var postgresDSN string

	flag.StringVar(&port, "http.server", defaultPort, "port to run HTTP server on")

	flag.Parse()

	if port == "" {
		port = defaultPort
	}

	postgresDSN = getDSN()
	if len(strings.Trim(postgresDSN, " ")) == 0 {
		log.Fatal("Please provide your postgres DSN")
	}

	postgres, err := postgres.New(postgresDSN)
	if err != nil {
		log.Fatalf("could not set up postgres connection... %v", err)
	}

	_ = postgres

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
