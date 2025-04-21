package inmemory

import (
	"context"
	"testing"

	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestRankInMemoryRepository(t *testing.T) {
	ctx := context.Background()
	r := &RankInMemoryRepository{}
	rank := mock.Rank
	id := rank.Id
	t.Run("Create", func(t *testing.T) {
		if err := r.Create(ctx, &rank); err != nil {
			t.Errorf("Create(%v, %v) got %v, want %v", ctx, rank, err, nil)
		}
		item, ok := ranks[id]
		if !ok {
			t.Fatal("item was not saved")
		}
		if *item != rank {
			t.Errorf("saved item does not match the expected one: got %v, want %v", item, rank)
		}
	})
	t.Run("FindById", func(t *testing.T) {
		if got, err := r.FindById(ctx, id); err != nil || *got != rank {
			t.Errorf("FindById(%v, %v) got (%v, %v), want (%v, %v)", ctx, id, got, err, rank, nil)
		}
		id := "8a2d5e7a-db3c-40c4-9bb2-8011f5eed6efs"
		if got, err := r.FindById(ctx, id); err != nil || got != nil {
			t.Errorf("FindById(%v, %v) got (%v, %v), want (%v, %v)", ctx, id, got, err, nil, nil)
		}
	})
	t.Run("Update", func(t *testing.T) {
		rank.Name = "Best Soccer Teams of All Time"
		rank.Public = false
		if err := r.Update(ctx, &rank); err != nil {
			t.Errorf("Update(%v, %v) got %v, want %v", ctx, rank, err, nil)
		}
		item, ok := ranks[id]
		if !ok {
			t.Fatal("item was not saved")
		}
		if *item != rank {
			t.Errorf("saved item does not match the expected one: got %v, want %v", item, rank)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		if err := r.Delete(ctx, &rank); err != nil {
			t.Errorf("Delete(%v, %v) got %v, want %v", ctx, rank, err, nil)
		}
		if _, ok := ranks[id]; ok {
			t.Fatal("item was not deleted from database")
		}
	})
}
