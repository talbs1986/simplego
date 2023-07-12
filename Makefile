all: deps build tidy

deps:
	cd ${DIR} && go get -v -t -d ./...

build:
	cd $(DIR) && go fmt ./...

tidy:
	cd $(DIR) && go mod tidy

tidy:
	cd $(DIR) && go test -v ./... -count=1

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cd $(DIR) && golangci-lint run

.PHONY: all deps build tidy lint

