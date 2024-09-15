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

# Request
curl -X POST -H "Content-Type: application/json" -d @query.json http://localhost:8080/graphql

*/

func main() {
	fields := getUserFields()
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
