package storage

import (
	"context"
	"io"
)

type InMemoryStorage struct {
	file io.Reader
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{}
}

func (s *InMemoryStorage) Upload(ctx context.Context, path string, file io.Reader) (string, error) {
	s.file = file
	return "http://fake-url/file.png", nil
}
