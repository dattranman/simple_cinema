# simple_cinema

Simple Cinema is a web application that allows users book seats for a movie.

## Features
- Create room
- Update Room information
- Delete Room
- Get Room detail by room id
- Booking seat
- Get available seats by room id
- Cancel booking

## Tech Stack
- Go (https://go.dev/doc/install)
- Gin (https://github.com/gin-gonic/gin)
- Gorm (https://gorm.io/)
- PostgreSQL (https://www.postgresql.org/)
- Redis (https://redis.io/)
- Docker (https://docs.docker.com/get-docker/)

## Architecture
- RESTful API
- Repository Pattern
- Dependency Injection
- SOLID Principle

## Prerequisites
- Docker (https://docs.docker.com/get-docker/)
- Make (https://www.gnu.org/software/make/manual/make.html)
### Step by step on new server
- install docker: https://docs.docker.com/get-docker/
- install make: https://www.gnu.org/software/make/manual/make.html
## How to run
- Clone the repository
- Run `make db` to start the database
- Run `make migrate` to run the migrations
- Run `make run` to start the application in local mode

## How to test
- Run `make test` to run the test
- Run `make test-coverage` to run the test and generate the coverage report

### How to create new migration
- Run `make create_migrate name=<migration_name>` to create a new migration file sql
- Run `make create_migrate_code name=<migration_name>` to create a new migration file use golang
 
### Summary about architecture code
- api: handle process request from client and call to app layer
- app: handle business logic and call to store layer
- config: handle logic read from file config.yaml or environment variable
- docs: store document generate from swagger
- migration: define the migration file for the database

- model: define model use in the application
    - model/schema: define the schema for the database
    - model/request: define the request for the api
    - model/response: define the response for the api
- store: handle database operation and call to schema layer
    - store/postgres: define the postgres operation
    - store/redis: define the redis operation
- util: define the util function for the application

### Simple flow test:
- Create room => Get room detail => Get available seats => Booking seat