package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/app"
)

// Delete an app
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.App.Delete(&app.DeleteRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
