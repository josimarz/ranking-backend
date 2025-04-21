package inmemory

import "github.com/josimarz/ranking-backend/internal/domain/entity"

func ClearDatabase() {
	ranks = make(map[string]*entity.Rank)
	attrs = make(map[string]*entity.Attribute)
	entries = make(map[string]*entity.Entry)
}
