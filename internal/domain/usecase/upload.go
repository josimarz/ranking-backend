package usecase

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/josimarz/ranking-backend/internal/infra/storage"
)

type UploadInput struct {
	RankId   string
	Filename string
	File     io.Reader
}

type UploadOutput struct {
	URL string `json:"url"`
}

type UploadUsecase struct {
	storage storage.FileStorage
}

func NewUploadUsecase(storage storage.FileStorage) *UploadUsecase {
	return &UploadUsecase{storage}
}

func (uc *UploadUsecase) Execute(ctx context.Context, input UploadInput) (*UploadOutput, error) {
	ext := filepath.Ext(input.Filename)
	filename := uuid.NewString()
	path := fmt.Sprintf("%s/%s%s", input.RankId, filename, ext)
	url, err := uc.storage.Upload(ctx, path, input.File)
	if err != nil {
		return nil, err
	}
	return &UploadOutput{url}, nil
}
