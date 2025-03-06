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
- Go (https://go.dev/doc/install)
- Make (https://www.gnu.org/software/make/manual/make.html)

## How to run
- Clone the repository
- Run `make db` to start the database
- Run `make run` to start the application
- Run `make build-docker` to build the docker image
- Run `make local` to start the application in local mode
- Run `make prod` to start the application in production mode

## How to test
- Run `make test` to run the test
- Run `make test-coverage` to run the test and generate the coverage report

## How to deploy to new server
- clone the repository
- create file config/config.yaml like config/config.yaml.example and fill in the correct values
- create file .env like .env.example and fill in the correct values
- Run `make build-docker` to build the docker image
- Run `make db` to start the database
- Run `make goose` to run the migrations
- Run `make prod` to start the application in production mode
- Run `make run` to start the application
