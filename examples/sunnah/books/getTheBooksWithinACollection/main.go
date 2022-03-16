package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/sunnah"
)

// Get a list of books from within a collection. A book can contain many chapters
// each with its own hadiths.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Sunnah.Books(&sunnah.BooksRequest{
		Collection: "bukhari",
	})
	fmt.Println(rsp, err)
}
