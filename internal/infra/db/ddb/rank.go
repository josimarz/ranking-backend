package ddb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

type rankRecord struct {
	record
	Id     string `dynamodbav:"id"`
	RankId string `dynamodbav:"rankid"`
	Name   string `dynamodbav:"name"`
	Public bool   `dynamodbav:"public"`
}

type RankDynamodbRepository struct {
	client *dynamodb.Client
}

func NewRankDynamodbRepository(client *dynamodb.Client) *RankDynamodbRepository {
	return &RankDynamodbRepository{client}
}

func (r *RankDynamodbRepository) Create(ctx context.Context, rank *entity.Rank) error {
	return r.putItem(ctx, rank)
}

func (r *RankDynamodbRepository) FindById(ctx context.Context, id string) (*entity.Rank, error) {
	key, err := attributevalue.MarshalMap(map[string]string{"id": id, "typ": "rank"})
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
	var rec rankRecord
	if err := attributevalue.UnmarshalMap(res.Item, &rec); err != nil {
		return nil, err
	}
	return &entity.Rank{
		Id:     rec.Id,
		Name:   rec.Name,
		Public: rec.Public,
	}, nil
}

func (r *RankDynamodbRepository) Update(ctx context.Context, rank *entity.Rank) error {
	return r.putItem(ctx, rank)
}

func (r *RankDynamodbRepository) Delete(ctx context.Context, rank *entity.Rank) error {
	key, err := attributevalue.MarshalMap(map[string]string{"id": rank.Id, "typ": "rank"})
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

func (r *RankDynamodbRepository) putItem(ctx context.Context, rank *entity.Rank) error {
	rec := &rankRecord{
		record: record{
			RecordType: "rank",
		},
		Id:     rank.Id,
		RankId: rank.Id,
		Name:   rank.Name,
		Public: rank.Public,
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
