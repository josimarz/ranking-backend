package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/josimarz/ranking-backend/internal/domain/repository"
	"github.com/josimarz/ranking-backend/internal/domain/usecase"
	"github.com/josimarz/ranking-backend/internal/infra/db/ddb"
	"github.com/josimarz/ranking-backend/internal/infra/storage"
	"github.com/josimarz/ranking-backend/internal/infra/web/handler"
	"github.com/josimarz/ranking-backend/internal/infra/web/server"
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
	upload        *usecase.UploadUsecase
}

type application struct {
	logger         *slog.Logger
	dynamodbClient *dynamodb.Client
	s3Client       *s3.Client
	storage        storage.FileStorage
	repos          *repositories
	usecases       *usecases
	handlers       server.Handlers
	server         server.Server
}

func newApplication() *application {
	return &application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (a *application) start() {
	a.connectToDatabase()
	a.connectToS3()
	a.initStorage()
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
	a.dynamodbClient = client
}

func (a *application) connectToS3() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := storage.NewS3Client(ctx)
	if err != nil {
		a.logger.Error(err.Error())
		os.Exit(1)
	}
	a.s3Client = client
}

func (a *application) initStorage() {
	a.storage = storage.NewFileS3Storage(a.s3Client)
}

func (a *application) initRepositories() {
	a.repos = &repositories{
		rank:      ddb.NewRankDynamodbRepository(a.dynamodbClient),
		attr:      ddb.NewAttributeDynamodbRepository(a.dynamodbClient),
		entry:     ddb.NewEntryDynamodbRepository(a.dynamodbClient),
		rankTable: ddb.NewRankTableDynamodbRepository(a.dynamodbClient),
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
		upload:        usecase.NewUploadUsecase(a.storage),
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
		"POST /rank/{id}/file":                 handler.NewPostFileHandler(a.logger, a.usecases.upload),
	}
}

func (a *application) startServer() {
	addr := fmt.Sprintf(":%s", a.getPort())
	a.server = *server.NewServer(addr, a.handlers)
	a.logger.Info("starting server", "addr", addr)
	if err := a.server.Start(); err != nil {
		a.logger.Error(err.Error())
		os.Exit(1)
	}
}

func (*application) getPort() string {
	if value, ok := os.LookupEnv("PORT"); ok {
		return value
	}
	return "8080"
}
