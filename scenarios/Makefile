DIRS = $(shell ls -d */ | sed 's|/||g')

list_dirs:
	@echo $(DIRS)

all: deps build tidy

dep: 
	cd ${DIR} && go get ${MODULE}

deps:
	cd ${DIR} && go get -v -t ./...

build:
	make clean ; \
	cd $(DIR) ; \
	go build -v -o ../../.build/"$$DIR" cmd/main.go ; \
	cd ..

clean:
	if [ -f "../.build/$${DIR}" ] ; then echo deleting ".build/$${DIR}.coverage.out" ; rm ../.build/"$$DIR" ; fi ; \
	if [ -f "../.build/$${DIR}.scenarios.coverage.out" ] ; then echo deleting ".build/$${DIR}.scenarios.coverage.out" ; rm ../.build/"$${DIR}.scenarios.coverage.out" ; fi 

tidy:
	cd $(DIR) && go mod tidy

test:
	cd $(DIR) && go test -v ./... -count=1

test_cover:
	cd $(DIR) && go test -covermode=count -coverprofile=../../.build/$${DIR}.scenarios.coverage.out -v ./... -count=1

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
    	make test_cover DIR=$$currDir ; \
	done

lint_fix_all:
	for currDir in $(DIRS) ; do \
    	make lint_fix DIR=$$currDir ; \
	done

clean_all:
	for currDir in $(DIRS) ; do \
    	make clean DIR=$$currDir ; \
	done

.PHONY: all deps build tidy lint dev_all test_all dep lint_fix_all lint_fix clean test_cover test_cover_all clean_all list_dirs

