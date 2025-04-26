package handler

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/josimarz/ranking-backend/internal/domain/usecase"
)

type PostFileHandler struct {
	baseHandler
	uc *usecase.UploadUsecase
}

func NewPostFileHandler(logger *slog.Logger, uc *usecase.UploadUsecase) *PostFileHandler {
	return &PostFileHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *PostFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("image")
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	if !h.isImage(buf) {
		h.errorResponse(w, r, http.StatusUnsupportedMediaType, "invalid file type")
		return
	}
	input := usecase.UploadInput{
		RankId:   r.PathValue("id"),
		Filename: handler.Filename,
		File:     bytes.NewReader(buf),
	}
	output, err := h.uc.Execute(r.Context(), input)
	if err != nil {
		h.serverErrorResponse(w, r, err)
		return
	}
	if err := h.writeJSON(w, http.StatusOK, output, nil); err != nil {
		h.serverErrorResponse(w, r, err)
	}
}

func (*PostFileHandler) isImage(buf []byte) bool {
	contentType := http.DetectContentType(buf)
	return strings.HasPrefix(contentType, "image/")
}
