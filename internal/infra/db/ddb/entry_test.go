package ddb

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestEntryDynamodbRepository(t *testing.T) {
	ctx := context.Background()
	r := NewEntryDynamodbRepository(client)
	entry := mock.Entries[0]
	id := fmt.Sprintf("%s/%s", entry.RankId, entry.Id)
	t.Run("Create", func(t *testing.T) {
		if err := r.Create(ctx, &entry); err != nil {
			t.Errorf("Create(%v, %v) got %v, want %v", ctx, entry, err, nil)
		}
		got, err := getItem[entryRecord](ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		want := &entryRecord{
			record: record{
				RecordType: "entry",
			},
			Id:       id,
			Name:     entry.Name,
			ImageURL: entry.ImageURL,
			Scores:   entry.Scores,
			RankId:   entry.RankId,
		}
		if !reflect.DeepEqual(*got, *want) {
			t.Errorf("saved item does not match the expected one: got %v, want %v", got, want)
		}
	})
	t.Run("FindById", func(t *testing.T) {
		if got, err := r.FindById(ctx, entry.RankId, entry.Id); err != nil || !reflect.DeepEqual(*got, entry) {
			t.Errorf("FindById(%v, %v, %v) got (%v, %v), want (%v, %v)", ctx, entry.RankId, entry.Id, got, err, entry, nil)
		}
		rankId := "afb055bc-b110-4efc-8c5b-bd18097603c9"
		id := "3f87c939-eea6-4a8f-946f-aff81983a306"
		if got, err := r.FindById(ctx, rankId, id); got != nil || err != nil {
			t.Errorf("FindById(%v, %v) got (%v, %v), want (%v, %v)", ctx, id, got, err, nil, nil)
		}
	})
	t.Run("Update", func(t *testing.T) {
		entry.Name = "Sega Dreamcast"
		entry.ImageURL = "https://videogame.com/dreamcast.png"
		entry.Scores = entity.Scores{
			"Controls": 90,
			"Graphics": 94,
			"Sound":    98,
		}
		if err := r.Update(ctx, &entry); err != nil {
			t.Errorf("Update(%v, %v) got %v, want %v", ctx, entry, err, nil)
		}
		got, err := getItem[entryRecord](ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		want := &entryRecord{
			record: record{
				RecordType: "entry",
			},
			Id:       id,
			Name:     entry.Name,
			ImageURL: entry.ImageURL,
			Scores:   entry.Scores,
			RankId:   entry.RankId,
		}
		if !reflect.DeepEqual(*got, *want) {
			t.Errorf("saved item does not match the expected one: got %v, want %v", got, want)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		if err := r.Delete(ctx, &entry); err != nil {
			t.Errorf("Delete(%v, %v) got %v, want %v", ctx, entry, err, nil)
		}
		got, err := getItem[entryRecord](ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		if got != nil {
			t.Errorf("item was not deleted from database")
		}
	})
}
