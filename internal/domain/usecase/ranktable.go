package usecase

import (
	"context"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/domain/repository"
)

type FindRankTableInput struct {
	Id string
}

type attributeOutput struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"description"`
	Order int    `json:"order"`
}

type entryOutput struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	ImageURL string        `json:"image_url"`
	Scores   entity.Scores `json:"scores"`
}

type FindRankTableOutput struct {
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Public  bool              `json:"public"`
	Attrs   []attributeOutput `json:"attributes"`
	Entries []entryOutput     `json:"entries"`
}

type FindRankTableUsecase struct {
	repo repository.RankTableRepository
}

func NewFindRankTableUsecase(repo repository.RankTableRepository) *FindRankTableUsecase {
	return &FindRankTableUsecase{repo}
}

func (uc *FindRankTableUsecase) Execute(ctx context.Context, input FindRankTableInput) (*FindRankTableOutput, error) {
	table, err := uc.repo.FindById(ctx, input.Id)
	if err != nil {
		return nil, err
	}
	if table == nil {
		return nil, &ResourceNotFoundError{name: "rank", id: input.Id}
	}
	output := &FindRankTableOutput{
		Id:     table.Id,
		Name:   table.Name,
		Public: table.Public,
	}
	for _, attr := range table.Attrs {
		output.Attrs = append(output.Attrs, attributeOutput{
			Id:    attr.Id,
			Name:  attr.Name,
			Desc:  attr.Desc,
			Order: attr.Order,
		})
	}
	for _, entry := range table.Entries {
		output.Entries = append(output.Entries, entryOutput{
			Id:       entry.Id,
			Name:     entry.Name,
			ImageURL: entry.ImageURL,
			Scores:   entry.Scores,
		})
	}
	return output, nil
}
