APP_NAME=swapi-planets
VERSION?=0.0.1
GOCMD=go
CGO_ENABLED=0

.PHONY: all

all: build
	./bin/api

clean: clean-bin
	$(GOCMD) clean

clean-bin:
	rm -rf ./bin/*

clean-gen:
	rm -rf ./gen/*

gen: clean-gen
	protoc \
		--go_out . \
		--go_opt paths=import \
    --go-grpc_out . \
		--go-grpc_opt paths=import \
		./proto/*.proto

build:
	$(GOCMD) build -o ./bin ./...

get:
	$(GOCMD) get ./...
	$(GOCMD) mod verify
	$(GOCMD) mod tidy

test:
	$(GOCMD) test -v -race ./...

test-mongo:
	$(GOCMD) test -v -race ./... -mongo

test-get-data:
	$(GOCMD) test -v -race ./... -get-data

req-create:
	./scripts/create-curl.sh

req-read-all:
	./scripts/read-all-curl.sh

req-read-one:
	./scripts/read-one-curl.sh

req-delete:
	./scripts/delete-curl.sh

lint:
	golangci-lint run

up:
	docker-compose up --build -d

down:
	docker-compose down
