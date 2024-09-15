package main

import "github.com/graphql-go/graphql"

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"authorID"`
}

var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"authorID": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

func init() {
	BookType.AddFieldConfig("author", &graphql.Field{
		Type: AuthorType,
		Args: nil,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			book, _ := p.Source.(Book)
			for _, author := range authors {
				if author.ID == book.AuthorID {
					return author, nil
				}
			}
			return nil, nil
		},
	})
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

var books = []Book{
	{ID: 1, Title: "Harry Potter and the Chamber of Secrets", AuthorID: 1},
	{ID: 2, Title: "Harry Potter and the Prisoner of Azkaban", AuthorID: 1},
	{ID: 3, Title: "Harry Potter and the Goblet of Fire", AuthorID: 1},
	{ID: 4, Title: "The Fellowship of the Ring", AuthorID: 2},
	{ID: 5, Title: "The Two Towers", AuthorID: 2},
	{ID: 6, Title: "The Return of the King", AuthorID: 2},
	{ID: 7, Title: "The Way of Shadows", AuthorID: 3},
	{ID: 8, Title: "Beyond the Shadows", AuthorID: 3},
}
