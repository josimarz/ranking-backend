package inmemory

import (
	"context"
	"fmt"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

var (
	entries map[string]*entity.Entry = make(map[string]*entity.Entry)
)

type EntryInMemoryRepository struct{}

func (r *EntryInMemoryRepository) Create(ctx context.Context, entry *entity.Entry) error {
	key := fmt.Sprintf("%s/%s", entry.RankId, entry.Id)
	entries[key] = entry
	return nil
}

func (r *EntryInMemoryRepository) FindById(ctx context.Context, rankId, id string) (*entity.Entry, error) {
	key := fmt.Sprintf("%s/%s", rankId, id)
	if entry, ok := entries[key]; ok {
		return entry, nil
	}
	return nil, nil
}

func (r *EntryInMemoryRepository) Update(ctx context.Context, entry *entity.Entry) error {
	key := fmt.Sprintf("%s/%s", entry.RankId, entry.Id)
	entries[key] = entry
	return nil
}

func (r *EntryInMemoryRepository) Delete(ctx context.Context, entry *entity.Entry) error {
	key := fmt.Sprintf("%s/%s", entry.RankId, entry.Id)
	delete(entries, key)
	return nil
}
