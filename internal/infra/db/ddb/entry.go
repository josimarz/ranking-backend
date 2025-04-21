package ddb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

type entryRecord struct {
	record
	Id       string        `dynamodbav:"id"`
	Name     string        `dynamodbav:"name"`
	ImageURL string        `dynamodbav:"imageurl"`
	Scores   entity.Scores `dynamodbav:"scores"`
	RankId   string        `dynamodbav:"rankid"`
}

type EntryDynamodbRepository struct {
	client *dynamodb.Client
}

func NewEntryDynamodbRepository(client *dynamodb.Client) *EntryDynamodbRepository {
	return &EntryDynamodbRepository{client}
}

func (r *EntryDynamodbRepository) Create(ctx context.Context, entry *entity.Entry) error {
	return r.putItem(ctx, entry)
}

func (r *EntryDynamodbRepository) FindById(ctx context.Context, rankId, id string) (*entity.Entry, error) {
	key, err := attributevalue.MarshalMap(map[string]string{
		"id":  fmt.Sprintf("%s/%s", rankId, id),
		"typ": "entry",
	})
	if err != nil {
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: tableName,
		Key:       key,
	}
	res, err := r.client.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}
	if res.Item == nil {
		return nil, nil
	}
	var rec entryRecord
	if err := attributevalue.UnmarshalMap(res.Item, &rec); err != nil {
		return nil, err
	}
	return &entity.Entry{
		Id:       id,
		Name:     rec.Name,
		ImageURL: rec.ImageURL,
		Scores:   rec.Scores,
		RankId:   rankId,
	}, nil
}

func (r *EntryDynamodbRepository) Update(ctx context.Context, entry *entity.Entry) error {
	return r.putItem(ctx, entry)
}

func (r *EntryDynamodbRepository) Delete(ctx context.Context, entry *entity.Entry) error {
	key, err := attributevalue.MarshalMap(map[string]string{
		"id":  fmt.Sprintf("%s/%s", entry.RankId, entry.Id),
		"typ": "entry",
	})
	if err != nil {
		return err
	}
	input := &dynamodb.DeleteItemInput{
		TableName:    tableName,
		Key:          key,
		ReturnValues: types.ReturnValueNone,
	}
	if _, err := r.client.DeleteItem(ctx, input); err != nil {
		return err
	}
	return nil
}

func (r *EntryDynamodbRepository) putItem(ctx context.Context, entry *entity.Entry) error {
	rec := &entryRecord{
		record: record{
			RecordType: "entry",
		},
		Id:       fmt.Sprintf("%s/%s", entry.RankId, entry.Id),
		Name:     entry.Name,
		ImageURL: entry.ImageURL,
		Scores:   entry.Scores,
		RankId:   entry.RankId,
	}
	item, err := attributevalue.MarshalMap(rec)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName:    tableName,
		Item:         item,
		ReturnValues: types.ReturnValueNone,
	}
	if _, err := r.client.PutItem(ctx, input); err != nil {
		return err
	}
	return nil
}
