package main

import "github.com/graphql-go/graphql"

type User struct {
	ID   int    `json:"id"`
	Naam string `json:"naam"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	// user -> ur, naam -> name
	// Intentionally making name mismatch to show that GraphQL knows nothing about the above struct.
	Name: "ur",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func getUserFields() graphql.Fields {
	return graphql.Fields{
		"usr": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if !ok {
					return nil, nil
				}
				for _, user := range users {
					if user.ID == id {
						// the returned value has to be of `UserType` , not `User`.
						// As we set that on field.Type
						s := struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
						}{
							ID:   user.ID,
							Name: user.Naam,
						}
						return s, nil
					}
				}
				return nil, nil
			},
		},
	}
}

var users = []User{
	User{ID: 1, Naam: "Alice"},
	User{ID: 2, Naam: "Bob"},
	User{ID: 3, Naam: "Charlie"},
}
