package utils

import (
	"github.com/graphql-go/graphql"
	"log"
)

func MakeNewSchema(fields graphql.Fields) graphql.Schema {
	gObj := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: gObj,
	})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}

func MakeQuery(schema graphql.Schema, query string) *graphql.Result {
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return r
}
