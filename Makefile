export GO111MODULE=on

GO_ENV=local

PGDUMP_PATH ?= /usr/local/bin/pg_dump
API_PATH ?= "cmd/api/main.go"

DB_USER ?= root
DB_HOST ?= 127.0.0.1
DB_NAME ?= postgres

DEPLOY_REPO = "oshou/awesome-music-api"
BINARY_NAME = main

LIST=./domain/repository ./domain/model ./usecase

$(GOPATH)/bin/sql-migrate:
	go get -v github.com/rubenv/sql-migrate/...

$(GOPATH)/bin/golangci-lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

$(GOPATH)/bin/gotests:
	go get -u github.com/cweill/gotests/...

gendoc:
	docker run -it --rm -p 10080:80 \
		-v $(shell pwd)/docs/:/usr/share/nginx/html/openapi/ \
		-e SPEC_URL=openapi/openapi.yaml \
		redocly/redoc

pg_local:
	cp -rp .env.local .env
	docker-compose -f deployments/postgres/docker-compose.yml up -d

migrate: $(GOPATH)/bin/sql-migrate
	sql-migrate up -config=_db/config.yml -env=$(GO_ENV)

rollback: $(GOPATH)/bin/sql-migrate
	sql-migrate down -config=_db/config.yml -env=$(GO_ENV)

schema:
	$(PGDUMP_PATH) -h $(DB_HOST) -U $(DB_USER) -s $(DB_NAME) -f _db/schema.sql

clean:
	go mod tidy

fmt: clean
	go fmt ./...

lint: fmt $(GOPATH)/bin/golangci-lint
	golangci-lint run

test: lint
	go test ./...

cov: lint
	go test ./... -cover

mockgen:
	@LIST="$(LIST)"; \
	for x in $$LIST; do \
		echo "$$x"; \
		mockgen -source "$$x" --destination mock/"$$x"/"$$x".go; \
	done

build_local:
	go build -o $(BINARY_NAME) $(API_PATH)

build_prd: test
	cp -rp .env.production .env
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o $(BINARY_NAME)

run:
	go run cmd/api/main.go

.PHONY: dep schema migrate rollback clean fmt lint test cov mockgen build_local build_prd run redoc
