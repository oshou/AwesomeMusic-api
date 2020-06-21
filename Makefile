# Prerequisite
# - Install golangci-lint
#   https://github.com/golangci/golangci-lint#install

export GO111MODULE=on

GO ?= $(shell which go)
GOLANGCILINT ?= golangci-lint
DEPLOY_REPO = "oshou/awesome-music-api"
BINARY_NAME = main

.PHONY: fmt lint test cov gen_test build_local build_prd run clean

clean:
	$(GO) mod tidy

fmt: clean
	$(GO) fmt ./...

lint: fmt
	$(GOLANGCILINT) run

test: lint
	$(GO) test ./...

cov: lint
	$(GO) test ./... -cover

build_local: clean fmt lint
	cp -rp .env.local .env
	$(GO) build -o $(BINARY_NAME)

build_prd:
	cp -rp .env.local .env
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -a -installsuffix cgo -ldflags="-s -w" -o $(BINARY_NAME)

run:
	./$(BINARY_NAME)

