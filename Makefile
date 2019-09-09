#export GOROOT=/usr/local/go
#export GOPATH=/root/go
#export GOBIN=/root/go/bin
export GO111MODULE=on

GOCMD=/usr/local/go/bin/go
BINARY_NAME=main

## 各環境別ビルドコマンド
build_local:
	go vet ./...
	go fmt ./...
	cp -rp .env.local .env
	$(GOCMD) build -o $(BINARY_NAME)

build_prd:
	go vet ./...
	go fmt ./...
	cp -rp .env.local .env
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -a -installsuffix cgo -ldflags="-s -w" -o $(BINARY_NAME)

deploy_hub:
	make build_prd
	docker build -t oshou/awesome-music-api .
	docker push oshou/awesome-music-api:latest

run:
	./$(BINARY_NAME)

clean:
	$(GOCMD) clean
