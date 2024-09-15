package main

import (
	"encoding/json"
	"fmt"
	"github.com/Arnobkumarsaha/graphql/utils"
	"github.com/graphql-go/graphql"
)

func main() {
	tutorials := populate()
	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type:        tutorialType,
			Description: "Get Tutorial By ID",
			// We can define arguments that allow us to
			// pick specific tutorials. In this case
			// we want to be able to specify the ID of the
			// tutorial we want to retrieve
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// take in the ID argument
				id, ok := p.Args["id"].(int)
				if ok {
					// Parse our tutorial array for the matching id
					for _, tutorial := range tutorials {
						if tutorial.ID == id {
							return tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Tutorial List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}

	schema := utils.MakeNewSchema(fields)
	testListField(schema)
	testTutorialField(schema)
}

func testListField(schema graphql.Schema) {
	query := `
    {
        list {
            id
            title
            comments {
                body
            }
            author {
                Name
                Tutorials
            }
        }
    }
`
	res := utils.MakeQuery(schema, query)
	rJSON, _ := json.Marshal(res.Data)
	fmt.Printf("%s \n", rJSON)
}

func testTutorialField(schema graphql.Schema) {
	query := `
    {
        tutorial(id:1) {
            title
            author {
                Name
            }
        }
    }
`
	res := utils.MakeQuery(schema, query)
	rJSON, _ := json.Marshal(res.Data)
	fmt.Printf("%s \n", rJSON)
}
