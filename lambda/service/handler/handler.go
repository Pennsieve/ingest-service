package handler

import (
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

// IngestHandler handles requests to the API V2 /manifest endpoints.
func IngestHandler(request events.APIGatewgo ayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {

	result := events.APIGatewayV2HTTPResponse{}

	log.info("hello")

	return &result, nil
}
