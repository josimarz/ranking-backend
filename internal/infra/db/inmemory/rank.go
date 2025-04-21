package inmemory

import (
	"context"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
)

var (
	ranks map[string]*entity.Rank = make(map[string]*entity.Rank)
)

type RankInMemoryRepository struct{}

func (r *RankInMemoryRepository) Create(ctx context.Context, rank *entity.Rank) error {
	ranks[rank.Id] = rank
	return nil
}

func (r *RankInMemoryRepository) FindById(ctx context.Context, id string) (*entity.Rank, error) {
	if rank, ok := ranks[id]; ok {
		return rank, nil
	}
	return nil, nil
}

func (r *RankInMemoryRepository) Update(ctx context.Context, rank *entity.Rank) error {
	ranks[rank.Id] = rank
	return nil
}

func (r *RankInMemoryRepository) Delete(ctx context.Context, rank *entity.Rank) error {
	delete(ranks, rank.Id)
	return nil
}
