include .envrc

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo 'Are you sure? [y/N]' && read ans && [ $${ans:-N} = y ]

## setup/local: setup resources to run application locally
.PHONY: setup/local
setup/local:
	@echo "Setting up resources to run application locally..."
	@cd scripts; ./setup.sh

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@go run ./cmd/api

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${DB_DSN} up

## tidy: format all .go files and tidy module dependencies
.PHONY: tidy
tidy:
	@echo 'Formatting .go files...'
	go fmt ./...
	@echo 'Tidying module dependencies...'
	go mod tidy
	@echo 'Verifying and vendoring module dependencies...'
	go mod verify
	go mod vendor

## test: run tests
.PHONY: test
test:
	@echo 'Running tests...'
	go test -v ./...

## test/cover: generate test coverage report
.PHONY: test/cover
test/cover:
	@echo 'Generating test coverage...'
	go test -coverprofile=c.out ./...
	@echo 'Generating test report...'
	go tool cover -html=c.out

## audit: run quality control checks
.PHONY: audit
audit:
	@echo 'Checking module dependencies...'
	go mod tidy -diff
	go mod verify
	@echo 'Vetting code...'
	go vet ./...
	$(shell go env GOPATH)/bin/staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api