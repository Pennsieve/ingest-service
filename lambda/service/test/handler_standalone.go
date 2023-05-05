package handler2

import (
	"fmt"
	"os"

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

func main() {

	log.Info("hello World Again")

	DownloadS3CSVFile("pennsieve-prod-discover-publish-use1", "2/2/metadata/records/Disease.csv")
	myfile := DownloadS3CSVFile("pennsieve-prod-discover-publish-use1", "2/2/metadata/records/Disease.csv")
	println("Filename:", myfile.Name())

}
