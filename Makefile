.DEFAULT_GOAL := build
SHELL := /bin/bash
GOBIN := $(shell go env GOPATH)/bin
GO_LINT_PACKAGES := $(shell go list ./... | grep -v /vendor/)

JOBS := $(shell getconf _NPROCESSORS_CONF)
GIT_REF := $(shell git rev-parse --short HEAD)
VERSION ?= commit-$(GIT_REF)
ifneq ($(shell git status --porcelain),)
	VERSION := $(VERSION)-dirty
endif


oapi_up: oapi_gen
	cd ./api/oapi && \
	docker-compose stop && \
	docker-compose build --no-cache && \
	docker-compose up -d --build && \
	echo 'oapi start'
.PHONY: oapi_up

oapi_down:
	cd ./api/oapi && \
	docker-compose down --remove-orphans && \
    echo 'oapi stop'

oapi_gen:
	swagger-cli bundle -o ./api/oapi/build/oapi.yaml -t yaml ./api/oapi/daradara.oapi.yaml && \
	oapi-codegen -config ./api/oapi/spec.config.yaml ./api/oapi/build/oapi.yaml && \
	oapi-codegen -config ./api/oapi/types.config.yaml ./api/oapi/build/oapi.yaml && \
	oapi-codegen -config ./api/oapi/server.config.yaml ./api/oapi/build/oapi.yaml && \
	echo 'oapi generate'

db_create_migration:
	migrate create -ext sql -dir db/migrations -seq ${migration_name}

#dev_create_db:
#	docker exec -it golang_develop_db createdb --username=dev01 --owner=dev01 dev01
#
#dev_drop_db:
#	docker exec -it golang_develop_db dropdb -Udev01 dev01

dev_db_migrate_up:
	migrate -path db/migrations -database "postgres://dev01:dev01@localhost:9432/dev01?sslmode=disable" -verbose up

dev_db_migrate_force:
	migrate -path db/migrations -database "postgres://dev01:dev01@localhost:9432/dev01?sslmode=disable" force ${force_number}

dev_db_migrate_down:
	migrate -path db/migrations -database "postgres://dev01:dev01@localhost:9432/dev01?sslmode=disable" -verbose down

test_db_migrate_up:
	migrate -path db/migrations -database "postgres://test01:test01@localhost:9433/test01?sslmode=disable" -verbose up

test_db_migrate_force:
	migrate -path db/migrations -database "postgres://test01:test01@localhost:9433/test01?sslmode=disable" force ${force_number}

test_db_migrate_down:
	migrate -path db/migrations -database "postgres://test01:test01@localhost:9433/test01?sslmode=disable" -verbose down

seed:
	go run db/fixtures/main.go

gogen: fmt
	go generate ./...
.PHONY: gogen

fmt:
	go fmt ./...
.PHONY: fmt

vet: fmt
	go vet ./...
.PHONY: vet

test: vet
	go test -v ./...
	go test ./... -count=1
	go test $(args) -race -cover ./...
.PHONY: test

build: test
	go mod tidy
	go build -o cmd/daradara/binary_main cmd/daradara/main.go
.PHONY: build

install:
	@./scripts/install.sh
.PHONY: install

run:
	@make fmt
	go run cmd/daradara/main.go
