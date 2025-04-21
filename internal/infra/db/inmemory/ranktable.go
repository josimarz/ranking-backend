package inmemory

import (
	"context"
	"sort"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

type RankTableInMemoryRepository struct{}

func (r *RankTableInMemoryRepository) FindById(ctx context.Context, id string) (*entity.RankTable, error) {
	rank, ok := ranks[id]
	if !ok {
		return nil, nil
	}
	rt := &entity.RankTable{
		Id:      rank.Id,
		Name:    rank.Name,
		Public:  rank.Public,
		Attrs:   r.filterAttributes(rank.Id),
		Entries: r.filterEntries(rank.Id),
	}
	sort.Slice(rt.Attrs, func(i, j int) bool {
		return rt.Attrs[i].Order < rt.Attrs[j].Order
	})
	sort.Slice(rt.Entries, func(i, j int) bool {
		return rt.Entries[i].Name < rt.Entries[j].Name
	})
	return rt, nil
}

func (r *RankTableInMemoryRepository) filterAttributes(rankId string) []entity.Attribute {
	var items []entity.Attribute
	for _, item := range attrs {
		if item.RankId == rankId {
			items = append(items, *item)
		}
	}
	return items
}

func (r *RankTableInMemoryRepository) filterEntries(rankId string) []entity.Entry {
	var items []entity.Entry
	for _, item := range entries {
		if item.RankId == rankId {
			items = append(items, *item)
		}
	}
	return items
}
