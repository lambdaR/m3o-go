package main

import (
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Update a record in the database. Include an "id" in the record to update.
func main() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Update(&db.UpdateRequest{
		Record: map[string]interface{}{
			"id":  "1",
			"age": 43,
		},
		Table: "example",
	})
	fmt.Println(rsp, err)
}
