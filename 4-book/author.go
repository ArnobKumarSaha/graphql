package main

import "github.com/graphql-go/graphql"

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var AuthorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// If we added books & author field directly in AuthorType & BookType respectively , that will show the `initialization cycle error`
// To solve that, we used AddFieldConfig technique (https://stackoverflow.com/questions/62031840/creating-cyclical-graphql-types-in-go)

func init() {
	AuthorType.AddFieldConfig("books", &graphql.Field{
		Type: graphql.NewList(BookType),
		Args: nil,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			author, _ := p.Source.(Author)
			var res []Book

			for _, book := range books {
				if book.AuthorID == author.ID {
					res = append(res, book)
				}
			}
			return res, nil
		},
	})
}

var authors = []Author{
	{ID: 1, Name: "J. K. Rowling"},
	{ID: 2, Name: "J. R. R. Tolkien"},
	{ID: 3, Name: "Brent Weeks"},
}
