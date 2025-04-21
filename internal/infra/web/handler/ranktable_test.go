package handler

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josimarz/ranking-backend/internal/domain/usecase"
	"github.com/josimarz/ranking-backend/internal/infra/db/inmemory"
	"github.com/josimarz/ranking-backend/internal/mock"
)

func TestGetRankTableHandler(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	repo := &inmemory.RankTableInMemoryRepository{}
	uc := usecase.NewFindRankTableUsecase(repo)
	h := NewGetRankTableHandler(logger, uc)
	mockRankTable(context.Background())
	t.Run("ServeHTTP", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/rank/{id}/table", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", "1ac85e34-cb6f-40c9-97bb-16267877bb13")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
			}
			want := `{"id":"1ac85e34-cb6f-40c9-97bb-16267877bb13","name":"Video Game Consoles","public":true,"attributes":[{"id":"be44503b-1fac-4d5a-aae0-0239159bdc4a","name":"Controls","description":"Evaluate the quality and accessibility of controls","order":1},{"id":"53e1515d-7fed-4d94-8b36-4cd49b2f11be","name":"Graphics","description":"Evaluate the graphics capacity of the console","order":2},{"id":"b2ac5f2c-a65c-4eb8-a0e1-a66a6bea4aac","name":"Sound","description":"Evaluate the sound capacity of the console","order":3}],"entries":[{"id":"d10961ca-e9ed-4d3b-b086-f756a3118894","name":"Neo Geo CD","image_url":"https://videogame.com/neo-geo-cd.png","scores":{"Controls":90,"Graphics":97,"Sound":97}},{"id":"e006f3be-88a4-4891-8c8e-f1de6d6b5324","name":"Nintendo Entertainment System","image_url":"https://videogame.com/nes.png","scores":{"Controls":70,"Graphics":72,"Sound":70}},{"id":"da2b4fc6-f933-4214-b742-4f199aec2481","name":"Sega Master System","image_url":"https://videogame.com/sms.png","scores":{"Controls":73,"Graphics":78,"Sound":76}},{"id":"25658fa3-6721-42ae-8e25-7ba9c8f1cd85","name":"Sega Mega Drive","image_url":"https://videogame.com/smd.png","scores":{"Controls":80,"Graphics":84,"Sound":83}},{"id":"959c559e-db6a-4c4a-9164-f3eab305e076","name":"Super Nintendo Entertainment System","image_url":"https://videogame.com/snes.png","scores":{"Controls":84,"Graphics":89,"Sound":87}}]}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
		t.Run("404", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/rank/{id}/table", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetPathValue("id", "6a4984c3-8a2e-4aba-86fb-7154b1ac2d15")
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusNotFound)
			}
			want := `{"error":"rank not found: 6a4984c3-8a2e-4aba-86fb-7154b1ac2d15"}`
			if body := rr.Body.String(); body != want {
				t.Errorf("handler returned wrong body: got %v, want %v", body, want)
			}
		})
	})
}

func mockRankTable(ctx context.Context) {
	inmemory.ClearDatabase()
	mockRank(ctx)
	mockAttributes(ctx)
	mockEntries(ctx)
}

func mockRank(ctx context.Context) {
	repo := &inmemory.RankInMemoryRepository{}
	repo.Create(ctx, &mock.Rank)
}

func mockAttributes(ctx context.Context) {
	repo := &inmemory.AttributeInMemoryRepository{}
	for _, attr := range mock.Attrs {
		repo.Create(ctx, &attr)
	}
}

func mockEntries(ctx context.Context) {
	repo := &inmemory.EntryInMemoryRepository{}
	for _, entry := range mock.Entries {
		repo.Create(ctx, &entry)
	}
}
