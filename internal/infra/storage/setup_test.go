package storage

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
)

var (
	client *s3.Client
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	ctr, err := startContainer(ctx)
	if err != nil {
		log.Fatalf("failed to run localstack container: %v", err)
	}
	if err := connect(ctx, ctr); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := createBucket(ctx); err != nil {
		log.Fatalf("failed to create bucket: %v", err)
	}
	if err := waitForBucket(ctx); err != nil {
		log.Fatalf("failed to wait for bucket exists: %v", err)
	}
	defer func() {
		if err := ctr.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %v", err)
		}
	}()
	os.Exit(m.Run())
}

func startContainer(ctx context.Context) (*localstack.LocalStackContainer, error) {
	return localstack.Run(ctx, "localstack/localstack:latest")
}

func connect(ctx context.Context, ctr *localstack.LocalStackContainer) error {
	port, err := ctr.MappedPort(ctx, "4566/tcp")
	if err != nil {
		return err
	}
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("xyz", "123", "")),
	)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("http://localhost:%s", port.Port())
	client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	})
	return nil
}

func createBucket(ctx context.Context) error {
	input := &s3.CreateBucketInput{
		Bucket: bucketName,
	}
	if _, err := client.CreateBucket(ctx, input); err != nil {
		return err
	}
	return nil
}

func waitForBucket(ctx context.Context) error {
	waiter := s3.NewBucketExistsWaiter(client)
	input := &s3.HeadBucketInput{
		Bucket: bucketName,
	}
	return waiter.Wait(ctx, input, 5*time.Second)
}
