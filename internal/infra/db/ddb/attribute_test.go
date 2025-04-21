package ddb

import (
	"context"
	"fmt"
	"testing"

	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestAttributeDynamodbRepository(t *testing.T) {
	ctx := context.Background()
	r := NewAttributeDynamodbRepository(client)
	attr := mock.Attrs[0]
	id := fmt.Sprintf("%s/%s", attr.RankId, attr.Id)
	t.Run("Create", func(t *testing.T) {
		if err := r.Create(ctx, &attr); err != nil {
			t.Errorf("Create(%v, %v) got %v, want %v", ctx, attr, err, nil)
		}
		got, err := getItem[attributeRecord](ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		want := &attributeRecord{
			record: record{
				RecordType: "attribute",
			},
			Id:     id,
			Name:   attr.Name,
			Desc:   attr.Desc,
			Order:  attr.Order,
			RankId: attr.RankId,
		}
		if *got != *want {
			t.Errorf("saved item does not match the expected one: got %v, want %v", got, want)
		}
	})
	t.Run("FindById", func(t *testing.T) {
		if got, err := r.FindById(ctx, attr.RankId, attr.Id); err != nil || *got != attr {
			t.Errorf("FindById(%v, %v, %v) got (%v, %v), want (%v, %v)", ctx, attr.RankId, attr.Id, got, err, attr, nil)
		}
		rankId := "5017e29b-5231-4ebe-a25f-74f832508011"
		id := "bb0c5d2c-f17d-429a-8b48-4bb48dffaff1"
		if got, err := r.FindById(ctx, rankId, id); got != nil || err != nil {
			t.Errorf("FindById(%v, %v, %v) got (%v, %v), want (%v, %v)", ctx, rankId, id, got, err, nil, nil)
		}
	})
	t.Run("Update", func(t *testing.T) {
		attr.Name = "Design"
		attr.Desc = "Evaluate the design of the console"
		attr.Order = 3
		if err := r.Update(ctx, &attr); err != nil {
			t.Errorf("Update(%v, %v) got %v, want %v", ctx, attr, err, nil)
		}
		got, err := getItem[attributeRecord](ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		want := &attributeRecord{
			record: record{
				RecordType: "attribute",
			},
			Id:     id,
			Name:   attr.Name,
			Desc:   attr.Desc,
			Order:  attr.Order,
			RankId: attr.RankId,
		}
		if *got != *want {
			t.Errorf("saved item does not match the expected one: got %v, want %v", got, want)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		if err := r.Delete(ctx, &attr); err != nil {
			t.Errorf("Delete(%v, %v) got %v, want %v", ctx, attr, err, nil)
		}
		got, err := getItem[attributeRecord](ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		if got != nil {
			t.Error("item was not delete from database")
		}
	})
}
