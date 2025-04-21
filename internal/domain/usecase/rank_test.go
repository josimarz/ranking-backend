package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/josimarz/ranking-backend/internal/infra/db/inmemory"
	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestCreateRankUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.RankInMemoryRepository{}
	uc := NewCreateRankUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		input := mock.Rank
		want := &CreateRankOutput{
			Id:     mock.Rank.Id,
			Name:   mock.Rank.Name,
			Public: mock.Rank.Public,
		}
		if got, err := uc.Execute(ctx, &input); err != nil || *got != *want {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, want, nil)
		}
	})
}

func TestFindRankUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.RankInMemoryRepository{}
	uc := NewFindRankUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		input := FindRankInput{
			Id: mock.Rank.Id,
		}
		want := &FindRankOutput{
			Id:     mock.Rank.Id,
			Name:   mock.Rank.Name,
			Public: mock.Rank.Public,
		}
		if got, err := uc.Execute(ctx, input); err != nil || *got != *want {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, want, nil)
		}
		input.Id = "4f1c876e-d35f-43d0-b411-470b513d07a0"
		notFoundErr := &ResourceNotFoundError{name: "rank", id: input.Id}
		if got, err := uc.Execute(ctx, input); got != nil || !errors.As(err, &notFoundErr) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, nil, notFoundErr)
		}
	})
}

func TestUpdateRankUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.RankInMemoryRepository{}
	uc := NewUpdateRankUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		rank := mock.Rank
		rank.Name = "Video Games"
		rank.Public = false
		want := &UpdateRankOutput{
			Id:     rank.Id,
			Name:   rank.Name,
			Public: rank.Public,
		}
		if got, err := uc.Execute(ctx, &rank); err != nil || *got != *want {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, rank, got, err, want, nil)
		}
		rank.Id = "40e9ede4-9443-45c9-a2c4-35c6a02f6c78"
		notFoundErr := &ResourceNotFoundError{name: "rank", id: rank.Id}
		if got, err := uc.Execute(ctx, &rank); got != nil || !errors.As(err, &notFoundErr) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, rank, got, err, nil, notFoundErr)
		}
	})
}

func TestDeleteRankUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.RankInMemoryRepository{}
	uc := NewDeleteRankUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		input := DeleteRankInput{
			Id: mock.Rank.Id,
		}
		want := &DeleteRankOutput{}
		if got, err := uc.Execute(ctx, input); err != nil || *got != *want {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, want, nil)
		}
		input.Id = "df0298bd-53bc-474e-87c2-2c260e867672"
		notFoundErr := &ResourceNotFoundError{name: "rank", id: input.Id}
		if got, err := uc.Execute(ctx, input); got != nil || !errors.As(err, &notFoundErr) {
			t.Errorf("Exeucte(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, nil, notFoundErr)
		}
	})
}
