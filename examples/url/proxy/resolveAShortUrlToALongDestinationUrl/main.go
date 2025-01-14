package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/url"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Url.Proxy(&url.ProxyRequest{})
	fmt.Println(rsp, err)
}
