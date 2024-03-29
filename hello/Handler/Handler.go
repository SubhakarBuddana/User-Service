package handler

import (
	"context"
	"net/http"

	Controllers "github.com/SubhakarBuddana/User-Service/hello/Controllers"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	switch request.Path {
	case "/create-user":
		response, _ = Controllers.CreateUser(ctx, request)
	case "/get-user":
		response, _ = Controllers.GetUser(ctx, request)
	case "/login-user":
		response, _ = Controllers.LoginUser(ctx, request)
	case "/delete-user":
		response, _ = Controllers.DeleteUser(ctx, request)
	case "/forgot-password":
		response, _ = Controllers.ForgotPassword(ctx, request)
	default:
		response = events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "Not Found"}
	}

	return response, nil
}
