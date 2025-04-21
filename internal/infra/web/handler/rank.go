package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/domain/usecase"
	"github.com/josimarz/ranking-backend/internal/validator"
)

type PostRankHandler struct {
	baseHandler
	uc *usecase.CreateRankUsecase
}

func NewPostRankHandler(logger *slog.Logger, uc *usecase.CreateRankUsecase) *PostRankHandler {
	return &PostRankHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *PostRankHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name   string `json:"name"`
		Public bool   `json:"public"`
	}
	if err := h.readJSON(w, r, &body); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	rank := entity.NewRank(body.Name, body.Public)
	v := validator.New()
	if entity.ValidateRank(v, rank); !v.Valid() {
		h.failedValidationResponse(w, r, v.Errors())
		return
	}
	output, err := h.uc.Execute(r.Context(), rank)
	if err != nil {
		h.serverErrorResponse(w, r, err)
		return
	}
	if err := h.writeJSON(w, http.StatusCreated, output, nil); err != nil {
		h.serverErrorResponse(w, r, err)
	}
}

type GetRankHandler struct {
	baseHandler
	uc *usecase.FindRankUsecase
}

func NewGetRankHandler(logger *slog.Logger, uc *usecase.FindRankUsecase) *GetRankHandler {
	return &GetRankHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *GetRankHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := usecase.FindRankInput{
		Id: r.PathValue("id"),
	}
	output, err := h.uc.Execute(r.Context(), input)
	if err != nil {
		var notFoundErr *usecase.ResourceNotFoundError
		if errors.As(err, &notFoundErr) {
			h.notFoundResponse(w, r, err)
			return
		}
		h.badRequestResponse(w, r, err)
		return
	}
	if err := h.writeJSON(w, http.StatusOK, output, nil); err != nil {
		h.serverErrorResponse(w, r, err)
	}
}

type PutRankHandler struct {
	baseHandler
	uc *usecase.UpdateRankUsecase
}

func NewPutRankHandler(logger *slog.Logger, uc *usecase.UpdateRankUsecase) *PutRankHandler {
	return &PutRankHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *PutRankHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name   string `json:"name"`
		Public bool   `json:"public"`
	}
	if err := h.readJSON(w, r, &body); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	id := r.PathValue("id")
	rank := &entity.Rank{
		Id:     id,
		Name:   body.Name,
		Public: body.Public,
	}
	v := validator.New()
	if entity.ValidateRank(v, rank); !v.Valid() {
		h.failedValidationResponse(w, r, v.Errors())
		return
	}
	output, err := h.uc.Execute(r.Context(), rank)
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

type DeleteRankHandler struct {
	baseHandler
	uc *usecase.DeleteRankUsecase
}

func NewDeleteRankHandler(logger *slog.Logger, uc *usecase.DeleteRankUsecase) *DeleteRankHandler {
	return &DeleteRankHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *DeleteRankHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := usecase.DeleteRankInput{
		Id: r.PathValue("id"),
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
		"message": "rank successfully deleted",
	}
	if err := h.writeJSON(w, http.StatusOK, data, nil); err != nil {
		h.serverErrorResponse(w, r, err)
	}
}
