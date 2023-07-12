all: deps build tidy

deps:
	cd logger && go get -v -t -d ./...

build:
	cd logger && go fmt ./...

tidy:
	cd logger && go mod tidy

.PHONY: all deps build tidy

