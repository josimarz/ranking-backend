package inmemory

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestEntryInMemoryRepository(t *testing.T) {
	ctx := context.Background()
	r := &EntryInMemoryRepository{}
	entry := mock.Entries[0]
	key := fmt.Sprintf("%s/%s", entry.RankId, entry.Id)
	t.Run("Create", func(t *testing.T) {
		if err := r.Create(ctx, &entry); err != nil {
			t.Errorf("Create(%v, %v) got %v, want %v", ctx, entry, err, nil)
		}
		item, ok := entries[key]
		if !ok {
			t.Fatal("item was not saved")
		}
		if !reflect.DeepEqual(*item, entry) {
			t.Errorf("saved item does not match the expected one: got %v, want %v", item, entry)
		}
	})
	t.Run("FindById", func(t *testing.T) {
		if got, err := r.FindById(ctx, entry.RankId, entry.Id); err != nil || !reflect.DeepEqual(*got, entry) {
			t.Errorf("FindById(%v, %v, %v) got (%v, %v), want (%v, %v)", ctx, entry.RankId, entry.Id, got, err, entry, nil)
		}
		rankId := "19622a05-4bb2-4f14-8d26-68d1ab1b7c2f"
		id := "a1d00621-e77d-4feb-81c9-53bc886fc05f"
		if got, err := r.FindById(ctx, rankId, id); err != nil || got != nil {
			t.Errorf("FindById(%v, %v, %v) got (%v, %v), want (%v, %v)", ctx, rankId, id, got, err, nil, nil)
		}
	})
	t.Run("Update", func(t *testing.T) {
		entry.Name = "Sega Dreamcast"
		entry.ImageURL = "https://videogame.com/dreamcast.png"
		entry.Scores = entity.Scores{
			"Controls": 92,
			"Graphics": 95,
			"Sound":    97,
		}
		if err := r.Update(ctx, &entry); err != nil {
			t.Errorf("Update(%v, %v) got %v, want %v", ctx, entry, err, nil)
		}
		item, ok := entries[key]
		if !ok {
			t.Fatal("item was not saved")
		}
		if !reflect.DeepEqual(*item, entry) {
			t.Errorf("saved item does not match the expected one: got %v, want %v", item, entry)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		if err := r.Delete(ctx, &entry); err != nil {
			t.Errorf("Delete(%v, %v) got %v, want %v", ctx, entry, err, nil)
		}
		if _, ok := entries[key]; ok {
			t.Fatal("item was not deleted from database")
		}
	})
}
