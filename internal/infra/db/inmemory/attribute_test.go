package inmemory

import (
	"context"
	"fmt"
	"testing"

	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestAttributeInMemoryRepository(t *testing.T) {
	ctx := context.Background()
	r := &AttributeInMemoryRepository{}
	attr := mock.Attrs[0]
	key := fmt.Sprintf("%s/%s", attr.RankId, attr.Id)
	t.Run("Create", func(t *testing.T) {
		if err := r.Create(ctx, &attr); err != nil {
			t.Errorf("Create(%v, %v) got %v, want %v", ctx, attr, err, nil)
		}
		item, ok := attrs[key]
		if !ok {
			t.Fatal("item was not saved")
		}
		if *item != attr {
			t.Errorf("saved item does not match the expected one: got %v, want %v", item, attr)
		}
	})
	t.Run("FindById", func(t *testing.T) {
		if got, err := r.FindById(ctx, attr.RankId, attr.Id); err != nil || *got != attr {
			t.Errorf("FindById(%v, %v, %v) got (%v, %v), want (%v, %v)", ctx, attr.RankId, attr.Id, got, err, attr, nil)
		}
		rankId := "022ba2ba-524a-4dce-82eb-3fd4e307687d"
		id := "a545e98d-b852-4266-81fe-a0bd286e5cd6"
		if got, err := r.FindById(ctx, rankId, id); err != nil || got != nil {
			t.Errorf("FindById(%v, %v, %v) got (%v, %v), want (%v, %v)", ctx, rankId, id, got, err, nil, nil)
		}
	})
	t.Run("Update", func(t *testing.T) {
		attr.Name = "Design"
		attr.Desc = "Evaluate the design of the console"
		attr.Order = 4
		if err := r.Update(ctx, &attr); err != nil {
			t.Errorf("Update(%v, %v) got %v, want %v", ctx, attr, err, nil)
		}
		item, ok := attrs[key]
		if !ok {
			t.Fatal("item was not saved")
		}
		if *item != attr {
			t.Errorf("saved item does not match the expected one: got %v, want %v", item, attr)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		if err := r.Delete(ctx, &attr); err != nil {
			t.Errorf("Delete(%v, %v) got %v, want %v", ctx, attr, err, nil)
		}
		if _, ok := attrs[key]; ok {
			t.Fatal("item was not deleted from database")
		}
	})
}
