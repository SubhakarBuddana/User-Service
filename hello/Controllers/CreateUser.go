package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/lib/pq"
)

func CreateUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

	// Insert user into the database
	_, err = db.Exec("INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("Failed to insert user into database: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	responseBody, _ := json.Marshal(map[string]string{"message": "User created successfully"})
	return events.APIGatewayProxyResponse{StatusCode: 201, Body: string(responseBody)}, nil
}
