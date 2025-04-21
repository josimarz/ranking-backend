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

func TestPostAttributeHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.AttributeInMemoryRepository{}
	uc := usecase.NewCreateAttributeUsecase(repo)
	h := NewPostAttributeHandler(logger, uc)
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("201", func(t *testing.T) {
			buf := []byte(`{
				"name": "Graphics",
				"description": "Evaluate the graphic capacity of the console",
				"order": 1
			}`)
			req, err := http.NewRequest("POST", "/rank/{rankId}/attribute", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusCreated {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
			}
		})
		t.Run("400", func(t *testing.T) {
			buf := []byte(`{
				"name": "Graphics",
				"description": "Evaluate the graphic capacity of the console",
				"order"
			}`)
			req, err := http.NewRequest("POST", "/rank/{rankId}/attribute", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "c329b8ae-8ac8-47c5-962c-63acb429255e")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusBadRequest)
			}
		})
		t.Run("422", func(t *testing.T) {
			buf := []byte(`{
				"name": "Graphics",
				"description": "Evaluate the graphic capacity of the console",
				"order": -3
			}`)
			req, err := http.NewRequest("POST", "/rank/{rankId}/attribute", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "c329b8ae-8ac8-47c5-962c-63acb429255e")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusUnprocessableEntity {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusUnprocessableEntity)
			}
		})
	})
}

func TestGetAttributeHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.AttributeInMemoryRepository{}
	uc := usecase.NewFindAttributeUsecase(repo)
	h := NewGetAttributeHandler(logger, uc)
	attr := mock.Attrs[0]
	repo.Create(context.Background(), &attr)
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/rank/{rankId}/attribute/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "be44503b-1fac-4d5a-aae0-0239159bdc4a")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong satus code: got %v, want %v", status, http.StatusOK)
			}
			want := `{"id":"be44503b-1fac-4d5a-aae0-0239159bdc4a","name":"Controls","description":"Evaluate the quality and accessibility of controls","order":1,"rank_id":"1ac85e34-cb6f-40c9-97bb-16267877bb13"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("404", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/rank/{rankId}/attribute/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "2c94d5b0-c659-4684-b3aa-add190486fe9")
			req.SetPathValue("id", "f68f93e2-1cae-4382-b94a-9dc84ae60db8")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong satus code: got %v, want %v", status, http.StatusNotFound)
			}
			want := `{"error":"attribute not found: f68f93e2-1cae-4382-b94a-9dc84ae60db8"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
	})
}

func TestPutAttributeHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.AttributeInMemoryRepository{}
	uc := usecase.NewUpdateAttributeUsecase(repo)
	h := NewPutAttributeHandler(logger, uc)
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			buf := []byte(`{
				"name": "Design",
				"description": "Evaluate the video game console design",
				"order": 2
			}`)
			req, err := http.NewRequest("PUT", "/rank/{rankId}/attribute/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "be44503b-1fac-4d5a-aae0-0239159bdc4a")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}
			want := `{"id":"be44503b-1fac-4d5a-aae0-0239159bdc4a","name":"Design","description":"Evaluate the video game console design","order":2,"rank_id":"1ac85e34-cb6f-40c9-97bb-16267877bb13"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("400", func(t *testing.T) {
			buf := []byte(`{
				"name": "Design",
				"description": "Evaluate the video game console design",
				"order":
			}`)
			req, err := http.NewRequest("PUT", "/rank/{rankId}/attribute/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "be44503b-1fac-4d5a-aae0-0239159bdc4a")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusBadRequest)
			}
		})
		t.Run("404", func(t *testing.T) {
			buf := []byte(`{
				"name": "Design",
				"description": "Evaluate the video game console design",
				"order": 2
			}`)
			req, err := http.NewRequest("PUT", "/rank/{rankId}/attribute/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "fa4cf7b4-6130-4d84-a9fa-e3e5be4a16d2")
			req.SetPathValue("id", "2908fdf0-a9c3-4a9f-a539-2b49ec9ad8cf")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusNotFound)
			}
			want := `{"error":"attribute not found: 2908fdf0-a9c3-4a9f-a539-2b49ec9ad8cf"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("422", func(t *testing.T) {
			buf := []byte(`{
				"name": "Design",
				"description": "Evaluate the video game console design",
				"order": -2
			}`)
			req, err := http.NewRequest("PUT", "/rank/{rankId}/attribute/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "be44503b-1fac-4d5a-aae0-0239159bdc4a")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusUnprocessableEntity {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusUnprocessableEntity)
			}
		})
	})
}

func TestDeleteAttributeHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.AttributeInMemoryRepository{}
	uc := usecase.NewDeleteAttributeUsecase(repo)
	h := NewDeleteAttributeHandler(logger, uc)
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/rank/{rankId}/attribute/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "be44503b-1fac-4d5a-aae0-0239159bdc4a")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}
			want := `{"message":"attribute successfully deleted"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("404", func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/rank/{rankId}/attribute/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "be44503b-1fac-4d5a-aae0-0239159bdc4a")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusNotFound)
			}
			want := `{"error":"attribute not found: be44503b-1fac-4d5a-aae0-0239159bdc4a"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
	})
}
