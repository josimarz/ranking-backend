package ddb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

type attributeRecord struct {
	record
	Id     string `dynamodbav:"id"`
	Name   string `dynamodbav:"name"`
	Desc   string `dynamodbav:"desc"`
	Order  int    `dynamodbav:"order"`
	RankId string `dynamodbav:"rankid"`
}

type AttributeDynamodbRepository struct {
	client *dynamodb.Client
}

func NewAttributeDynamodbRepository(client *dynamodb.Client) *AttributeDynamodbRepository {
	return &AttributeDynamodbRepository{client}
}

func (r *AttributeDynamodbRepository) Create(ctx context.Context, attr *entity.Attribute) error {
	return r.putItem(ctx, attr)
}

func (r *AttributeDynamodbRepository) FindById(ctx context.Context, rankId, id string) (*entity.Attribute, error) {
	key, err := attributevalue.MarshalMap(map[string]string{
		"id":  fmt.Sprintf("%s/%s", rankId, id),
		"typ": "attribute",
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
	var rec attributeRecord
	if err := attributevalue.UnmarshalMap(res.Item, &rec); err != nil {
		return nil, err
	}
	return &entity.Attribute{
		Id:     id,
		Name:   rec.Name,
		Desc:   rec.Desc,
		Order:  rec.Order,
		RankId: rankId,
	}, nil
}

func (r *AttributeDynamodbRepository) Update(ctx context.Context, attr *entity.Attribute) error {
	return r.putItem(ctx, attr)
}

func (r *AttributeDynamodbRepository) Delete(ctx context.Context, attr *entity.Attribute) error {
	key, err := attributevalue.MarshalMap(map[string]string{
		"id":  fmt.Sprintf("%s/%s", attr.RankId, attr.Id),
		"typ": "attribute",
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

func (r *AttributeDynamodbRepository) putItem(ctx context.Context, attr *entity.Attribute) error {
	rec := &attributeRecord{
		record: record{
			RecordType: "attribute",
		},
		Id:     fmt.Sprintf("%s/%s", attr.RankId, attr.Id),
		Name:   attr.Name,
		Desc:   attr.Desc,
		Order:  attr.Order,
		RankId: attr.RankId,
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
