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
	log.Println("📤 Starting R2 upload...")
	client := NewR2Client()

	bucket := os.Getenv("R2_BUCKET_NAME")
	accountID := os.Getenv("R2_ACCOUNT_ID")

	log.Printf("🪣 Bucket: %s\n", bucket)
	log.Printf("👤 Account ID: %s\n", accountID)
	log.Printf("📁 Filename: %s\n", filename)

	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	}

	log.Println("📡 Sending PutObject request to R2...")
	_, err := client.PutObject(context.TODO(), input)
	if err != nil {
		log.Printf("❌ Failed to upload to R2: %v\n", err)
		return "", err
	}

	url := fmt.Sprintf("https://%s.r2.cloudflarestorage.com/%s/%s", accountID, bucket, filename)
	log.Printf("✅ Upload successful. URL: %s\n", url)

	return url, nil
}
