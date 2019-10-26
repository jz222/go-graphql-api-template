package main

import (
	"fmt"
	gql "graphql_api/graphQL"
	grapqhlconfig "graphql_api/graphQL/configuration"
	"graphql_api/initialization"
	"graphql_api/keys"
	"graphql_api/middlewares"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
)

var port = fmt.Sprintf(":%s", keys.GetKeys().PORT)
var router = mux.NewRouter()

/*
	Initializes environment variables and establishes
	a connection to MongoDB.
*/
func init() {
	initialization.InitEnv()
	initialization.InitDatabase()
}

func main() {
	router.Use(middlewares.VerifyJwt)

	router.Handle("/playground", handler.Playground("playground", "/graphql"))

	router.Handle("/graphql", handler.GraphQL(gql.NewExecutableSchema(grapqhlconfig.GetConfig())))

	srv := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	log.Println("🆙 Server listening on port", port)

	srv.ListenAndServe()
}
