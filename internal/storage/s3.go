package storage

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3StorageClient struct {
	Client     *s3.Client
	BucketName string
	Region     string
}

func NewS3StorageClient(accessKey string, secretKey string, region string, bucketName string) (FileStorageClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &S3StorageClient{
		Client:     client,
		BucketName: bucketName,
		Region:     region,
	}, nil
}

func (s *S3StorageClient) UploadFile(file multipart.File, filename string) (string, error) {
	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	fileUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.BucketName, s.Region, filename)

	return fileUrl, nil
}
