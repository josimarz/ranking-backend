package repository

import (
	"context"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

type EntryRepository interface {
	Create(context.Context, *entity.Entry) error
	FindById(context.Context, string, string) (*entity.Entry, error)
	Update(context.Context, *entity.Entry) error
	Delete(context.Context, *entity.Entry) error
}
