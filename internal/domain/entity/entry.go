package entity

import (
	"github.com/google/uuid"
	"github.com/josimarz/ranking-backend/internal/validator"
)

type Scores map[string]int

type Entry struct {
	Id       string
	Name     string
	ImageURL string
	Scores   Scores
	RankId   string
}

func NewEntry(name, imageURL string, scores Scores, rankId string) *Entry {
	return &Entry{
		Id:       uuid.NewString(),
		Name:     name,
		ImageURL: imageURL,
		Scores:   scores,
		RankId:   rankId,
	}
}

func ValidateEntry(v *validator.Validator, entry *Entry) {
	v.Check(validator.IsUUID(entry.Id), "id", "must be a valid UUID")
	v.Check(len(entry.Name) >= 3 && len(entry.Name) <= 60, "name", "must be between 3 and 60 characters long")
	v.Check(validator.IsURL(entry.ImageURL), "image_url", "must be a valid URL")
	v.Check(validator.IsUUID(entry.RankId), "rank_id", "must be a valid UUID")
}
