package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/stock"
)

// Get the last quote for the stock
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Stock.Quote(&stock.QuoteRequest{
		Symbol: "AAPL",
	})
	fmt.Println(rsp, err)
}
