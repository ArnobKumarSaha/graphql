package main

import (
	"fmt"
	"github.com/Arnobkumarsaha/graphql/utils"
	"net/http"

	"github.com/graphql-go/handler"
)

/*
To try the simple example (users)
> cat query.json
{
    "query": "{
        user(id: 2) {
            id
            name
        }
    }"
}


To try the book-author example
> cat query.json
{
    "query": "{
        author(id: 2) {
            id
            name
            books {
                title
            }
        }
        book(id: 1) {
            id
            title
            author {
                name
            }
        }
    }"
}

> In the above example, `author(id)` & `book(id)` is being resolved in fields.go
> But, the internal `books {}` & `author {}` is being resolved in types.go file

# Request
curl -X POST -H "Content-Type: application/json" -d @query.json http://localhost:8080/graphql

*/

func main() {
	fields := getAuthorAndBookFields() // Or try simple -> getUserFields()
	schema := utils.MakeNewSchema(fields)

	// Create a GraphQL handler
	handler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// Serve the GraphQL API
	http.Handle("/graphql", handler)
	fmt.Println("Server is running on port 8080")
	_ = http.ListenAndServe(":8080", nil)
}
