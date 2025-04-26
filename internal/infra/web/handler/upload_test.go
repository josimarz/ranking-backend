package handler

import (
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/josimarz/ranking-backend/internal/domain/usecase"
	"github.com/josimarz/ranking-backend/internal/infra/storage"
)

func TestPostFileHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	storage := storage.NewInMemoryStorage()
	uc := usecase.NewUploadUsecase(storage)
	h := NewPostFileHandler(logger, uc)
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			dataPart, err := writer.CreateFormFile("image", "image.png")
			if err != nil {
				t.Fatal(err)
			}
			fileContent := getFile()
			if _, err := io.Copy(dataPart, fileContent); err != nil {
				t.Fatal(err)
			}
			if err := writer.Close(); err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest("POST", "/rank/{id}/file", body)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", "910cbfa5-526f-4781-8e89-76d0a35ca861")
			req.Header.Set("Content-Type", writer.FormDataContentType())
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}
		})
		t.Run("400", func(t *testing.T) {
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			dataPart, err := writer.CreateFormFile("file", "image.png")
			if err != nil {
				t.Fatal(err)
			}
			fileContent := strings.NewReader("invalid file content")
			if _, err := io.Copy(dataPart, fileContent); err != nil {
				t.Fatal(err)
			}
			if err := writer.Close(); err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest("POST", "/rank/{id}/file", body)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", "910cbfa5-526f-4781-8e89-76d0a35ca861")
			req.Header.Set("Content-Type", writer.FormDataContentType())
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusBadRequest)
			}
		})
		t.Run("415", func(t *testing.T) {
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			dataPart, err := writer.CreateFormFile("image", "image.png")
			if err != nil {
				t.Fatal(err)
			}
			fileContent := strings.NewReader("invalid file content")
			if _, err := io.Copy(dataPart, fileContent); err != nil {
				t.Fatal(err)
			}
			if err := writer.Close(); err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest("POST", "/rank/{id}/file", body)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", "910cbfa5-526f-4781-8e89-76d0a35ca861")
			req.Header.Set("Content-Type", writer.FormDataContentType())
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusUnsupportedMediaType {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusUnsupportedMediaType)
			}
		})
	})
}

func getFile() io.Reader {
	base64Str := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQAAAAA3bvkkAAAACklEQVR4AWNgAAAAAgABc3UBGAAAAABJRU5ErkJggg=="
	index := strings.Index(base64Str, ",")
	if index == -1 {
		log.Fatal("invalid base64 string")
		return nil
	}
	base64Data := base64Str[index+1:]
	decodedData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		log.Fatal("error decoding base64 string")
		return nil
	}
	reader := strings.NewReader(string(decodedData))
	return reader
}
