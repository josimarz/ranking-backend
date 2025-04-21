package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/josimarz/ranking-backend/internal/infra/db/inmemory"
	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestCreateAttributeUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.AttributeInMemoryRepository{}
	uc := NewCreateAttributeUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		want := &CreateAttributeOutput{
			Id:     mock.Attrs[0].Id,
			Name:   mock.Attrs[0].Name,
			Desc:   mock.Attrs[0].Desc,
			Order:  mock.Attrs[0].Order,
			RankId: mock.Attrs[0].RankId,
		}
		if got, err := uc.Execute(ctx, &mock.Attrs[0]); err != nil || *got != *want {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, mock.Attrs, got, err, want, nil)
		}
	})
}

func TestFindAttributeUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.AttributeInMemoryRepository{}
	uc := NewFindAttributeUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		attr := mock.Attrs[0]
		input := FindAttributeInput{
			RankId: attr.RankId,
			Id:     attr.Id,
		}
		want := &FindAttributeOutput{
			Id:     attr.Id,
			Name:   attr.Name,
			Desc:   attr.Desc,
			Order:  attr.Order,
			RankId: attr.RankId,
		}
		if got, err := uc.Execute(ctx, input); err != nil || *got != *want {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, want, nil)
		}
		input.RankId = "6d61415f-c2e1-40e5-a2bf-8d85e767f04e"
		input.Id = "a7a5af12-c59e-4e7d-be85-4ae6fb9d3622"
		notFoundErr := &ResourceNotFoundError{name: "attribute", id: input.Id}
		if got, err := uc.Execute(ctx, input); got != nil || !errors.As(err, &notFoundErr) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, nil, notFoundErr)
		}
	})
}

func TestUpdateAttributeUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.AttributeInMemoryRepository{}
	uc := NewUpdateAttributeUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		attr := mock.Attrs[0]
		attr.Name = "Design"
		attr.Desc = "Evaluate the video game console design"
		attr.Order = 4
		want := &UpdateAttributeOutput{
			Id:     attr.Id,
			Name:   attr.Name,
			Desc:   attr.Desc,
			Order:  attr.Order,
			RankId: attr.RankId,
		}
		if got, err := uc.Execute(ctx, &attr); err != nil || *got != *want {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, attr, got, err, want, nil)
		}
		attr.Id = "0e87e789-6b24-4818-b52f-c754583fe59f"
		attr.RankId = "799b3bcb-0536-4bcc-97a1-ad1b4142a128"
		notFoundErr := &ResourceNotFoundError{name: "attribute", id: attr.Id}
		if got, err := uc.Execute(ctx, &attr); got != nil || !errors.As(err, &notFoundErr) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, attr, got, err, nil, notFoundErr)
		}
	})
}

func TestDeleteAttributeUsecase(t *testing.T) {
	ctx := context.Background()
	repo := &inmemory.AttributeInMemoryRepository{}
	uc := NewDeleteAttributeUsecase(repo)
	t.Run("Execute", func(t *testing.T) {
		attr := mock.Attrs[0]
		input := DeleteAttributeInput{
			RankId: attr.RankId,
			Id:     attr.Id,
		}
		want := &DeleteAttributeOutput{}
		if got, err := uc.Execute(ctx, input); err != nil || *got != *want {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, want, nil)
		}
		input.RankId = "a0a86620-8912-4124-9f6f-59ead1a40219"
		input.Id = "f0594487-ccc9-432d-86a6-2bf1418c3c66"
		notFoundErr := &ResourceNotFoundError{name: "attribute", id: input.Id}
		if got, err := uc.Execute(ctx, input); got != nil || !errors.As(err, &notFoundErr) {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, nil, notFoundErr)
		}
	})
}
