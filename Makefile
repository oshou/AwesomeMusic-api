# Prerequisite
# - Install golangci-lint
#   https://github.com/golangci/golangci-lint#install

export GO111MODULE=on

LIST=./domain/repository ./domain/model ./usecase
GO ?= $(shell which go)
GOLINT ?= golangci-lint
PG_DUMP ?= /usr/local/bin/pg_dump
DB_USER ?= root
DB_HOST ?= 127.0.0.1
DB_NAME ?= work
DEPLOY_REPO = "oshou/awesome-music-api"
API_CMD_PATH = "cmd/api/main.go"
BINARY_NAME = main

clean:
	$(GO) mod tidy

fmt: clean
	$(GO) fmt ./...

lint: fmt
	$(GOLINT) run

schema:
	$(PG_DUMP) -h $(DB_HOST) -U $(DB_USER) -s $(DB_NAME) -f _db/schema.sql

#mockgen:
#	@LIST="$(LIST)"; \
#	for x in $$LIST; do \
#		echo "$$x"; \
#		mockgen -source "$$x" --destination mock/"$$x"/"$$x".go; \
#	done

test: lint
	$(GO) test ./...

cov: lint
	$(GO) test ./... -cover

mockgen:
	@LIST="$(LIST)"; \
	for x in $$LIST; do \
		echo "$$x"; \
		mockgen -source "$$x" --destination mock/"$$x"/"$$x".go; \
	done

build_local: test
	$(GO) build -o $(BINARY_NAME) $(API_CMD_PATH)

build_prd: test
	cp -rp .env.production .env
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -a -installsuffix cgo -ldflags="-s -w" -o $(BINARY_NAME)

run:
	./$(BINARY_NAME)

.PHONY: fmt lint test cov gen_test build_local build_prd run clean
