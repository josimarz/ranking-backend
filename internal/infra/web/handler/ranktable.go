package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/josimarz/ranking-backend/internal/domain/usecase"
)

type GetRankTableHandler struct {
	baseHandler
	uc *usecase.FindRankTableUsecase
}

func NewGetRankTableHandler(logger *slog.Logger, uc *usecase.FindRankTableUsecase) *GetRankTableHandler {
	return &GetRankTableHandler{
		baseHandler: baseHandler{logger},
		uc:          uc,
	}
}

func (h *GetRankTableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := usecase.FindRankTableInput{
		Id: r.PathValue("id"),
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
