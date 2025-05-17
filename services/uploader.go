package services

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadToR2(file multipart.File, filename string) (string, error) {
	client := NewR2Client()

	bucket := os.Getenv("R2_BUCKET_NAME")
	publicBaseURL := os.Getenv("R2_PUBLIC_BASE_URL")

	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	}

	_, err := client.PutObject(context.TODO(), input)
	if err != nil {
		log.Printf("❌ Failed to upload to R2: %v\n", err)
		return "", err
	}

	url := fmt.Sprintf("%s/%s", publicBaseURL, filename)

	return url, nil
}
