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

func GetUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}
	defer db.Close()

	// Retrieve user from the database based on user ID
	userID := request.QueryStringParameters["id"]
	var user Models.User
	row := db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", userID)
	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		log.Printf("Failed to retrieve user from database: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 404}, nil
	}

	responseBody, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}
