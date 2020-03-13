# Prerequisite
# - go vet
# - go fmt
# - golangci-lint
#   - brew install "golangci/tap/golangci-lint"

#export GOROOT=/usr/local/go
#export GOPATH=/root/go
#export GOBIN=/root/go/bin
export GO111MODULE=on

#GOCMD=/usr/local/go/bin/go
GOCMD=$(shell which go)
DEPLOY_REPO="oshou/awesome-music-api"
BINARY_NAME=main

## 各環境別ビルドコマンド
build_local:
	$(GOCMD) fmt ./...
	$(GOCMD) vet ./...
	golangci-lint run --enable-all
	cp -rp .env.local .env
	$(GOCMD) build -o $(BINARY_NAME)

build_prd:
	$(GOCMD) fmt ./...
	$(GOCMD) vet ./...
	golangci-lint run --enable-all
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
