package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

// Read an account by id, username or email. Only one need to be specified.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.Read(&user.ReadRequest{
		Id: "user-1",
	})
	fmt.Println(rsp, err)
}
