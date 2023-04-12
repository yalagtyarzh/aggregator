export

include .env

build-user-api:
	go build -o ./.bin/user ./cmd/user-api/main.go

run-user-api: build-user-api
	./.bin/user

build-admin-api:
	go build -o ./.bin/admin ./cmd/admin-api/main.go

run-admin-api: build-admin-api
	./.bin/user

compose-build:
	docker-compose -f docker-compose.yml build

compose-up:
	docker-compose -f docker-compose.yml up --remove-orphans --abort-on-container-exit

compose-down:
	docker-compose -f docker-compose.yml down --remove-orphans

migrate-up-local:
	docker build -t db-migrate db
	docker run --rm --net backend -v "${CURDIR}/db":/src:ro -w /src db-migrate sql-migrate up -env="local-pg"

migrate-down-local:
	docker build -t db-migrate db
	docker run --rm --net backend -v "${CURDIR}/db":/src:ro -w /src db-migrate sql-migrate down -env="local-pg"

linter-golangci:
	golangci-lint run

lint:
	docker run --rm -v "${CURDIR}":/app:ro -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v
