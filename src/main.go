package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(event events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("%s - Successful", event.HTTPMethod),
	}
}

func main() {
	runtime.Start(handleRequest)
}
