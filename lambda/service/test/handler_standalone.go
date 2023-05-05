package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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

func InferCorrectDtypeString(data_as_string string) (dtype_as_string string) {
	// Try conversions one by one, return result any time there is no error.
	// Fallback condition is leave it as the string
	dtype_as_string = "string"

	// INTEGER
	_, err := strconv.Atoi(data_as_string)
	if err != nil {
		dtype_as_string = "int"
	}

	// FLOAT
	_, err = strconv.ParseFloat(data_as_string, 32) // Just default to 64-bit int
	if err != nil {
		dtype_as_string = "float"
	}

	// DATETIME

	return
}

func InferProperties(csv_header, csv_first_line []string) (properties_as_lines [][]string) {
	properties_as_lines = [][]string{}
	for i := 0; i < len(csv_header); i++ {
		properties_as_lines = append(properties_as_lines, []string{
			csv_header[i],
			InferCorrectDtypeString(csv_first_line[i]),
		})
	}

	return
}

func main() {

	log.Info("hello World Again")

	test_filename := "../handler/data/Disease.csv"
	test_username := "Joe"

	// myfile := DownloadS3CSVFile("pennsieve-prod-discover-publish-use1", "2/2/metadata/records/Disease.csv")
	myfile := FetchLocalCSVFile(test_filename)
	println("Filename:", myfile.Name())
	lines := ReadCSVFile(myfile)

	csv_header := lines[0]
	csv_rows := lines[1:]

	file_base := filepath.Base(test_filename)
	model_name := strings.Split(file_base, ".")[0]

	println("Parsing file:", file_base)
	println("Inferred model name:", model_name)

	// Model CSV
	model_csv_lines := [][]string{
		[]string{"modelName", "createdBy", "createdDate"},
		[]string{model_name, test_username, time.Now().String()},
	}
	println(model_csv_lines)

	// Properties CSV
	properties_csv_lines := [][]string{
		[]string{"propertyName", "dataType"},
	}
	properties_csv_lines = append(properties_csv_lines, InferProperties(csv_header, csv_rows[0])...)

	// Record CSV

}
