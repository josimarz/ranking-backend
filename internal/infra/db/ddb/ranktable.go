package ddb

import (
	"context"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

type RankTableDynamodbRepository struct {
	client *dynamodb.Client
}

func NewRankTableDynamodbRepository(client *dynamodb.Client) *RankTableDynamodbRepository {
	return &RankTableDynamodbRepository{client}
}

func (r *RankTableDynamodbRepository) FindById(ctx context.Context, id string) (*entity.RankTable, error) {
	keyEx := expression.Key("rankid").Equal(expression.Value(id))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, err
	}
	input := &dynamodb.QueryInput{
		TableName:                 tableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		IndexName:                 aws.String("gsi"),
	}
	output, err := r.client.Query(ctx, input)
	if err != nil {
		return nil, err
	}
	if output.Count == 0 {
		return nil, nil
	}
	var rankTable entity.RankTable
	for _, item := range output.Items {
		var typ string
		err := attributevalue.Unmarshal(item["typ"], &typ)
		if err != nil {
			return nil, err
		}
		switch typ {
		case "rank":
			var rec rankRecord
			if err := attributevalue.UnmarshalMap(item, &rec); err != nil {
				return nil, err
			}
			rankTable.Id = rec.Id
			rankTable.Name = rec.Name
			rankTable.Public = rec.Public
		case "attribute":
			var rec attributeRecord
			if err := attributevalue.UnmarshalMap(item, &rec); err != nil {
				return nil, err
			}
			attr := entity.Attribute{
				Id:     strings.Split(rec.Id, "/")[1],
				Name:   rec.Name,
				Desc:   rec.Desc,
				Order:  rec.Order,
				RankId: rec.RankId,
			}
			rankTable.Attrs = append(rankTable.Attrs, attr)
		case "entry":
			var rec entryRecord
			if err := attributevalue.UnmarshalMap(item, &rec); err != nil {
				return nil, err
			}
			entry := entity.Entry{
				Id:       strings.Split(rec.Id, "/")[1],
				Name:     rec.Name,
				ImageURL: rec.ImageURL,
				Scores:   rec.Scores,
				RankId:   rec.RankId,
			}
			rankTable.Entries = append(rankTable.Entries, entry)
		}
	}
	sort.Slice(rankTable.Attrs, func(i, j int) bool {
		return rankTable.Attrs[i].Order < rankTable.Attrs[j].Order
	})
	sort.Slice(rankTable.Entries, func(i, j int) bool {
		return rankTable.Entries[i].Name < rankTable.Entries[j].Name
	})
	return &rankTable, nil
}
