package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

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

func FetchLocalCSVFile(file_path string) (file *os.File) {

	file, err := os.Open(file_path)

	if err != nil {
		log.Fatalf("Unable to open file %q, %v", file_path, err)
	}

	return
}

func ReadCSVFile(csv_file *os.File) (lines [][]string) {
	lines, err := csv.NewReader(csv_file).ReadAll()

	if err != nil {
		return nil
	}

	return lines
}

func main() {

	log.Info("hello World Again")

	test_filename := "../handler/data/Disease.csv"

	// myfile := DownloadS3CSVFile("pennsieve-prod-discover-publish-use1", "2/2/metadata/records/Disease.csv")
	myfile := FetchLocalCSVFile(test_filename)
	println("Filename:", myfile.Name())
	// lines := ReadCSVFile(myfile)

	// csv_header := lines[0]
	// csv_rows := lines[1:]

	file_base := filepath.Base(test_filename)
	model_name := strings.Split(file_base, ".")[0]

	println(file_base)
	println(model_name)

	// Model CSV

	// Properties CSV

	// Record CSV

}
