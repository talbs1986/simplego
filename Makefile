all: deps build tidy

deps:
	cd logger && go get -v -t -d ./...

build:
	cd logger && go fmt ./...

tidy:
	cd logger && go mod tidy

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cd logger && golangci-lint run

.PHONY: all deps build tidy lint

