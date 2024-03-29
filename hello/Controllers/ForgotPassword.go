package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	Models "github.com/SubhakarBuddana/User-Service/hello/Models"
	"github.com/aws/aws-lambda-go/events"
)

func ForgotPassword(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse request body
	var user Models.User
	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		log.Printf("Failed to unmarshal request body: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// Implement forgot password functionality (dummy implementation for demonstration)
	resetToken := "dummy_reset_token"
	resetLink := fmt.Sprintf("https://example.com/reset?token=%s", resetToken)

	responseBody, _ := json.Marshal(map[string]string{"reset_link": resetLink})
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}
