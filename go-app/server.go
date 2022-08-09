package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/kaenova/go-elastic-kibana/go-app/graph"
	"github.com/kaenova/go-elastic-kibana/go-app/graph/generated"
	"github.com/kaenova/go-elastic-kibana/go-app/pkg/logger"
)

const defaultPort = "8080"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logElastic := logger.NewLogger(logger.LoggerOptions{
		AppName:     os.Getenv("APP_NAME"),
		ElasticHost: os.Getenv("ELASTICSEARCH_HOST"),
		ElasticPort: os.Getenv("ELASTICSEARCH_PORT"),
	})

	srv := handler.NewDefaultServer(generated.
		NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
			Logger: logElastic,
		}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
