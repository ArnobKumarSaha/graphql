package main

import "github.com/graphql-go/graphql"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	User{ID: 1, Name: "Alice"},
	User{ID: 2, Name: "Bob"},
	User{ID: 3, Name: "Charlie"},
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"authorID"`
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

var authors = []Author{
	{ID: 1, Name: "J. K. Rowling"},
	{ID: 2, Name: "J. R. R. Tolkien"},
	{ID: 3, Name: "Brent Weeks"},
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
