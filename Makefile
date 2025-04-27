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
	cd scripts; ENDPOINT_URL=http://localhost:4566 TABLE_NAME=${AWS_TABLE} BUCKET_NAME=${AWS_BUCKET} ./setup.sh

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	PORT=${PORT} AWS_ENDPOINT_URL=${AWS_ENDPOINT_URL} AWS_TABLE=${AWS_TABLE} AWS_BUCKET=${AWS_BUCKET} go run ./cmd/api

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