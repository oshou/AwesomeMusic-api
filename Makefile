#export GOROOT=/usr/local/go
#export GOPATH=/root/go
#export GOBIN=/root/go/bin
export GO111MODULE=on

GOCMD=/usr/local/go/bin/go
BINARY_NAME=main

## 各環境別ビルドコマンド
build_local:
	cp -rp .env.local .env
	$(GOCMD) build -o $(BINARY_NAME)
	go vet ./...
	go fmt ./...

run:
	./$(BINARY_NAME)

clean:
	$(GOCMD) clean
