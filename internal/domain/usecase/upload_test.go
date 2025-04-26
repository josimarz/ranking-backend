package usecase

import (
	"context"
	"strings"
	"testing"

	"github.com/josimarz/ranking-backend/internal/infra/storage"
)

func TestUploadUsecase(t *testing.T) {
	ctx := context.Background()
	storage := storage.NewInMemoryStorage()
	uc := NewUploadUsecase(storage)
	t.Run("Execute", func(t *testing.T) {
		input := UploadInput{
			RankId:   "58b95233-b624-44c5-8175-cbbbd03a37ef",
			Filename: "file.png",
			File:     strings.NewReader("file content"),
		}
		var want *UploadOutput
		if got, err := uc.Execute(ctx, input); got == nil || err != nil {
			t.Errorf("Execute(%v, %v) got (%v, %v), want (%v, %v)", ctx, input, got, err, want, nil)
		}
	})
}
