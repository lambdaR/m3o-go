package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/search"
)

// Delete an index by name
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Search.DeleteIndex(&search.DeleteIndexRequest{
		Index: "customers",
	})
	fmt.Println(rsp, err)
}
