package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/stream"
)

// Create a channel with a given name and description. Channels are created automatically but
// this allows you to specify a description that's persisted for the lifetime of the channel.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Stream.CreateChannel(&stream.CreateChannelRequest{
		Description: "The channel for all things",
		Name:        "general",
	})
	fmt.Println(rsp, err)
}
