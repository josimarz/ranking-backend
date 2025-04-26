package repository

import (
	"context"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

type RankTableRepository interface {
	FindById(context.Context, string) (*entity.RankTable, error)
}
