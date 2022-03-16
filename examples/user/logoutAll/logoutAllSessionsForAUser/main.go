package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

// Logout of all user's sessions
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.LogoutAll(&user.LogoutAllRequest{})
	fmt.Println(rsp, err)
}
