package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kentliuqiao/go-graphql/postgres"
)

type Resolver struct {
	db *postgres.DB
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	name, ok := p.Args["name"].(string)
	if ok {
		users, err := r.db.GetUserByName(name)
		return users, err
	}

	return nil, nil
}
