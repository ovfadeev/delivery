.PHONY: build
build:
	go build -v ./cmd/server && go build -v ./cmd/delivery

.PHONY: test
test:
	go test -v -race -timeout 10s ./...

.DEFAULT_GOAL := build