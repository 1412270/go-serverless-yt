package handlers

import (
	"net/http"

	"github.com/1412270/go-serverless-yt/pkg/user"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		user.FetchUser
	}
}

func CreateUser() {

}

func UpdateUser() {

}

func DeleteUser() {

}

func UnhandleMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
