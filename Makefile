.DEFAULT_GOAL := build

oapi_up:
	oapi-codegen -package openapi_service -generate "spec" ./api/oapi/oapi.yaml > ./internal/api/infrastructure/api/openapi_service/spec.gen.go && \
	cd ./api/oapi && \
	docker-compose down && \
	docker-compose build --no-cache && \
	docker-compose up -d --build && \
	cd ../../ && \
	echo 'oapi start'

oapi_down:
	cd ./api/oapi && \
	docker-compose down && \
	cd ../../ && \
    echo 'oapi stop'

oapi_gen:
	oapi-codegen -package openapi_service -generate "spec" ./api/oapi/oapi.yaml > ./internal/api/infrastructure/api/openapi_service/spec.gen.go && \
	oapi-codegen -package openapi_service -generate "types" ./api/oapi/oapi.yaml > ./internal/api/infrastructure/api/openapi_service/types.gen.go && \
	oapi-codegen -package openapi_service -generate "server" ./api/oapi/oapi.yaml > ./internal/api/infrastructure/api/openapi_service/server.gen.go && \
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

fmt:
	go fmt ./...
.PHONY: fmt

lint: fmt
	staticcheck
.PHONY: lint

vet: fmt
	go vet ./...
.PHONY: vet

build: vet
	go mod tidy
	go build cmd/daradara/main.go
.PHONY: build

run:
	go run cmd/daradara/main.go
