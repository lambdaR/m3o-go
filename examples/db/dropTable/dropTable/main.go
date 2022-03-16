package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/db"
)

// Drop a table in the DB
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Db.DropTable(&db.DropTableRequest{
		Table: "example",
	})
	fmt.Println(rsp, err)
}
