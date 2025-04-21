package inmemory

import (
	"context"
	"fmt"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

var (
	attrs map[string]*entity.Attribute = make(map[string]*entity.Attribute)
)

type AttributeInMemoryRepository struct{}

func (r *AttributeInMemoryRepository) Create(ctx context.Context, attr *entity.Attribute) error {
	key := fmt.Sprintf("%s/%s", attr.RankId, attr.Id)
	attrs[key] = attr
	return nil
}

func (r *AttributeInMemoryRepository) FindById(ctx context.Context, rankId, id string) (*entity.Attribute, error) {
	key := fmt.Sprintf("%s/%s", rankId, id)
	if attr, ok := attrs[key]; ok {
		return attr, nil
	}
	return nil, nil
}

func (r *AttributeInMemoryRepository) Update(ctx context.Context, attr *entity.Attribute) error {
	key := fmt.Sprintf("%s/%s", attr.RankId, attr.Id)
	attrs[key] = attr
	return nil
}

func (r *AttributeInMemoryRepository) Delete(ctx context.Context, attr *entity.Attribute) error {
	key := fmt.Sprintf("%s/%s", attr.RankId, attr.Id)
	delete(attrs, key)
	return nil
}
