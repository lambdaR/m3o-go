package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/comments"
)

// Delete a comment
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Comments.Delete(&comments.DeleteRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",
	})
	fmt.Println(rsp, err)
}
