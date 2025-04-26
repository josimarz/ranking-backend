package ddb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
)

var (
	client *dynamodb.Client
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	ctr, err := startContainer(ctx)
	if err != nil {
		log.Fatalf("failed to run locastack container: %v", err)
	}
	if err := connect(ctx, ctr); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := createTable(ctx); err != nil {
		log.Fatalf("failed to create table on dynamodb: %v", err)
	}
	if err := waitForTable(ctx); err != nil {
		log.Fatalf("failed to wait for table exists: %v", err)
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
	cfg, err := config.LoadDefaultConfig(ctx, func(lo *config.LoadOptions) error {
		lo.Region = "us-east-1"
		lo.Credentials = credentials.NewStaticCredentialsProvider("xyz", "123", "")
		return nil
	})
	if err != nil {
		return err
	}
	port, err := ctr.MappedPort(ctx, "4566/tcp")
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("http://localhost:%s", port)
	client = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
	return nil
}

func createTable(ctx context.Context) error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("id"),
			AttributeType: types.ScalarAttributeTypeS,
		}, {
			AttributeName: aws.String("rankid"),
			AttributeType: types.ScalarAttributeTypeS,
		}, {
			AttributeName: aws.String("typ"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("id"),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String("typ"),
			KeyType:       types.KeyTypeRange,
		}},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{{
			IndexName: aws.String("gsi"),
			KeySchema: []types.KeySchemaElement{{
				AttributeName: aws.String("rankid"),
				KeyType:       types.KeyTypeHash,
			}, {
				AttributeName: aws.String("typ"),
				KeyType:       types.KeyTypeRange,
			}},
			Projection: &types.Projection{
				ProjectionType: types.ProjectionTypeAll,
			},
		}},
		TableName:   tableName,
		BillingMode: types.BillingModePayPerRequest,
	}
	if _, err := client.CreateTable(ctx, input); err != nil {
		return err
	}
	return nil
}

func waitForTable(ctx context.Context) error {
	waiter := dynamodb.NewTableExistsWaiter(client)
	input := &dynamodb.DescribeTableInput{
		TableName: tableName,
	}
	return waiter.Wait(ctx, input, 5*time.Second)
}

func putItem[T any](ctx context.Context, rec *T) error {
	item, err := attributevalue.MarshalMap(rec)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: tableName,
		Item:      item,
	}
	if _, err := client.PutItem(ctx, input); err != nil {
		return err
	}
	return nil
}

func getItem[T any](ctx context.Context, id string) (*T, error) {
	var rec T
	var typ string
	switch any(rec).(type) {
	case rankRecord:
		typ = "rank"
	case attributeRecord:
		typ = "attribute"
	case entryRecord:
		typ = "entry"
	default:
		return nil, errors.New("unknown record type")
	}
	key, err := attributevalue.MarshalMap(map[string]string{"id": id, "typ": typ})
	if err != nil {
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: tableName,
		Key:       key,
	}
	res, err := client.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}
	if res.Item == nil {
		return nil, nil
	}
	if err := attributevalue.UnmarshalMap(res.Item, &rec); err != nil {
		return nil, err
	}
	return &rec, nil
}
