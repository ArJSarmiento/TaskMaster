package main

import (
	"log"
	"net/http"

	"crud_ql/auth"
	"crud_ql/graph"
	"crud_ql/repository"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func main() {
	defaultPort := repository.GoDotEnvVariable("DEFAULT_PORT")
	router := chi.NewRouter()

	db := repository.Connect()
	auth_instance := auth.Init()

	router.Use(auth.Middleware(db, auth_instance))
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			DB:   db,
			AUTH: auth_instance,
		},
	}))
	router.Handle("/", playground.Handler("Taskmaster", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}
