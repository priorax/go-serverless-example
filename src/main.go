package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

type Resp struct {
	Body string
}

func handleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := Resp{
		Body: fmt.Sprintf("%s - Successful", event.HTTPMethod),
	}
	msg, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(msg),
	}, nil
}

func main() {
	runtime.Start(handleRequest)
}
