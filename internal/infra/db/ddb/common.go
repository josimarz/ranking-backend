package ddb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

var (
	tableName = aws.String("rank")
)

type record struct {
	RecordType string `dynamodbav:"typ"`
}
