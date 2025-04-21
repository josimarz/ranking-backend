package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/josimarz/ranking-backend/internal/infra/db/inmemory"
	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestFindRankTableUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.RankTableInMemoryRepository{}
	uc := NewFindRankTableUsecase(repo)
	mockRankTable(ctx)
	t.Run("Execute", func(t *testing.T) {
		want := &FindRankTableOutput{
			Id:     mock.Rank.Id,
			Name:   mock.Rank.Name,
			Public: mock.Rank.Public,
		}
		for _, attr := range mock.Attrs {
			want.Attrs = append(want.Attrs, attributeOutput{
				Id:    attr.Id,
				Name:  attr.Name,
				Desc:  attr.Desc,
				Order: attr.Order,
			})
		}
		for _, entry := range mock.Entries {
			want.Entries = append(want.Entries, entryOutput{
				Id:       entry.Id,
				Name:     entry.Name,
				ImageURL: entry.ImageURL,
				Scores:   entry.Scores,
			})
		}
		input := FindRankTableInput{
			Id: mock.Rank.Id,
		}
		if got, err := uc.Execute(ctx, input); err != nil || !reflect.DeepEqual(*got, *want) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, want, nil)
		}
		input.Id = "b981a90e-181e-40f1-8b10-8b122e9c35af"
		notFoundErr := &ResourceNotFoundError{name: "rank", id: input.Id}
		if got, err := uc.Execute(ctx, input); got != nil || !errors.As(err, &notFoundErr) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, nil, notFoundErr)
		}
	})
}

func mockRankTable(ctx context.Context) {
	inmemory.ClearDatabase()
	mockRank(ctx)
	mockAttributes(ctx)
	mockEntries(ctx)
}

func mockRank(ctx context.Context) {
	repo := &inmemory.RankInMemoryRepository{}
	repo.Create(ctx, &mock.Rank)
}

func mockAttributes(ctx context.Context) {
	repo := &inmemory.AttributeInMemoryRepository{}
	for _, attr := range mock.Attrs {
		repo.Create(ctx, &attr)
	}
}

func mockEntries(ctx context.Context) {
	repo := &inmemory.EntryInMemoryRepository{}
	for _, entry := range mock.Entries {
		repo.Create(ctx, &entry)
	}
}
