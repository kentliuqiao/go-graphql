package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/kentliuqiao/go-graphql/gql"
)

type Server struct {
	GQLSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

func (s *Server) GraphQL() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(rw, "Must provide graphql query in request body", http.StatusBadRequest)
			return
		}

		var rBody reqBody
		if err := json.NewDecoder(r.Body).Decode(&rBody); err != nil {
			http.Error(rw, "Error parsing JSON request body", http.StatusBadRequest)
			return
		}

		result := gql.ExecuteQuery(rBody.Query, *s.GQLSchema)

		render.JSON(rw, r, result)
	}
}
