# Answer

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/answer/api](https://m3o.com/answer/api).

Endpoints:

## Question

Ask a question and receive an instant answer


[https://m3o.com/answer/api#Question](https://m3o.com/answer/api#Question)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/answer"
)

// Ask a question and receive an instant answer
func AskAquestion() {
	answerService := answer.NewAnswerService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := answerService.Question(&answer.QuestionRequest{
		Query: "microsoft",

	})
	fmt.Println(rsp, err)
	
}
```
