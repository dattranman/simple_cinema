.EXPORT_ALL_VARIABLES:
include .env
NAME         =simple_cinema
MAIN_PATH    =
MAIN_FILE    =main.go
ENV_FILE ?= .env
VERSION ?= $(shell date +"%Y%m%d")

default: 
	@echo "USAGE: make <command>"
	@echo ""
	@echo "    db: Running postgres and redis as daemon service"
	@echo "    build: Build new bin file"
	@echo "    clean: Clean up docker containers and volumes"
	@echo ""
	@echo "    build-docker: Build docker image"
	@echo "    local: Running as daemon service"
	@echo "    goose: Run migrations"
	@echo "    down: Stop docker containers"
	@echo "    clean: Clean up docker containers and volumes"
	@echo "    create_migrate: Create a new migration file"
	@echo "    create_migrate_code: Create a new migration file use golang"
	@echo ""

.PHONY: db
## db is a command to run the postgres and redis as daemon service
db:
	@docker compose up -d simple_cinema_postgres simple_cinema_redis

.PHONY: db-local
## db-local is a command to run the postgres and redis as daemon service
db-local:
	@docker compose -f docker-compose.local.yaml -f docker-compose.yaml up -d simple_cinema_postgres simple_cinema_redis

.PHONY: build
build:
	CGO_ENABLED=0 go build  -v \
		-o simple_cinema ./main.go

.PHONY: goose
## goose is a tool that manages database migrations
goose:
	@goose up -dir migration/postgres

.PHONY: create_migrate
## create_migrate is a command to create a new migration file
create_migrate:
	@goose create -dir migration/postgres $(name) sql

.PHONY: create_migrate_code
## create_migrate_code is a command to create a new migration file use golang
create_migrate_code:
	@goose create -dir migration/postgres $(name) go

.PHONY: local
## local is a command to run the application locally
local:
	@docker compose -f docker-compose.yaml -f docker-compose.local.yaml up -d simple_cinema

.PHONY: build-docker
## build-docker is a command to build the application locally
build-docker:
	@docker build -t simple_cinema:$(VERSION) .
	@docker compose -f docker-compose.yaml -f docker-compose.local.yaml build simple_cinema

.PHONY: down
## down is a command to stop the docker containers
down:
	@docker compose -f docker-compose.yaml -f docker-compose.local.yaml down

.PHONY: clean
## clean is a command to clean up the docker containers and volumes
clean: 
	@docker compose down --remove-orphans -v

.PHONY: prod
## prod is a command to run the application in production mode
prod:
	@docker compose -f docker-compose.yaml -f docker-compose.prod.yaml up -d simple_cinema

.PHONY: test
## test is a command to run the test
test:
	@go test -v ./...

.PHONY: test-coverage
## test-coverage is a command to run the test and generate the coverage report
test-coverage:
	@go test -v ./... -coverprofile=coverage.out
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@open coverage.html
