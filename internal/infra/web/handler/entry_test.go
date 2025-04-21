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

func TestPostEntryHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.EntryInMemoryRepository{}
	uc := usecase.NewCreateEntryUsecase(repo)
	h := NewPostEntryHandler(logger, uc)
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("201", func(t *testing.T) {
			buf := []byte(`{
				"name": "Super Nintendo Entertainment System",
				"image_url": "https://videogame.com/snes.png",
				"scores": {
					"Controls": 84,
					"Graphics": 90,
					"Sound": 88
				}
			}`)
			req, err := http.NewRequest("POST", "/rank/{rankId}/entry", bytes.NewBuffer(buf))
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
				"name": "Super Nintendo Entertainment System",
				"image_url": "https://videogame.com/snes.png",
				"scores": {
					"Controls": 84,
					"Graphics": 90,
					"Sound"
				}
			}`)
			req, err := http.NewRequest("POST", "/rank/{rankId}/entry", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusBadRequest)
			}
		})
		t.Run("422", func(t *testing.T) {
			buf := []byte(`{
				"name": "Super Nintendo Entertainment System",
				"image_url": "videgame/snes.png",
				"scores": {
					"Controls": 84,
					"Graphics": 90,
					"Sound": 88
				}
			}`)
			req, err := http.NewRequest("POST", "/rank/{rankId}/entry", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "ad183111-c022-4812-9081-ebec903a3903")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusUnprocessableEntity {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusUnprocessableEntity)
			}
		})
	})
}

func TestGetEntryHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.EntryInMemoryRepository{}
	uc := usecase.NewFindEntryUsecase(repo)
	h := NewGetEntryHandler(logger, uc)
	repo.Create(context.Background(), &mock.Entries[0])
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/rank/{rankId}/entry/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "d10961ca-e9ed-4d3b-b086-f756a3118894")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}
			want := `{"id":"d10961ca-e9ed-4d3b-b086-f756a3118894","name":"Neo Geo CD","image_url":"https://videogame.com/neo-geo-cd.png","scores":{"Controls":90,"Graphics":97,"Sound":97},"rank_id":"1ac85e34-cb6f-40c9-97bb-16267877bb13"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("404", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/rank/{rankId}/entry/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "2e15269d-8db9-4fff-a28e-fa5dd7501c20")
			req.SetPathValue("id", "7ca55092-80f7-40cb-a629-361037baa6be")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusNotFound)
			}
			want := `{"error":"entry not found: 7ca55092-80f7-40cb-a629-361037baa6be"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
	})
}

func TestPutEntryHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.EntryInMemoryRepository{}
	uc := usecase.NewUpdateEntryUsecase(repo)
	h := NewPutEntryHandler(logger, uc)
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			buf := []byte(`{
				"name": "Nintendo 64",
				"image_url": "https://videogame.com/n64.png",
				"scores": {
					"Controls": 91,
					"Graphics": 93,
					"Sound": 90
				}
			}`)
			req, err := http.NewRequest("PUT", "/rank/{rankId}/entry/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "d10961ca-e9ed-4d3b-b086-f756a3118894")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}
			want := `{"id":"d10961ca-e9ed-4d3b-b086-f756a3118894","name":"Nintendo 64","image_url":"https://videogame.com/n64.png","scores":{"Controls":91,"Graphics":93,"Sound":90},"rank_id":"1ac85e34-cb6f-40c9-97bb-16267877bb13"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("400", func(t *testing.T) {
			buf := []byte(`{
				"name": "Nintendo 64",
				"image_url": "https://videogame.com/n64.png",
				"scores": {
					"Controls": 91,
					"Graphics": 93,
					"Sound"
				}
			}`)
			req, err := http.NewRequest("PUT", "/rank/{rankId}/entry/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "d10961ca-e9ed-4d3b-b086-f756a3118894")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusBadRequest)
			}
		})
		t.Run("404", func(t *testing.T) {
			buf := []byte(`{
				"name": "Nintendo 64",
				"image_url": "https://videogame.com/n64.png",
				"scores": {
					"Controls": 91,
					"Graphics": 93,
					"Sound": 90
				}
			}`)
			req, err := http.NewRequest("PUT", "/rank/{rankId}/entry/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "464b7cf7-a0c6-49a1-a528-a9357f5ceef3")
			req.SetPathValue("id", "b59287e3-e544-44a9-a20a-f28520c6beea")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusNotFound)
			}
			want := `{"error":"entry not found: b59287e3-e544-44a9-a20a-f28520c6beea"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("422", func(t *testing.T) {
			buf := []byte(`{
				"name": "64",
				"image_url": "https://videogame.com/n64.png",
				"scores": {
					"Controls": 91,
					"Graphics": 93,
					"Sound": 90
				}
			}`)
			req, err := http.NewRequest("PUT", "/rank/{rankId}/entry/{id}", bytes.NewBuffer(buf))
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "464b7cf7-a0c6-49a1-a528-a9357f5ceef3")
			req.SetPathValue("id", "b59287e3-e544-44a9-a20a-f28520c6beea")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusUnprocessableEntity {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusUnprocessableEntity)
			}
		})
	})
}

func TestDeleteEntryHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.EntryInMemoryRepository{}
	uc := usecase.NewDeleteEntryUsecase(repo)
	h := NewDeleteEntryHandler(logger, uc)
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/rank/{rankId}/entry/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "d10961ca-e9ed-4d3b-b086-f756a3118894")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}
			want := `{"message":"entry successfully deleted"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("404", func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/rank/{rankId}/entry/{id}", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("rankId", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			req.SetPathValue("id", "d10961ca-e9ed-4d3b-b086-f756a3118894")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusNotFound)
			}
			want := `{"error":"entry not found: d10961ca-e9ed-4d3b-b086-f756a3118894"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
	})
}
