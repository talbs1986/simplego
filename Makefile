DIRS = $(sort $(dir $(wildcard ./*/)))
all: deps build tidy

deps:
	cd ${DIR} && go get -v -t -d ./...

build:
	cd $(DIR) && go fmt ./...

tidy:
	cd $(DIR) && go mod tidy

test:
	cd $(DIR) && go test -v ./... -count=1

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cd $(DIR) && golangci-lint run

dev_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
    		make all DIR=$$currDir ; \
		fi \
	done

.PHONY: all deps build tidy lint dev_all

