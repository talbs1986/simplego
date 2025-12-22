DIRS = $(shell ls -d */ | sed 's|/||g')

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

test_cover:
	cd $(DIR) && go test -covermode=count -coverprofile=../.build/$${DIR}.coverage.out -v ./... -count=1

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cd $(DIR) && golangci-lint run

lint_fix:
	cd $(DIR) && golangci-lint run --fix

clean:
	if [ -f ".build/$${DIR}" ] ; then echo deleting ".build/$${DIR}.coverage.out" ; rm .build/"$$DIR" ; fi ; \
	if [ -f ".build/$${DIR}.coverage.out" ] ; then echo deleting ".build/$${DIR}.coverage.out" ; rm .build/"$${DIR}.coverage.out" ; fi 

dev_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "scenarios" ] ; then \
    		make all DIR=$$currDir ; \
		else \
			cd scenarios && make dev_all && cd .. ; \
		fi \
	done

test_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "scenarios" ] ; then \
    		make test DIR=$$currDir ; \
		else \
			cd scenarios && make test_all && cd .. ; \
		fi \
	done

test_cover_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "scenarios" ] ; then \
    		make test_cover DIR=$$currDir ; \
		else \
			cd scenarios && make test_cover_all && cd .. ; \
		fi \
	done

lint_fix_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "scenarios" ] ; then \
    		make lint_fix DIR=$$currDir ; \
		else \
			cd scenarios && make lint_fix_all && cd .. ; \
		fi \
	done

clean_all:
	for currDir in $(DIRS) ; do \
		if [ "$$currDir" != "scenarios" ] ; then \
    		make clean DIR=$$currDir ; \
		else \
			cd scenarios && make clean_all && cd .. ; \
		fi \
	done

.PHONY: all deps build tidy lint dev_all test_all dep lint_fix_all lint_fix list_dirs test_cover test_cover_all clean_all clean
