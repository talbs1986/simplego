DIRS = $(sort $(dir $(wildcard ./*/)))

list_dirs:
	@echo $(DIRS)

all: deps build tidy

dep: 
	cd ${DIR} && go get ${MODULE}

deps:
	cd ${DIR} && go get -v -t ./...

build:
	cd $(DIR) && go build ./...

tidy:
	cd $(DIR) && go mod tidy

test:
	cd $(DIR) && go test -v ./... -count=1

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cd $(DIR) && golangci-lint run

lint_fix:
	cd $(DIR) && golangci-lint run --fix

dev_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
			if [ "$$currDir" != "./scenarios/" ] ; then \
    			make all DIR=$$currDir ; \
			else \
				cd scenarios && make dev_all && cd .. ; \
			fi \
		fi \
	done

test_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
			if [ "$$currDir" != "./scenarios/" ] ; then \
				make test DIR=$$currDir ; \
			else \
				cd scenarios && make test_all && cd .. ; \
			fi \
		fi \
	done

lint_fix_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
			if [ "$$currDir" != "./scenarios/" ] ; then \
				make lint_fix DIR=$$currDir ; \
			else \
				cd scenarios && make lint_fix_all && cd .. ; \
			fi \
		fi \
	done

.PHONY: all deps build tidy lint dev_all test_all dep lint_fix_all lint_fix list_dirs

