package entity

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/josimarz/ranking-backend/internal/validator"
)

func TestValidateAttribute(t *testing.T) {
	v := validator.New()
	attr := NewAttribute("Graphics", "Evaluate graphic capacity", 1, uuid.NewString())
	ValidateAttribute(v, attr)
	if got := v.Valid(); !got {
		t.Errorf("attribute validation failed: got %v, want %v", got, true)
	}

	attr.Id = "123"
	attr.Name = ""
	attr.Desc = "Phasellus libero felis, mollis id purus ut, interdum malesuada nunc. Aenean dignissim ac mi eu sollicitudin. Proin viverra eget mi non condimentum. Proin quis dolor velit."
	attr.Order = -1
	attr.RankId = "123"
	ValidateAttribute(v, attr)
	if got := v.Valid(); got {
		t.Errorf("attribute validation failed: got %v, want %v", got, false)
	}

	want := map[string]string{
		"id":          "must be a valid UUID",
		"name":        "must be between 3 and 15 characters long",
		"description": "must be a maximum of 150 characters long",
		"order":       "must be a positive number",
		"rank_id":     "must be a valid UUID",
	}
	if got := v.Errors(); !reflect.DeepEqual(got, want) {
		t.Errorf("validation returned wrong errors: got %v, want %v", got, want)
	}
}
