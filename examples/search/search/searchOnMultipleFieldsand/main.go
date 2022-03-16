package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/search"
)

// Search for records in a given in index
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Search.Search(&search.SearchRequest{
		Index: "customers",
		Query: "name == 'John' AND starsign == 'Leo'",
	})
	fmt.Println(rsp, err)
}
