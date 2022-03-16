package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/rss"
)

// Get an RSS feed by name. If no name is given, all feeds are returned. Default limit is 25 entries.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Rss.Feed(&rss.FeedRequest{
		Name: "bbc",
	})
	fmt.Println(rsp, err)
}
