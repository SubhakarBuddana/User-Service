package controllers

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

func DeleteUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}
	defer db.Close()

	// Retrieve user ID from request path parameters
	userID := request.PathParameters["id"]

	// Delete user from the database based on user ID
	result, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		log.Printf("Failed to delete user from database: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	if rowsAffected == 0 {
		return events.APIGatewayProxyResponse{StatusCode: 404}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 204}, nil
}
