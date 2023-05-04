package handler

import (
	"github.com/aws/aws-lambda-go/events"
)

// IngestHandler handles requests to the API V2 /manifest endpoints.
func IngestHandler(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {

	result := events.APIGatewayV2HTTPResponse{}

	return &result, nil
}
