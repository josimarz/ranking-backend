package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/domain/usecase"
	"github.com/josimarz/ranking-backend/internal/validator"
)

type PostEntryHandler struct {
	baseHandler
	uc *usecase.CreateEntryUsecase
}

func NewPostEntryHandler(logger *slog.Logger, uc *usecase.CreateEntryUsecase) *PostEntryHandler {
	return &PostEntryHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *PostEntryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name     string        `json:"name"`
		ImageURL string        `json:"image_url"`
		Scores   entity.Scores `json:"scores"`
	}
	if err := h.readJSON(w, r, &body); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	rankId := r.PathValue("rankId")
	entry := entity.NewEntry(body.Name, body.ImageURL, body.Scores, rankId)
	v := validator.New()
	if entity.ValidateEntry(v, entry); !v.Valid() {
		h.failedValidationResponse(w, r, v.Errors())
		return
	}
	output, err := h.uc.Execute(r.Context(), entry)
	if err != nil {
		h.serverErrorResponse(w, r, err)
		return
	}
	if err := h.writeJSON(w, http.StatusCreated, output, nil); err != nil {
		h.serverErrorResponse(w, r, err)
	}
}

type GetEntryHandler struct {
	baseHandler
	uc *usecase.FindEntryUsecase
}

func NewGetEntryHandler(logger *slog.Logger, uc *usecase.FindEntryUsecase) *GetEntryHandler {
	return &GetEntryHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *GetEntryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := usecase.FindEntryInput{
		RankId: r.PathValue("rankId"),
		Id:     r.PathValue("id"),
	}
	output, err := h.uc.Execute(r.Context(), input)
	if err != nil {
		var notFoundErr *usecase.ResourceNotFoundError
		if errors.As(err, &notFoundErr) {
			h.notFoundResponse(w, r, err)
			return
		}
		h.serverErrorResponse(w, r, err)
		return
	}
	if err := h.writeJSON(w, http.StatusOK, output, nil); err != nil {
		h.serverErrorResponse(w, r, err)
	}
}

type PutEntryHandler struct {
	baseHandler
	uc *usecase.UpdateEntryUsecase
}

func NewPutEntryHandler(logger *slog.Logger, uc *usecase.UpdateEntryUsecase) *PutEntryHandler {
	return &PutEntryHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *PutEntryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name     string        `json:"name"`
		ImageURL string        `json:"image_url"`
		Scores   entity.Scores `json:"scores"`
	}
	if err := h.readJSON(w, r, &body); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	rankId := r.PathValue("rankId")
	id := r.PathValue("id")
	entry := &entity.Entry{
		Id:       id,
		Name:     body.Name,
		ImageURL: body.ImageURL,
		Scores:   body.Scores,
		RankId:   rankId,
	}
	v := validator.New()
	if entity.ValidateEntry(v, entry); !v.Valid() {
		h.failedValidationResponse(w, r, v.Errors())
		return
	}
	output, err := h.uc.Execute(r.Context(), entry)
	if err != nil {
		var notFoundErr *usecase.ResourceNotFoundError
		if errors.As(err, &notFoundErr) {
			h.notFoundResponse(w, r, err)
			return
		}
		h.serverErrorResponse(w, r, err)
		return
	}
	if err := h.writeJSON(w, http.StatusOK, output, nil); err != nil {
		h.serverErrorResponse(w, r, err)
	}
}

type DeleteEntryHandler struct {
	baseHandler
	uc *usecase.DeleteEntryUsecase
}

func NewDeleteEntryHandler(logger *slog.Logger, uc *usecase.DeleteEntryUsecase) *DeleteEntryHandler {
	return &DeleteEntryHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *DeleteEntryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := usecase.DeleteEntryInput{
		RankId: r.PathValue("rankId"),
		Id:     r.PathValue("id"),
	}
	if _, err := h.uc.Execute(r.Context(), input); err != nil {
		var notFoundErr *usecase.ResourceNotFoundError
		if errors.As(err, &notFoundErr) {
			h.notFoundResponse(w, r, err)
			return
		}
		h.serverErrorResponse(w, r, err)
		return
	}
	data := map[string]any{
		"message": "entry successfully deleted",
	}
	if err := h.writeJSON(w, http.StatusOK, data, nil); err != nil {
		h.serverErrorResponse(w, r, err)
	}
}
