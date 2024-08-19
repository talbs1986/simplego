DIRS = $(sort $(dir $(wildcard ./*/)))
all: deps build tidy

dep: 
	cd ${DIR} && go get -d ${MODULE}

deps:
	cd ${DIR} && go get -v -t -d ./...

build:
	make clean ; \
	cd $(DIR) ; \
	tmp=$(DIR) ; \
	out=$(DIR:./%/=%) ; \
	go build -v -o ../../.build/"$$out" cmd/main.go ; \
	cd ..

clean:
	tmp="$(DIR)" ; \
	out=$(tmp:./%/=%) ; \
	if [ -f "$$out" ] ; then rm ../../.build/"$$out" ; fi 

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
    		make all DIR=$$currDir ; \
		fi \
	done

test_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
    		make test DIR=$$currDir ; \
		fi \
	done

lint_fix_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "./" ] ; then \
    		make lint_fix DIR=$$currDir ; \
		fi \
	done

.PHONY: all deps build tidy lint dev_all test_all dep lint_fix_all lint_fix clean

