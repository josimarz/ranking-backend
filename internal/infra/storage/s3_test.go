package storage

import (
	"context"
	"strings"
	"testing"
)

func TestFileS3Storage(t *testing.T) {
	ctx := context.Background()
	storage := NewFileS3Storage(client)
	t.Run("Upload", func(t *testing.T) {
		file := strings.NewReader("file content")
		path := "file/path.png"
		want := "http://localhost:4566/ranking/file/path.png"
		if got, err := storage.Upload(ctx, path, file); err != nil || got != want {
			t.Errorf("Upload(%v, %v, %v) got (%v, %v), want (%v, %v)", ctx, path, file, got, err, want, nil)
		}
	})
}
