# Prerequisite
# - Install golangci-lint
#   https://github.com/golangci/golangci-lint#install

export GO111MODULE=on

GO_CMD=$(shell which go)
GOTESTS_CMD=gotests
GOLANGCILINT_CMD=golangci-lint
DEPLOY_REPO="oshou/awesome-music-api"
BINARY_NAME=main

## 各環境別ビルドコマンド
lint:
	$(GO_CMD) fmt ./...
	$(GOLANGCILINT_CMD) run

test:
	$(GO_CMD) test -v ./...

coverage:
	$(GO_CMD) test ./... -cover

gen_test:
	$(GOTESTS_CMD) -all -w ./*

build_local:
	cp -rp .env.local .env
	$(GO_CMD) build -o $(BINARY_NAME)

build_prd:
	cp -rp .env.local .env
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -a -installsuffix cgo -ldflags="-s -w" -o $(BINARY_NAME)

deploy_hub:
	docker build -t $(DEPLOY_REPO) .
	docker push $(DEPLOY_REPO):latest

run:
	./$(BINARY_NAME)

clean:
	$(GO_CMD) mod tidy
	$(GO_CMD) clean
