package inmemory

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestRankTableInMemoryRepository(t *testing.T) {
	ctx := context.Background()
	r := &RankTableInMemoryRepository{}
	mockRankTable()
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
		id := "caa458f6-e666-4631-981a-97c8a264c2ee"
		if got, err := r.FindById(ctx, id); err != nil || got != nil {
			t.Errorf("FindById(%v, %v) got (%v, %v), want (%v, %v)", ctx, id, got, err, nil, nil)
		}
	})
}

func mockRankTable() {
	mockRank()
	mockAttributes()
	mockEntries()
}

func mockRank() {
	ranks[mock.Rank.Id] = &mock.Rank
}

func mockAttributes() {
	for _, attr := range mock.Attrs {
		key := fmt.Sprintf("%s/%s", attr.RankId, attr.Id)
		attrs[key] = &attr
	}
}

func mockEntries() {
	for _, entry := range mock.Entries {
		key := fmt.Sprintf("%s/%s", entry.RankId, entry.Id)
		entries[key] = &entry
	}
}
