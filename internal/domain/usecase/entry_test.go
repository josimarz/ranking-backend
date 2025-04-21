package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/infra/db/inmemory"
	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestCreateEntryUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.EntryInMemoryRepository{}
	uc := NewCreateEntryUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		want := &CreateEntryOutput{
			Id:       mock.Entries[0].Id,
			Name:     mock.Entries[0].Name,
			ImageURL: mock.Entries[0].ImageURL,
			Scores:   mock.Entries[0].Scores,
			RankId:   mock.Entries[0].RankId,
		}
		if got, err := uc.Execute(ctx, &mock.Entries[0]); err != nil || !reflect.DeepEqual(*got, *want) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, mock.Entries[0], got, err, want, nil)
		}
	})
}

func TestFindEntryUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.EntryInMemoryRepository{}
	uc := NewFindEntryUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		entry := mock.Entries[0]
		input := FindEntryInput{
			RankId: entry.RankId,
			Id:     entry.Id,
		}
		want := &FindEntryOutput{
			Id:       entry.Id,
			Name:     entry.Name,
			ImageURL: entry.ImageURL,
			Score:    entry.Scores,
			RankId:   entry.RankId,
		}
		if got, err := uc.Execute(ctx, input); err != nil || !reflect.DeepEqual(*got, *want) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, want, nil)
		}
		input.RankId = "bfa9c9c5-32ac-424a-bff8-736d4a40fb59"
		input.Id = "d5fcf68d-6db0-4167-a426-12d5fe3f0ee2"
		notFoundErr := &ResourceNotFoundError{name: "entry", id: input.Id}
		if got, err := uc.Execute(ctx, input); got != nil || !errors.As(err, &notFoundErr) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, nil, notFoundErr)
		}
	})
}

func TestUpdateEntryUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.EntryInMemoryRepository{}
	uc := NewUpdateEntryUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		entry := mock.Entries[0]
		entry.Name = "Sega Dreamcast"
		entry.ImageURL = "https://videogame.com/dreamcast.png"
		entry.Scores = entity.Scores{
			"Controls": 94,
			"Graphics": 96,
			"Sound":    98,
		}
		want := &UpdateEntryOutput{
			Id:       entry.Id,
			Name:     entry.Name,
			ImageURL: entry.ImageURL,
			Scores:   entry.Scores,
			RankId:   entry.RankId,
		}
		if got, err := uc.Execute(ctx, &entry); err != nil || !reflect.DeepEqual(*got, *want) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, entry, got, err, want, nil)
		}
		entry.Id = "cdd8c04f-cbde-4400-8448-7eaf7440a030"
		entry.RankId = "bc288ada-cfb4-425b-80e3-bb0b5b3ff4f5"
		notFoundErr := &ResourceNotFoundError{name: "entry", id: entry.Id}
		if got, err := uc.Execute(ctx, &entry); got != nil || !errors.As(err, &notFoundErr) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, entry, got, err, nil, notFoundErr)
		}
	})
}

func TestDeleteEntryUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.EntryInMemoryRepository{}
	uc := NewDeleteEntryUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		entry := mock.Entries[0]
		input := DeleteEntryInput{
			RankId: entry.RankId,
			Id:     entry.Id,
		}
		want := &DeleteEntryOutput{}
		if got, err := uc.Execute(ctx, input); err != nil || *got != *want {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, want, nil)
		}
		input.Id = "8724cac5-bf06-4ee1-aa48-2ca071c4da22"
		input.RankId = "379f5dc2-2371-4412-adc9-cdc74013f849"
		notFoundErr := &ResourceNotFoundError{name: "entry", id: input.Id}
		if got, err := uc.Execute(ctx, input); got != nil || !errors.As(err, &notFoundErr) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, nil, notFoundErr)
		}
	})
}
