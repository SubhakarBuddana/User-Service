package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"

	Models "github.com/SubhakarBuddana/User-Service/hello/Models"
	"github.com/aws/aws-lambda-go/events"
)

func LoginUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse request body
	var user Models.User
	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		log.Printf("Failed to unmarshal request body: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}
	defer db.Close()

	// Validate user credentials (dummy validation for demonstration)
	validUser := Models.User{Username: "admin", Password: "password"}
	if user.Username != validUser.Username || user.Password != validUser.Password {
		return events.APIGatewayProxyResponse{StatusCode: 401}, nil
	}

	// Generate JWT token (dummy token for demonstration)
	token := "dummy_jwt_token"

	responseBody, _ := json.Marshal(map[string]string{"token": token})
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}
