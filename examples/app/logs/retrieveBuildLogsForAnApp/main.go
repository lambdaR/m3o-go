package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/app"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.App.Logs(&app.LogsRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
