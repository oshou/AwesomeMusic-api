export GO111MODULE=on

GO ?= $(shell which go)
PG_DUMP ?= /usr/local/bin/pg_dump
PG_DUMPALL ?= /usr/local/bin/pg_dumpall
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

pg_local:
	cp -rp .env.local .env
	docker run -d --rm \
		--name postgres \
		-p 5432:5432 \
		-e POSTGRES_DB=postgres \
		-e POSTGRES_HOST_AUTH_METHOD=trust \
		postgres:12-alpine

apidoc:
	docker run --rm \
		--name apidoc \
		-p 10080:80 \
		-v $(shell pwd)/docs/:/usr/share/nginx/html/openapi/ \
		-e SPEC_URL=openapi/openapi.yaml \
		redocly/redoc

dbdoc:
	docker run --rm \
		--name dbdoc \
		-v $(PWD):/work \
		--network=host \
		k1low/tbls doc postgres://postgres:postgres@localhost/postgres?sslmode=disable

migrate: $(GOPATH)/bin/sql-migrate
	sql-migrate up -config=_db/config.yaml

rollback: $(GOPATH)/bin/sql-migrate
	sql-migrate down -config=_db/config.yaml

schema:
	$(PGDUMP_PATH) -h $(DB_HOST) -U $(DB_USER) -s $(DB_NAME) -f _db/schema.sql

dumpall:
	$(PG_DUMP) -h $(DB_HOST) -U $(DB_USER) $(DB_NAME) -f _db/dumpall.sql

restore:
	psql -h $(DB_HOST) -U $(DB_USER) $(DB_NAME) < _db/dumpall.sql

seed:
	go run cmd/seed/main.go

clean:
	go mod tidy

fmt: clean
	go fmt ./...

lint: fmt $(GOPATH)/bin/golangci-lint
	golangci-lint run

apitest:
	docker build -t tavern -f deployments/apitest/Dockerfile ./test/ \
	&& docker run --rm tavern

test: fmt
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
