package usecase

import (
	"context"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/domain/repository"
)

type CreateAttributeInput *entity.Attribute

type CreateAttributeOutput struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"description"`
	Order  int    `json:"order"`
	RankId string `json:"rank_id"`
}

type CreateAttributeUsecase struct {
	repo repository.AttributeRepository
}

func NewCreateAttributeUsecase(repo repository.AttributeRepository) *CreateAttributeUsecase {
	return &CreateAttributeUsecase{repo}
}

func (uc *CreateAttributeUsecase) Execute(ctx context.Context, input CreateAttributeInput) (*CreateAttributeOutput, error) {
	if err := uc.repo.Create(ctx, input); err != nil {
		return nil, err
	}
	return &CreateAttributeOutput{
		Id:     input.Id,
		Name:   input.Name,
		Desc:   input.Desc,
		Order:  input.Order,
		RankId: input.RankId,
	}, nil
}

type FindAttributeInput struct {
	RankId string
	Id     string
}

type FindAttributeOutput struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"description"`
	Order  int    `json:"order"`
	RankId string `json:"rank_id"`
}

type FindAttributeUsecase struct {
	repo repository.AttributeRepository
}

func NewFindAttributeUsecase(repo repository.AttributeRepository) *FindAttributeUsecase {
	return &FindAttributeUsecase{repo}
}

func (uc *FindAttributeUsecase) Execute(ctx context.Context, input FindAttributeInput) (*FindAttributeOutput, error) {
	attr, err := uc.repo.FindById(ctx, input.RankId, input.Id)
	if err != nil {
		return nil, err
	}
	if attr == nil {
		return nil, &ResourceNotFoundError{name: "attribute", id: input.Id}
	}
	return &FindAttributeOutput{
		Id:     attr.Id,
		Name:   attr.Name,
		Desc:   attr.Desc,
		Order:  attr.Order,
		RankId: attr.RankId,
	}, nil
}

type UpdateAttributeInput *entity.Attribute

type UpdateAttributeOutput struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"description"`
	Order  int    `json:"order"`
	RankId string `json:"rank_id"`
}

type UpdateAttributeUsecase struct {
	repo repository.AttributeRepository
}

func NewUpdateAttributeUsecase(repo repository.AttributeRepository) *UpdateAttributeUsecase {
	return &UpdateAttributeUsecase{repo}
}

func (uc *UpdateAttributeUsecase) Execute(ctx context.Context, input UpdateAttributeInput) (*UpdateAttributeOutput, error) {
	attr, err := uc.repo.FindById(ctx, input.RankId, input.Id)
	if err != nil {
		return nil, err
	}
	if attr == nil {
		return nil, &ResourceNotFoundError{name: "attribute", id: input.Id}
	}
	if err := uc.repo.Update(ctx, input); err != nil {
		return nil, err
	}
	return &UpdateAttributeOutput{
		Id:     input.Id,
		Name:   input.Name,
		Desc:   input.Desc,
		Order:  input.Order,
		RankId: input.RankId,
	}, nil
}

type DeleteAttributeInput struct {
	RankId string
	Id     string
}

type DeleteAttributeOutput struct{}

type DeleteAttributeUsecase struct {
	repo repository.AttributeRepository
}

func NewDeleteAttributeUsecase(repo repository.AttributeRepository) *DeleteAttributeUsecase {
	return &DeleteAttributeUsecase{repo}
}

func (uc *DeleteAttributeUsecase) Execute(ctx context.Context, input DeleteAttributeInput) (*DeleteAttributeOutput, error) {
	attr, err := uc.repo.FindById(ctx, input.RankId, input.Id)
	if err != nil {
		return nil, err
	}
	if attr == nil {
		return nil, &ResourceNotFoundError{name: "attribute", id: input.Id}
	}
	if err := uc.repo.Delete(ctx, attr); err != nil {
		return nil, err
	}
	return &DeleteAttributeOutput{}, nil
}
