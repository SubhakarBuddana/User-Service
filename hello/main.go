package main

import (
	Handler "github.com/SubhakarBuddana/User-Service/hello/Handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler.Handler)
}
