package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/josimarz/ranking-backend/internal/domain/usecase"
	"github.com/josimarz/ranking-backend/internal/infra/db/ddb"
	"github.com/josimarz/ranking-backend/internal/infra/web/handler"
	"github.com/josimarz/ranking-backend/internal/infra/web/server"
	"github.com/josimarz/ranking-backend/internal/repository"
)

type repositories struct {
	rank      repository.RankRepository
	attr      repository.AttributeRepository
	entry     repository.EntryRepository
	rankTable repository.RankTableRepository
}

type usecases struct {
	createRank    *usecase.CreateRankUsecase
	findRank      *usecase.FindRankUsecase
	updateRank    *usecase.UpdateRankUsecase
	deleteRank    *usecase.DeleteRankUsecase
	createAttr    *usecase.CreateAttributeUsecase
	findAttr      *usecase.FindAttributeUsecase
	updateAttr    *usecase.UpdateAttributeUsecase
	deleteAttr    *usecase.DeleteAttributeUsecase
	createEntry   *usecase.CreateEntryUsecase
	findEntry     *usecase.FindEntryUsecase
	updateEntry   *usecase.UpdateEntryUsecase
	deleteEntry   *usecase.DeleteEntryUsecase
	findRankTable *usecase.FindRankTableUsecase
}

type application struct {
	logger   *slog.Logger
	cfg      config
	client   *dynamodb.Client
	repos    *repositories
	usecases *usecases
	handlers server.Handlers
	server   server.Server
}

func newApplication() *application {
	return &application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		cfg:    loadConfig(),
	}
}

func (a *application) start() {
	a.connectToDatabase()
	a.initRepositories()
	a.initUsecases()
	a.initHandlers()
	a.startServer()
}

func (a *application) connectToDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := ddb.NewDynamodbClient(ctx)
	if err != nil {
		a.logger.Error(err.Error())
		os.Exit(1)
	}
	a.client = client
}

func (a *application) initRepositories() {
	a.repos = &repositories{
		rank:      ddb.NewRankDynamodbRepository(a.client),
		attr:      ddb.NewAttributeDynamodbRepository(a.client),
		entry:     ddb.NewEntryDynamodbRepository(a.client),
		rankTable: ddb.NewRankTableDynamodbRepository(a.client),
	}
}

func (a *application) initUsecases() {
	a.usecases = &usecases{
		createRank:    usecase.NewCreateRankUsecase(a.repos.rank),
		findRank:      usecase.NewFindRankUsecase(a.repos.rank),
		updateRank:    usecase.NewUpdateRankUsecase(a.repos.rank),
		deleteRank:    usecase.NewDeleteRankUsecase(a.repos.rank),
		createAttr:    usecase.NewCreateAttributeUsecase(a.repos.attr),
		findAttr:      usecase.NewFindAttributeUsecase(a.repos.attr),
		updateAttr:    usecase.NewUpdateAttributeUsecase(a.repos.attr),
		deleteAttr:    usecase.NewDeleteAttributeUsecase(a.repos.attr),
		createEntry:   usecase.NewCreateEntryUsecase(a.repos.entry),
		findEntry:     usecase.NewFindEntryUsecase(a.repos.entry),
		updateEntry:   usecase.NewUpdateEntryUsecase(a.repos.entry),
		deleteEntry:   usecase.NewDeleteEntryUsecase(a.repos.entry),
		findRankTable: usecase.NewFindRankTableUsecase(a.repos.rankTable),
	}
}

func (a *application) initHandlers() {
	a.handlers = server.Handlers{
		"POST /rank":                           handler.NewPostRankHandler(a.logger, a.usecases.createRank),
		"GET /rank/{id}":                       handler.NewGetRankHandler(a.logger, a.usecases.findRank),
		"PUT /rank/{id}":                       handler.NewPutRankHandler(a.logger, a.usecases.updateRank),
		"DELETE /rank/{id}":                    handler.NewDeleteRankHandler(a.logger, a.usecases.deleteRank),
		"POST /rank/{rankId}/attribute":        handler.NewPostAttributeHandler(a.logger, a.usecases.createAttr),
		"GET /rank/{rankId}/attribute/{id}":    handler.NewGetAttributeHandler(a.logger, a.usecases.findAttr),
		"PUT /rank/{rankId}/attribute/{id}":    handler.NewPutAttributeHandler(a.logger, a.usecases.updateAttr),
		"DELETE /rank/{rankId}/attribute/{id}": handler.NewDeleteAttributeHandler(a.logger, a.usecases.deleteAttr),
		"POST /rank/{rankId}/entry":            handler.NewPostEntryHandler(a.logger, a.usecases.createEntry),
		"GET /rank/{rankId}/entry/{id}":        handler.NewGetEntryHandler(a.logger, a.usecases.findEntry),
		"PUT /rank/{rankId}/entry/{id}":        handler.NewPutEntryHandler(a.logger, a.usecases.updateEntry),
		"DELETE /rank/{rankId}/entry/{id}":     handler.NewDeleteEntryHandler(a.logger, a.usecases.deleteEntry),
		"GET /rank/{id}/table":                 handler.NewGetRankTableHandler(a.logger, a.usecases.findRankTable),
	}
}

func (a *application) startServer() {
	addr := fmt.Sprintf(":%d", a.cfg.port)
	a.server = *server.NewServer(addr, a.handlers)
	a.logger.Info("starting server", "addr", addr)
	if err := a.server.Start(); err != nil {
		a.logger.Error(err.Error())
		os.Exit(1)
	}
}
