package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/kentliuqiao/go-graphql/gql"
	"github.com/kentliuqiao/go-graphql/postgres"
	"github.com/kentliuqiao/go-graphql/server"
)

func main() {
	// Initialize our api and return a pointer to our router for http.ListenAndServe
	// and a pointer to our db to defer its closing when main() is finished
	router, db := initializeAPI()
	defer db.Close()

	// Listen on port 4000 and if there's an error log it and exit
	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *postgres.DB) {
	db, err := postgres.New(postgres.ConnString("localhost", 5432, "postgres", "root", "go_graphql_db"))
	if err != nil {
		log.Fatal(err)
	}
	rootQuery := gql.NewRoot(db)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery.Query})
	if err != nil {
		log.Fatal(err)
	}

	server := server.Server{GQLSchema: &schema}

	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/graphql", server.GraphQL())

	return router, db
}
