# Prerequisite
# - Install golangci-lint
#   - brew install "golangci/tap/golangci-lint"

export GO111MODULE=on

GOCMD=$(shell which go)
DEPLOY_REPO="oshou/awesome-music-api"
BINARY_NAME=main

## 各環境別ビルドコマンド
lint:
	$(GOCMD) fmt ./...
	$(GOCMD) vet ./...
	golangci-lint run --enable-all

build_local:
	make lint
	cp -rp .env.local .env
	$(GOCMD) build -o $(BINARY_NAME)

build_prd:
	make lint
	cp -rp .env.local .env
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -a -installsuffix cgo -ldflags="-s -w" -o $(BINARY_NAME)

deploy_hub:
	make build_prd
	docker build -t $(DEPLOY_REPO) .
	docker push $(DEPLOY_REPO):latest

run:
	./$(BINARY_NAME)

clean:
	$(GOCMD) clean
