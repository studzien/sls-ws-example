package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

type response events.APIGatewayProxyResponse
type request events.APIGatewayWebsocketProxyRequest

func handler(ctx context.Context, request request) (response, error) {
	body := request.Body
	log.Printf("Message received: %v", body)

	requestContext := request.RequestContext

	session := session.Must(session.NewSession())
	config := aws.NewConfig().WithEndpoint(requestContext.DomainName + "/" + requestContext.Stage)
	apiGateway := apigatewaymanagementapi.New(session, config)
	input := apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: &requestContext.ConnectionID,
		Data:         []byte(body)}

	_, err := apiGateway.PostToConnection(&input)
	if err != nil {
		return response{StatusCode: 500}, err
	}

	return response{StatusCode: 200, Body: body}, nil
}

func main() {
	lambda.Start(handler)
}
