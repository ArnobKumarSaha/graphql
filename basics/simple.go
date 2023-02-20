package main

import (
	"encoding/json"
	"fmt"
	"github.com/Arnobkumarsaha/graphql/utils"
	"github.com/graphql-go/graphql"
)

func main() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	schema := utils.MakeNewSchema(fields)
	query := `
		{
			hello
		}
	`
	res := utils.MakeQuery(schema, query)
	rJSON, _ := json.Marshal(res.Data)
	fmt.Printf("%s \n", rJSON) // {"hello":"world"}
}
