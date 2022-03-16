package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/function"
)

// Update a function. Downloads the source, builds and redeploys
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Function.Update(&function.UpdateRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
