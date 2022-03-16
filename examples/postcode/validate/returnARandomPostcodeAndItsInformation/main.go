package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/postcode"
)

// Validate a postcode.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Postcode.Validate(&postcode.ValidateRequest{
		Postcode: "SW1A 2AA",
	})
	fmt.Println(rsp, err)
}
