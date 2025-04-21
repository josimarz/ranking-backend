package entity

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/josimarz/ranking-backend/internal/validator"
)

func TestValidateEntry(t *testing.T) {
	v := validator.New()
	entry := NewEntry("Sega Mega Drive", "https://videogame.com/smd.png", Scores{"Graphics": 84, "Sound": 81}, uuid.NewString())
	ValidateEntry(v, entry)
	if got := v.Valid(); !got {
		t.Errorf("entry validation failed: got %v, want %v", got, true)
	}

	entry.Id = "123"
	entry.Name = "MD"
	entry.ImageURL = "videogame/smd.png"
	entry.RankId = "123"
	ValidateEntry(v, entry)
	if got := v.Valid(); got {
		t.Errorf("entry validation failed: got %v, want %v", got, false)
	}

	want := map[string]string{
		"id":        "must be a valid UUID",
		"name":      "must be between 3 and 60 characters long",
		"image_url": "must be a valid URL",
		"rank_id":   "must be a valid UUID",
	}
	if got := v.Errors(); !reflect.DeepEqual(got, want) {
		t.Errorf("validation returned wrong errors: got %v, want %v", got, want)
	}
}
