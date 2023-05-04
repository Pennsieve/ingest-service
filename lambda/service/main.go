package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pennsieve/ingest-service/service/handler"
)

func main() {
	lambda.Start(handler.IngestHandler)
}
