package storage

import (
	"context"
	"io"
)

type FileStorage interface {
	Upload(context.Context, string, io.Reader) (string, error)
}
