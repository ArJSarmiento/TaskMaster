package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"crud_ql/auth"
	"crud_ql/graph"
	"crud_ql/repository"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-chi/chi"
)

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	router := chi.NewRouter()

	// Connect to the database and auth
	db := repository.Connect()
	auth_instance := auth.Init()

	router.Use(auth.Middleware(db, auth_instance))
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			DB:   db,
			AUTH: auth_instance,
		},
	}))

	// Setup handlers
	stage := os.Getenv("STAGE")
	query_sub := fmt.Sprintf("/%s/query", stage)
	router.Handle("/", playground.Handler("Taskmaster", query_sub))
	router.Handle("/query", srv)

	// Create a new request using the data from the Lambda event.
	httpReq, err := http.NewRequest(http.MethodGet, req.Path, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// Use an httptest ResponseRecorder to capture the response.
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, httpReq)

	// Convert the httptest.ResponseRecorder result to an APIGatewayProxyResponse.
	return events.APIGatewayProxyResponse{
		StatusCode: rec.Code,
		Body:       rec.Body.String(),
		Headers: map[string]string{
			"Content-Type":                 "*/*",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST, GET, OPTIONS",
			"Access-Control-Allow-Headers": "Origin, Content-Type, Accept, Authorization",
		},
	}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
