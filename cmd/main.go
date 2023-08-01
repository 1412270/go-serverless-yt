package main

import "github.com/aws/aws-lambda-go/events"

func main() {

}

const tableName = "LambdaInGoUser"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handler
	}
}
