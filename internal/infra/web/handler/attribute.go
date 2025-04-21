package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/josimarz/ranking-backend/internal/domain/entity"
	"github.com/josimarz/ranking-backend/internal/domain/usecase"
	"github.com/josimarz/ranking-backend/internal/validator"
)

type PostAttributeHandler struct {
	baseHandler
	uc *usecase.CreateAttributeUsecase
}

func NewPostAttributeHandler(logger *slog.Logger, uc *usecase.CreateAttributeUsecase) *PostAttributeHandler {
	return &PostAttributeHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *PostAttributeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name  string `json:"name"`
		Desc  string `json:"description"`
		Order int    `json:"order"`
	}
	if err := h.readJSON(w, r, &body); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	rankId := r.PathValue("rankId")
	attr := entity.NewAttribute(body.Name, body.Desc, body.Order, rankId)
	v := validator.New()
	if entity.ValidateAttribute(v, attr); !v.Valid() {
		h.failedValidationResponse(w, r, v.Errors())
		return
	}
	output, err := h.uc.Execute(r.Context(), attr)
	if err != nil {
		h.serverErrorResponse(w, r, err)
		return
	}
	if err := h.writeJSON(w, http.StatusCreated, output, nil); err != nil {
		h.serverErrorResponse(w, r, err)
	}
}

type GetAttributeHandler struct {
	baseHandler
	uc *usecase.FindAttributeUsecase
}

func NewGetAttributeHandler(logger *slog.Logger, uc *usecase.FindAttributeUsecase) *GetAttributeHandler {
	return &GetAttributeHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *GetAttributeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := usecase.FindAttributeInput{
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

type PutAttributeHandler struct {
	baseHandler
	uc *usecase.UpdateAttributeUsecase
}

func NewPutAttributeHandler(logger *slog.Logger, uc *usecase.UpdateAttributeUsecase) *PutAttributeHandler {
	return &PutAttributeHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *PutAttributeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name  string `json:"name"`
		Desc  string `json:"description"`
		Order int    `json:"order"`
	}
	if err := h.readJSON(w, r, &body); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	rankId := r.PathValue("rankId")
	id := r.PathValue("id")
	attr := &entity.Attribute{
		Id:     id,
		Name:   body.Name,
		Desc:   body.Desc,
		Order:  body.Order,
		RankId: rankId,
	}
	v := validator.New()
	if entity.ValidateAttribute(v, attr); !v.Valid() {
		h.failedValidationResponse(w, r, v.Errors())
		return
	}
	output, err := h.uc.Execute(r.Context(), attr)
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

type DeleteAttributeHandler struct {
	baseHandler
	uc *usecase.DeleteAttributeUsecase
}

func NewDeleteAttributeHandler(logger *slog.Logger, uc *usecase.DeleteAttributeUsecase) *DeleteAttributeHandler {
	return &DeleteAttributeHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *DeleteAttributeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := usecase.DeleteAttributeInput{
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
		"message": "attribute successfully deleted",
	}
	if err := h.writeJSON(w, http.StatusOK, data, nil); err != nil {
		h.serverErrorResponse(w, r, err)
	}
}
