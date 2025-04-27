package ddb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/josimarz/ranking-backend/internal/infra"
)

func NewDynamodbClient(ctx context.Context) (*dynamodb.Client, error) {
	if infra.IsRunningOnLambda() {
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
		o.BaseEndpoint = aws.String(infra.EndpointURL())
	}), nil
}
