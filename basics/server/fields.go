package main

import "github.com/graphql-go/graphql"

func getUserFields() graphql.Fields {
	return graphql.Fields{
		"user": &graphql.Field{
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
						return user, nil
					}
				}
				return nil, nil
			},
		},
	}
}

func getAuthorAndBookFields() graphql.Fields {
	return graphql.Fields{
		"book": &graphql.Field{
			Type: BookType,
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
				for _, book := range books {
					if book.ID == id {
						return book, nil
					}
				}
				return nil, nil
			},
		},
		"books": &graphql.Field{
			Type: &graphql.List{OfType: BookType},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return books, nil
			},
		},
		"author": &graphql.Field{
			Type: AuthorType,
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
				for _, author := range authors {
					if author.ID == id {
						return author, nil
					}
				}
				return nil, nil
			},
		},
		"authors": &graphql.Field{
			Type: &graphql.List{OfType: AuthorType},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return authors, nil
			},
		},
	}
}
