package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"crud_ql/auth"
	"crud_ql/graph"
	"crud_ql/repository"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

var muxRouter *mux.Router

func init() {
	muxRouter = mux.NewRouter()

	// Connect to the database and auth
	db := repository.Connect()
	auth_instance := auth.Init()

	muxRouter.Use(auth.Middleware(db, auth_instance))

	schema := graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			DB:   db,
			AUTH: auth_instance,
		},
	})
	server := handler.NewDefaultServer(schema)
	muxRouter.Handle("/query", server)
	muxRouter.Handle("/", playground.Handler("GraphQL playground", "/query"))
}

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	muxAdapter := gorillamux.New(muxRouter)
	rsp, err := muxAdapter.Proxy(*core.NewSwitchableAPIGatewayRequestV1(&req))
	if err != nil {
		log.Println(err)
	}
	newRsp := &events.APIGatewayProxyResponse{
		Body:       rsp.Version1().Body,
		StatusCode: rsp.Version1().StatusCode,
		Headers:    rsp.Version1().Headers,
	}
	return *newRsp, err
}

func main() {
	isRunningAtLambda := strings.Contains(os.Getenv("AWS_EXECUTION_ENV"), "AWS_Lambda_")

	if isRunningAtLambda {
		lambda.Start(lambdaHandler)
	} else {
		defaultPort := "7010"
		port := os.Getenv("PORT")

		if port == "" {
			port = defaultPort
		}

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, muxRouter))
	}
}
