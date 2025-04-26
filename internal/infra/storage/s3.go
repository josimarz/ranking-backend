package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/josimarz/ranking-backend/internal/infra"
)

var (
	bucketName = aws.String("ranking")
)

func NewS3Client(ctx context.Context) (*s3.Client, error) {
	if infra.IsRunningOnLambda() {
		cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
		if err != nil {
			return nil, err
		}
		return s3.NewFromConfig(cfg), nil
	}
	cfg, err := config.LoadDefaultConfig(ctx, func(lo *config.LoadOptions) error {
		lo.Region = "us-east-1"
		lo.Credentials = credentials.NewStaticCredentialsProvider("xyz", "123", "")
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String("http://localhost:4566")
		o.UsePathStyle = true
	}), nil
}

type FileS3Storage struct {
	client *s3.Client
}

func NewFileS3Storage(client *s3.Client) *FileS3Storage {
	return &FileS3Storage{client}
}

func (s *FileS3Storage) Upload(ctx context.Context, path string, file io.Reader) (string, error) {
	input := &s3.PutObjectInput{
		Bucket: bucketName,
		Key:    aws.String(path),
		Body:   file,
		ACL:    types.ObjectCannedACLPublicRead,
	}
	if _, err := s.client.PutObject(ctx, input); err != nil {
		return "", err
	}
	return s.buildURL(path), nil
}

func (*FileS3Storage) buildURL(path string) string {
	if infra.IsRunningOnLambda() {
		region := os.Getenv("AWS_REGION")
		return fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", *bucketName, region, path)
	}
	return fmt.Sprintf("http://localhost:4566/%s/%s", *bucketName, path)
}
