package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response events.APIGatewayProxyResponse
type request events.APIGatewayWebsocketProxyRequest

func handler(ctx context.Context, request request) (response, error) {
	requestContext := request.RequestContext
	log.Printf("Websocket connected, connection id: %v", requestContext.ConnectionID)

	resp := response{StatusCode: 200}
	return resp, nil
}

func main() {
	lambda.Start(handler)
}
