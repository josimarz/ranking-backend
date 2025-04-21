package server

import (
	"net/http"
	"testing"
)

type mockHandler struct{}

func (h *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte("Hello, World!")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func TestServer(t *testing.T) {
	addr := ":8080"
	handlers := Handlers{
		"GET /": &mockHandler{},
	}
	NewServer(addr, handlers)
}
