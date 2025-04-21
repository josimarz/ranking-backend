package usecase

import (
	"context"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/repository"
)

type CreateEntryInput *entity.Entry

type CreateEntryOutput struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	ImageURL string        `json:"image_url"`
	Scores   entity.Scores `json:"scores"`
	RankId   string        `json:"rank_id"`
}

type CreateEntryUsecase struct {
	repo repository.EntryRepository
}

func NewCreateEntryUsecase(repo repository.EntryRepository) *CreateEntryUsecase {
	return &CreateEntryUsecase{repo}
}

func (uc *CreateEntryUsecase) Execute(ctx context.Context, input CreateEntryInput) (*CreateEntryOutput, error) {
	if err := uc.repo.Create(ctx, input); err != nil {
		return nil, err
	}
	return &CreateEntryOutput{
		Id:       input.Id,
		Name:     input.Name,
		ImageURL: input.ImageURL,
		Scores:   input.Scores,
		RankId:   input.RankId,
	}, nil
}

type FindEntryInput struct {
	RankId string
	Id     string
}

type FindEntryOutput struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	ImageURL string        `json:"image_url"`
	Score    entity.Scores `json:"scores"`
	RankId   string        `json:"rank_id"`
}

type FindEntryUsecase struct {
	repo repository.EntryRepository
}

func NewFindEntryUsecase(repo repository.EntryRepository) *FindEntryUsecase {
	return &FindEntryUsecase{repo}
}

func (uc *FindEntryUsecase) Execute(ctx context.Context, input FindEntryInput) (*FindEntryOutput, error) {
	entry, err := uc.repo.FindById(ctx, input.RankId, input.Id)
	if err != nil {
		return nil, err
	}
	if entry == nil {
		return nil, &ResourceNotFoundError{name: "entry", id: input.Id}
	}
	return &FindEntryOutput{
		Id:       entry.Id,
		Name:     entry.Name,
		ImageURL: entry.ImageURL,
		Score:    entry.Scores,
		RankId:   entry.RankId,
	}, nil
}

type UpdateEntryInput *entity.Entry

type UpdateEntryOutput struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	ImageURL string        `json:"image_url"`
	Scores   entity.Scores `json:"scores"`
	RankId   string        `json:"rank_id"`
}

type UpdateEntryUsecase struct {
	repo repository.EntryRepository
}

func NewUpdateEntryUsecase(repo repository.EntryRepository) *UpdateEntryUsecase {
	return &UpdateEntryUsecase{repo}
}

func (uc *UpdateEntryUsecase) Execute(ctx context.Context, input UpdateEntryInput) (*UpdateEntryOutput, error) {
	entry, err := uc.repo.FindById(ctx, input.RankId, input.Id)
	if err != nil {
		return nil, err
	}
	if entry == nil {
		return nil, &ResourceNotFoundError{name: "entry", id: input.Id}
	}
	if err := uc.repo.Update(ctx, input); err != nil {
		return nil, err
	}
	return &UpdateEntryOutput{
		Id:       input.Id,
		Name:     input.Name,
		ImageURL: input.ImageURL,
		Scores:   input.Scores,
		RankId:   input.RankId,
	}, nil
}

type DeleteEntryInput struct {
	RankId string
	Id     string
}

type DeleteEntryOutput struct{}

type DeleteEntryUsecase struct {
	repo repository.EntryRepository
}

func NewDeleteEntryUsecase(repo repository.EntryRepository) *DeleteEntryUsecase {
	return &DeleteEntryUsecase{repo}
}

func (uc *DeleteEntryUsecase) Execute(ctx context.Context, input DeleteEntryInput) (*DeleteEntryOutput, error) {
	entry, err := uc.repo.FindById(ctx, input.RankId, input.Id)
	if err != nil {
		return nil, err
	}
	if entry == nil {
		return nil, &ResourceNotFoundError{name: "entry", id: input.Id}
	}
	if err := uc.repo.Delete(ctx, entry); err != nil {
		return nil, err
	}
	return &DeleteEntryOutput{}, nil
}
