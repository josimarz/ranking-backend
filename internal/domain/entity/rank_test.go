package entity

import (
	"reflect"
	"testing"

	"github.com/josimarz/ranking-backend/internal/validator"
)

func TestValidateRank(t *testing.T) {
	v := validator.New()
	rank := NewRank("Video Game Consoles", true)
	ValidateRank(v, rank)
	if got := v.Valid(); !got {
		t.Errorf("rank validation failed: got %v, want %v", got, true)
	}

	rank.Id = ""
	rank.Name = ""
	ValidateRank(v, rank)
	if got := v.Valid(); got {
		t.Errorf("rank validation failed: got %v, want %v", got, false)
	}

	want := map[string]string{
		"id":   "must be a valid UUID",
		"name": "must be between 5 and 50 characters long",
	}
	if got := v.Errors(); !reflect.DeepEqual(got, want) {
		t.Errorf("rank validation returned wrong errors: got %v, want %v", got, want)
	}
}
