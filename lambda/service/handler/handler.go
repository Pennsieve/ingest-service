package handler

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/pennsieve/pennsieve-go-core/pkg/authorizer"
	"github.com/pennsieve/pennsieve-go-core/pkg/models/permissions"
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
)

// init runs on cold start of lambda and fetches variables and created neo4j driver.
func init() {

	log.SetFormatter(&log.JSONFormatter{})
	ll, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(ll)
	}
}

// IngestHandler handles requests to the API V2 /manifest endpoints.
func IngestHandler(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	var err error
	var apiResponse *events.APIGatewayV2HTTPResponse

	r := regexp.MustCompile(`(?P<method>) (?P<pathKey>.*)`)
	routeKeyParts := r.FindStringSubmatch(request.RouteKey)
	routeKey := routeKeyParts[r.SubexpIndex("pathKey")]

	claims := authorizer.ParseClaims(request.RequestContext.Authorizer.Lambda)
	authorized := false

	switch routeKey {
	case "/ingest/csv":
		switch request.RequestContext.HTTP.Method {
		case "POST":
			//	Return all models for a specific dataset
			if authorized = authorizer.HasRole(*claims, permissions.CreateDeleteRecord); authorized {
				log.Info("hello World Again")
			}
		}
	}

	// Return unauthorized if
	if !authorized {
		apiResponse := events.APIGatewayV2HTTPResponse{
			StatusCode: 403,
			Body:       `{"message": "User is not authorized to perform this action on the dataset."}`,
		}
		return &apiResponse, nil
	}

	// Response
	if err != nil {
		log.Fatalln("Something is wrong with creating the response", err)
	}

	return apiResponse, nil
}
