package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/vehicle"
)

// Lookup a UK vehicle by it's registration number
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Vehicle.Lookup(&vehicle.LookupRequest{
		Registration: "LC60OTA",
	})
	fmt.Println(rsp, err)
}
