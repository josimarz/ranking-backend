package repository

import (
	"context"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

type AttributeRepository interface {
	Create(context.Context, *entity.Attribute) error
	FindById(context.Context, string, string) (*entity.Attribute, error)
	Update(context.Context, *entity.Attribute) error
	Delete(context.Context, *entity.Attribute) error
}
