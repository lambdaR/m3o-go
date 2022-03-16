package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/space"
)

// List the objects in space
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Space.List(&space.ListRequest{
		Prefix: "images/",
	})
	fmt.Println(rsp, err)
}
