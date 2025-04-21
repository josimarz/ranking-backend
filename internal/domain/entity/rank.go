package entity

import (
	"github.com/google/uuid"
	"github.com/josimarz/ranking-backend/internal/validator"
)

type Rank struct {
	Id     string
	Name   string
	Public bool
}

func NewRank(name string, public bool) *Rank {
	return &Rank{
		Id:     uuid.NewString(),
		Name:   name,
		Public: public,
	}
}

func ValidateRank(v *validator.Validator, rank *Rank) {
	v.Check(validator.IsUUID(rank.Id), "id", "must be a valid UUID")
	v.Check(len(rank.Name) >= 5 && len(rank.Name) <= 50, "name", "must be between 5 and 50 characters long")
}
