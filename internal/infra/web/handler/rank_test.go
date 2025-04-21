package handler

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josimarz/ranking-backend/internal/domain/usecase"
	"github.com/josimarz/ranking-backend/internal/infra/db/inmemory"
	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestPostRankHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.RankInMemoryRepository{}
	uc := usecase.NewCreateRankUsecase(repo)
	h := NewPostRankHandler(logger, uc)
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("201", func(t *testing.T) {
			buf := []byte(`{
				"name": "Video Game Consoles",
				"public": true
			}`)
			req, err := http.NewRequest("POST", "/rank", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusCreated {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
			}
		})
		t.Run("400", func(t *testing.T) {
			buf := []byte(`{
				"name": "Video Game Consoles",
				"public":
			}`)
			req, err := http.NewRequest("POST", "/rank", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusBadRequest)
			}
		})
		t.Run("422", func(t *testing.T) {
			buf := []byte(`{
				"name": "",
				"public": true
			}`)
			req, err := http.NewRequest("POST", "/rank", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusUnprocessableEntity {
				t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusUnprocessableEntity)
			}
		})
	})
}

func TestGetRankHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.RankInMemoryRepository{}
	uc := usecase.NewFindRankUsecase(repo)
	h := NewGetRankHandler(logger, uc)
	rank := mock.Rank
	repo.Create(context.Background(), &rank)
	t.Run("ServeHttp", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/rank/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", rank.Id)
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}
			want := `{"id":"1ac85e34-cb6f-40c9-97bb-16267877bb13","name":"Video Game Consoles","public":true}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("404", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/rank/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", "738344c2-6a48-4acd-bf18-5727e285ba5d")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusNotFound)
			}
		})
	})
}

func TestPutRankHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.RankInMemoryRepository{}
	uc := usecase.NewUpdateRankUsecase(repo)
	h := NewPutRankHandler(logger, uc)
	rank := mock.Rank
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			buf := []byte(`{
				"name": "Video Games",
				"public": false
			}`)
			req, err := http.NewRequest("PUT", "/rank/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", rank.Id)
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}
			want := `{"id":"1ac85e34-cb6f-40c9-97bb-16267877bb13","name":"Video Games","public":false}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("400", func(t *testing.T) {
			buf := []byte(`{
				"name": "Video Games",
				"public":
			}`)
			req, err := http.NewRequest("PUT", "/rank/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", rank.Id)
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusBadRequest)
			}
		})
		t.Run("404", func(t *testing.T) {
			buf := []byte(`{
				"name": "Video Games",
				"public": false
			}`)
			req, err := http.NewRequest("PUT", "/rank/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", "9812bc0e-129b-42b6-896d-48c798168160")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusNotFound)
			}
			want := `{"error":"rank not found: 9812bc0e-129b-42b6-896d-48c798168160"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("422", func(t *testing.T) {
			buf := []byte(`{
				"name": "",
				"public": false
			}`)
			req, err := http.NewRequest("PUT", "/rank/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", rank.Id)
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusUnprocessableEntity {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusUnprocessableEntity)
			}
		})
	})
}

func TestDeleteRankHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.RankInMemoryRepository{}
	uc := usecase.NewDeleteRankUsecase(repo)
	h := NewDeleteRankHandler(logger, uc)
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/rank/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}
			want := `{"message":"rank successfully deleted"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("404", func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/rank/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusNotFound)
			}
			want := `{"error":"rank not found: 1ac85e34-cb6f-40c9-97bb-16267877bb13"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
	})
}
