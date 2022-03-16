package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/notes"
)

// Subscribe to notes events
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	stream, err := client.Notes.Events(&notes.EventsRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(rsp)
	}
}
