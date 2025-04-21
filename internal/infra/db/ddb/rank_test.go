package ddb

import (
	"context"
	"testing"

	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestRankDynamodbRepository(t *testing.T) {
	ctx := context.Background()
	r := NewRankDynamodbRepository(client)
	rank := mock.Rank
	t.Run("Create", func(t *testing.T) {
		if err := r.Create(ctx, &rank); err != nil {
			t.Errorf("Create(%v, %v) got %v, want %v", ctx, mock.Rank, err, nil)
		}
		got, err := getItem[rankRecord](ctx, rank.Id)
		if err != nil {
			t.Fatal(err)
		}
		want := &rankRecord{
			record: record{
				RecordType: "rank",
			},
			Id:     rank.Id,
			RankId: rank.Id,
			Name:   rank.Name,
			Public: rank.Public,
		}
		if *got != *want {
			t.Errorf("saved item does not match the expected one: got %v, want %v", got, want)
		}
	})
	t.Run("FindById", func(t *testing.T) {
		if got, err := r.FindById(ctx, rank.Id); err != nil || *got != rank {
			t.Errorf("FindById(%v, %v) got (%v, %v), want (%v, %v)", ctx, rank.Id, got, err, rank, nil)
		}
		id := "2a7f843d-5084-41ae-b6a0-b23ce80d78c4"
		if got, err := r.FindById(ctx, id); got != nil || err != nil {
			t.Errorf("FindById(%v, %v) got (%v, %v), want (%v, %v)", ctx, id, got, err, nil, nil)
		}
	})
	t.Run("Update", func(t *testing.T) {
		rank.Name = "Video Games"
		rank.Public = false
		if err := r.Update(ctx, &rank); err != nil {
			t.Errorf("Update(%v, %v) got %v, want %v", ctx, rank, err, nil)
		}
		got, err := getItem[rankRecord](ctx, rank.Id)
		if err != nil {
			t.Fatal(err)
		}
		want := &rankRecord{
			record: record{
				RecordType: "rank",
			},
			Id:     rank.Id,
			RankId: rank.Id,
			Name:   rank.Name,
			Public: rank.Public,
		}
		if *got != *want {
			t.Errorf("saved item does not match the expected one: got %v, wnat %v", got, want)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		if err := r.Delete(ctx, &rank); err != nil {
			t.Errorf("Delete(%v, %v) got %v, want %v", ctx, rank, err, nil)
		}
		got, err := getItem[rankRecord](ctx, rank.Id)
		if err != nil {
			t.Fatal(err)
		}
		if got != nil {
			t.Errorf("item was not deleted from database")
		}
	})
}
