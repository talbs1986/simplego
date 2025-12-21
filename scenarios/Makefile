DIRS = $(sort $(dir $(wildcard ./*/)))

all: deps build tidy

dep: 
	cd ${DIR} && go get ${MODULE}

deps:
	cd ${DIR} && go get -v -t ./...

build:
	make clean ; \
	cd $(DIR) ; \
	out=$(DIR:./%/=%) ; \
	go build -v -o ../../.build/"$$out" cmd/main.go ; \
	cd ..

clean:
	out=$(DIR:./%/=%) ; \
	if [ -f "../.build/$${out}" ] ; then echo deleting ".build/$${out}.coverage.out" ; rm ../.build/"$$out" ; fi ; \
	if [ -f "../.build/$${out}.scenarios.coverage.out" ] ; then echo deleting ".build/$${out}.scenarios.coverage.out" ; rm ../.build/"$${out}.scenarios.coverage.out" ; fi 

tidy:
	cd $(DIR) && go mod tidy

test:
	cd $(DIR) && go test -v ./... -count=1

test_cover:
	out=$(DIR:./%/=%) ; \
	cd $(DIR) && go test -covermode=count -coverprofile=../../.build/$${out}.scenarios.coverage.out -v ./... -count=1

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cd $(DIR) && golangci-lint run

lint_fix:
	cd $(DIR) && golangci-lint run --fix

dev_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
    		make all DIR=$$currDir ; \
		fi \
	done

test_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
    		make test DIR=$$currDir ; \
		fi \
	done

test_cover_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
    		make test_cover DIR=$$currDir ; \
		fi \
	done

lint_fix_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
    		make lint_fix DIR=$$currDir ; \
		fi \
	done

clean_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
    		make clean DIR=$$currDir ; \
		fi \
	done

.PHONY: all deps build tidy lint dev_all test_all dep lint_fix_all lint_fix clean test_cover test_cover_all clean_all

