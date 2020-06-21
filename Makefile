# Prerequisite
# - Install golangci-lint
#   https://github.com/golangci/golangci-lint#install

export GO111MODULE=on

LIST=./domain/repository ./domain/model ./usecase
GO ?= $(shell which go)
GOLINT ?= golangci-lint
DEPLOY_REPO = "oshou/awesome-music-api"
BINARY_NAME = main

clean:
	$(GO) mod tidy

fmt: clean
	$(GO) fmt ./...

lint: fmt
	$(GOLINT) run

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

build_local: clean fmt lint
	cp -rp .env.local .env
	$(GO) build -o $(BINARY_NAME)

build_prd:
	cp -rp .env.production .env
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -a -installsuffix cgo -ldflags="-s -w" -o $(BINARY_NAME)

run:
	./$(BINARY_NAME)

.PHONY: fmt lint test cov gen_test build_local build_prd run clean
