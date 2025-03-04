# simple_cinema

Vulcan Cinema is a web application that allows users book seats for a movie.

## Features
- Create room
- Update Room information
- Delete Room
- Get Room detail
- Booking seat
- Get available seats
- Get booking detail
- Cancel booking

## Tech Stack
- Go
- Gin
- Gorm
- PostgreSQL
- Redis
- Docker

## Architecture
- RESTful API
- Repository Pattern
- Dependency Injection
- SOLID Principle

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
