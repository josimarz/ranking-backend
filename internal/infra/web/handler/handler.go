package handler

import (
	"encoding/json"
	"log/slog"
	"maps"
	"net/http"
)

const (
	maxBytes = 1_048_576
)

type baseHandler struct {
	logger *slog.Logger
}

func (h *baseHandler) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.logError(r, err)
	msg := "the server has encountered a problem and could not process your request"
	h.errorResponse(w, r, http.StatusInternalServerError, msg)
}

func (h *baseHandler) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.errorResponse(w, r, http.StatusNotFound, err.Error())
}

func (h *baseHandler) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (h *baseHandler) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	h.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (h *baseHandler) errorResponse(w http.ResponseWriter, r *http.Request, status int, msg any) {
	data := map[string]any{"error": msg}
	if err := h.writeJSON(w, status, data, nil); err != nil {
		h.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *baseHandler) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}

func (h *baseHandler) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}
	maps.Insert(w.Header(), maps.All(headers))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(output)
	return nil
}

func (h *baseHandler) logError(r *http.Request, err error) {
	h.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
}
