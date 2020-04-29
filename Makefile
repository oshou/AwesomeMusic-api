# Prerequisite
# - Install golangci-lint
#   https://github.com/golangci/golangci-lint#install

export GO111MODULE=on

GO ?= $(shell which go)
GOTESTS ?= gotests
GOLANGCILINT ?= golangci-lint
DEPLOY_REPO = "oshou/awesome-music-api"
BINARY_NAME = main

.PHONY: lint
lint:
	$(GO) fmt ./...
	$(GOLANGCILINT) run

.PHONY: test
test:
	$(GO) test -v ./...

.PHONY: cov
cov:
	$(GO) test ./... -cover

.PHONY: gen_test
gen_test:
	$(GOTESTS) -all -w ./*

.PHONY: build_local
build_local:
	cp -rp .env.local .env
	$(GO) build -o $(BINARY_NAME)

.PHONY: build_prd
build_prd:
	cp -rp .env.local .env
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -a -installsuffix cgo -ldflags="-s -w" -o $(BINARY_NAME)

.PHONY: run
run:
	./$(BINARY_NAME)

.PHONY: clean
clean:
	$(GO) mod tidy
	$(GO) clean
