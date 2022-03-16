package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/currency"
)

// Returns the historic rates for a currency on a given date
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Currency.History(&currency.HistoryRequest{
		Code: "USD",
		Date: "2021-05-30",
	})
	fmt.Println(rsp, err)
}
