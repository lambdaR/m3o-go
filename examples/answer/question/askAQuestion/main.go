package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/answer"
)

// Ask a question and receive an instant answer
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Answer.Question(&answer.QuestionRequest{
		Query: "microsoft",
	})
	fmt.Println(rsp, err)
}
