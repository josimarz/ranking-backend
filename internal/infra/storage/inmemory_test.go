package storage

import (
	"context"
	"strings"
	"testing"
)

func TestInMemoryStorage(t *testing.T) {
	ctx := context.Background()
	storage := NewInMemoryStorage()
	t.Run("Upload", func(t *testing.T) {
		path := "file/path.png"
		file := strings.NewReader("file content")
		want := "http://fake-url/file.png"
		if got, err := storage.Upload(ctx, path, file); got != want || err != nil {
			t.Errorf("Upload(%v, %v, %v) got (%v, %v), want (%v, %v)", ctx, path, file, got, err, want, nil)
		}
	})
}
