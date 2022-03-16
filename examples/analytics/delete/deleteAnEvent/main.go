package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/analytics"
)

// Delete an event
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Analytics.Delete(&analytics.DeleteRequest{
		Name: "click",
	})
	fmt.Println(rsp, err)
}
