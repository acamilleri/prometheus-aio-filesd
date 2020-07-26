.PHONY: lint run up build binary test

VERSION = $(shell git rev-parse --verify HEAD)

run: lint run

build: binary build

up:
	docker-compose up -d

binary:
	go build -ldflags "-X main.version=$(VERSION)" cmd/prometheus-aio-filesd.go

build:
	docker build --build-arg version=$(VERSION) .

run:
	FILESD_PROVIDER_NAME=docker FILESD_WRITER_NAME=stdout go run cmd/prometheus-aio-filesd.go

lint:
	@go get -u golang.org/x/lint/golint
	golint ./...

test:
	go test -v -cover ./...
