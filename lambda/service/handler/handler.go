package handler

import (
	"fmt"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pennsieve/pennsieve-go-core/pkg/authorizer"
	"github.com/pennsieve/pennsieve-go-core/pkg/models/permissions"
	log "github.com/sirupsen/logrus"
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

func DownloadS3CSVFile(bucket_name, key_name string) (file *os.File) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(key_name)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket_name),
			Key:    aws.String(key_name),
		})

	if err != nil {
		log.Fatalf("Unable to download item %q, %v", key_name, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	return
}

// IngestHandler handles requests to the API V2 /manifest endpoints.
func IngestHandler(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	var err error
	var apiResponse *events.APIGatewayV2HTTPResponse

	log.Info("Running the Ingest Handler")

	r := regexp.MustCompile(`(?P<method>) (?P<pathKey>.*)`)
	routeKeyParts := r.FindStringSubmatch(request.RouteKey)
	routeKey := routeKeyParts[r.SubexpIndex("pathKey")]

	claims := authorizer.ParseClaims(request.RequestContext.Authorizer.Lambda)
	authorized := false

	switch routeKey {
	case "/ingest":
		switch request.RequestContext.HTTP.Method {
		case "POST":
			//	Return all models for a specific dataset
			if authorized = authorizer.HasRole(*claims, permissions.CreateDeleteRecord); authorized {
				log.Info("hello World Again")

				s3file := DownloadS3CSVFile("pennsieve-prod-discover-publish-use1", "2/2/metadata/records/Disease.csv")

				fmt.Println("File name:", s3file.Name())
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
