package usecase

import (
	"context"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/domain/repository"
)

type CreateRankInput *entity.Rank

type CreateRankOutput struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Public bool   `json:"public"`
}

type CreateRankUsecase struct {
	repo repository.RankRepository
}

func NewCreateRankUsecase(repo repository.RankRepository) *CreateRankUsecase {
	return &CreateRankUsecase{repo}
}

func (uc *CreateRankUsecase) Execute(ctx context.Context, input CreateRankInput) (*CreateRankOutput, error) {
	if err := uc.repo.Create(ctx, input); err != nil {
		return nil, err
	}
	return &CreateRankOutput{
		Id:     input.Id,
		Name:   input.Name,
		Public: input.Public,
	}, nil
}

type FindRankInput struct {
	Id string
}

type FindRankOutput struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Public bool   `json:"public"`
}

type FindRankUsecase struct {
	repo repository.RankRepository
}

func NewFindRankUsecase(repo repository.RankRepository) *FindRankUsecase {
	return &FindRankUsecase{repo}
}

func (uc *FindRankUsecase) Execute(ctx context.Context, input FindRankInput) (*FindRankOutput, error) {
	rank, err := uc.repo.FindById(ctx, input.Id)
	if err != nil {
		return nil, err
	}
	if rank == nil {
		return nil, &ResourceNotFoundError{name: "rank", id: input.Id}
	}
	return &FindRankOutput{
		Id:     rank.Id,
		Name:   rank.Name,
		Public: rank.Public,
	}, nil
}

type UpdateRankInput *entity.Rank

type UpdateRankOutput struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Public bool   `json:"public"`
}

type UpdateRankUsecase struct {
	repo repository.RankRepository
}

func NewUpdateRankUsecase(repo repository.RankRepository) *UpdateRankUsecase {
	return &UpdateRankUsecase{repo}
}

func (uc *UpdateRankUsecase) Execute(ctx context.Context, input UpdateRankInput) (*UpdateRankOutput, error) {
	rank, err := uc.repo.FindById(ctx, input.Id)
	if err != nil {
		return nil, err
	}
	if rank == nil {
		return nil, &ResourceNotFoundError{name: "rank", id: input.Id}
	}
	if err := uc.repo.Update(ctx, input); err != nil {
		return nil, err
	}
	return &UpdateRankOutput{
		Id:     input.Id,
		Name:   input.Name,
		Public: input.Public,
	}, nil
}

type DeleteRankInput struct {
	Id string
}

type DeleteRankOutput struct{}

type DeleteRankUsecase struct {
	repo repository.RankRepository
}

func NewDeleteRankUsecase(repo repository.RankRepository) *DeleteRankUsecase {
	return &DeleteRankUsecase{repo}
}

func (uc *DeleteRankUsecase) Execute(ctx context.Context, input DeleteRankInput) (*DeleteRankOutput, error) {
	rank, err := uc.repo.FindById(ctx, input.Id)
	if err != nil {
		return nil, err
	}
	if rank == nil {
		return nil, &ResourceNotFoundError{name: "rank", id: input.Id}
	}
	if err := uc.repo.Delete(ctx, rank); err != nil {
		return nil, err
	}
	return &DeleteRankOutput{}, nil
}
