package entity

import (
	"github.com/google/uuid"
	"github.com/josimarz/ranking-backend/internal/validator"
)

type Attribute struct {
	Id     string
	Name   string
	Desc   string
	Order  int
	RankId string
}

func NewAttribute(name, desc string, order int, rankId string) *Attribute {
	return &Attribute{
		Id:     uuid.NewString(),
		Name:   name,
		Desc:   desc,
		Order:  order,
		RankId: rankId,
	}
}

func ValidateAttribute(v *validator.Validator, attr *Attribute) {
	v.Check(validator.IsUUID(attr.Id), "id", "must be a valid UUID")
	v.Check(len(attr.Name) >= 3 && len(attr.Name) <= 15, "name", "must be between 3 and 15 characters long")
	v.Check(len(attr.Desc) <= 150, "description", "must be a maximum of 150 characters long")
	v.Check(attr.Order > 0, "order", "must be a positive number")
	v.Check(validator.IsUUID(attr.RankId), "rank_id", "must be a valid UUID")
}
