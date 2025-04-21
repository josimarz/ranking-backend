package ddb

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamodbClient(ctx context.Context) (*dynamodb.Client, error) {
	if isRunningOnLambda() {
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return nil, err
		}
		return dynamodb.NewFromConfig(cfg), nil
	}
	cfg, err := config.LoadDefaultConfig(ctx, func(lo *config.LoadOptions) error {
		lo.Region = "us-east-1"
		lo.Credentials = credentials.NewStaticCredentialsProvider("xyz", "123", "")
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://localhost:4566")
	}), nil
}

func isRunningOnLambda() bool {
	_, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	return exists
}
