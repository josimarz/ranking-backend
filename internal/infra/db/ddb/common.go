package ddb

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var (
	tableName = aws.String("rank")
)

type record struct {
	RecordType string `dynamodbav:"typ"`
}

func init() {
	if value, ok := os.LookupEnv("AWS_TABLE"); ok {
		tableName = aws.String(value)
	}
}
