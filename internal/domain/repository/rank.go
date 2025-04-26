package repository

import (
	"context"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

type RankRepository interface {
	Create(context.Context, *entity.Rank) error
	FindById(context.Context, string) (*entity.Rank, error)
	Update(context.Context, *entity.Rank) error
	Delete(context.Context, *entity.Rank) error
}
