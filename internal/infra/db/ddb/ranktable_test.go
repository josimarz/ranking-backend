package ddb

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestRankTableDynamodbRepository(t *testing.T) {
	ctx := context.Background()
	r := NewRankTableDynamodbRepository(client)
	if err := mockRankTable(ctx); err != nil {
		t.Fatal(err)
	}
	t.Run("FindById", func(t *testing.T) {
		want := &entity.RankTable{
			Id:      mock.Rank.Id,
			Name:    mock.Rank.Name,
			Public:  mock.Rank.Public,
			Attrs:   mock.Attrs,
			Entries: mock.Entries,
		}
		sort.Slice(want.Attrs, func(i, j int) bool {
			return want.Attrs[i].Order < want.Attrs[j].Order
		})
		sort.Slice(want.Entries, func(i, j int) bool {
			return want.Entries[i].Name < want.Entries[j].Name
		})
		if got, err := r.FindById(ctx, mock.Rank.Id); err != nil || !reflect.DeepEqual(got, want) {
			t.Errorf("FindById(%v, %v) got (%v, %v), want (%v, %v)", ctx, mock.Rank.Id, got, err, want, nil)
		}
		id := "1c87a2f3-15ee-473c-9ad6-adb889dbef36"
		if got, err := r.FindById(ctx, id); err != nil || got != nil {
			t.Errorf("FindById(%v, %v) got (%v, %v), want (%v, %v)", ctx, id, got, err, nil, nil)
		}
	})
}

func mockRankTable(ctx context.Context) error {
	if err := mockRank(ctx); err != nil {
		return err
	}
	if err := mockAttributes(ctx); err != nil {
		return err
	}
	return mockEntries(ctx)
}

func mockRank(ctx context.Context) error {
	rec := &rankRecord{
		record: record{
			RecordType: "rank",
		},
		Id:     mock.Rank.Id,
		RankId: mock.Rank.Id,
		Name:   mock.Rank.Name,
		Public: mock.Rank.Public,
	}
	return putItem(ctx, rec)
}

func mockAttributes(ctx context.Context) error {
	for _, attr := range mock.Attrs {
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
		if err := putItem(ctx, rec); err != nil {
			return err
		}
	}
	return nil
}

func mockEntries(ctx context.Context) error {
	for _, entry := range mock.Entries {
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
		if err := putItem(ctx, rec); err != nil {
			return err
		}
	}
	return nil
}
