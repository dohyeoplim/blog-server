package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadToR2(file multipart.File, filename string) (string, error) {
	client := NewR2Client()

	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key:         aws.String(filename),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.r2.cloudflarestorage.com/%s/%s",
		os.Getenv("R2_ACCOUNT_ID"),
		os.Getenv("R2_BUCKET_NAME"),
		filename,
	)

	return url, nil
}
