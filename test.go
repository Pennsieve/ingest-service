package main

import (
	"encoding/csv"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// FETCH FILE FROM AWS ########################

	// Uses example from https://stackoverflow.com/a/49270166
	test_bucket := "pennsieve-prod-discover-publish-use1"
	test_item := "2/2/metadata/records/Disease.csv"

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(test_item)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(test_bucket),
			Key:    aws.String(test_item),
		})

	if err != nil {
		log.Fatalf("Unable to download item %q, %v", test_item, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	// READ CSV FILE ############################

	r := csv.NewReader(strings.NewReader(file.Name()))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}

	// PROCESS CSV ##############################

	// DELIVER PROCESSED DATA ###################

	// CLEAN UP LAMBDA ENVIRONMENT ##############

}
