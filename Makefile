DIRS          := $(shell ls -d */ | sed 's|/||g')
SCENARIOS_DIR := scenarios

define run_or_scenarios
	@if [ "$(DIR)" = "$(SCENARIOS_DIR)" ]; then \
		cd $(SCENARIOS_DIR) && $(MAKE) $(1); \
	else \
		$(2); \
	fi
endef

list_dirs:
	@echo $(DIRS)

all:
	$(call run_or_scenarios, \
		dev_all, \
		$(MAKE) deps  DIR=$(DIR) && \
		$(MAKE) build DIR=$(DIR) && \
		$(MAKE) tidy  DIR=$(DIR) \
	)

dep:
	cd $(DIR) && go get $(MODULE)

deps:
	cd $(DIR) && go get -v -t ./...

build:
	cd $(DIR) && go build ./...

tidy:
	cd $(DIR) && go mod tidy

test:
	$(call run_or_scenarios, \
		test_all, \
		cd $(DIR) && go test -v ./... -count=1 \
	)

test_cover:
	$(call run_or_scenarios, \
		test_cover_all, \
		cd $(DIR) && \
		go test -covermode=count \
			-coverprofile=../.build/$${DIR}.coverage.out \
			-v ./... -count=1 \
	)

lint_install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	$(call run_or_scenarios, \
		lint_all, \
		cd $(DIR) && golangci-lint run \
	)

lint_fix:
	$(call run_or_scenarios, \
		lint_fix_all, \
		cd $(DIR) && golangci-lint run --fix \
	)

clean:
	$(call run_or_scenarios, \
		clean_all, \
		if [ -f ".build/$${DIR}" ]; then \
			echo deleting ".build/$${DIR}.coverage.out"; \
			rm .build/"$$DIR"; \
		fi; \
		if [ -f ".build/$${DIR}.coverage.out" ]; then \
			echo deleting ".build/$${DIR}.coverage.out"; \
			rm .build/"$${DIR}.coverage.out"; \
		fi \
	)

dev_all:
	for currDir in $(DIRS); do \
		$(MAKE) all DIR=$$currDir; \
	done

test_all:
	for currDir in $(DIRS); do \
		$(MAKE) test DIR=$$currDir; \
	done

test_cover_all:
	for currDir in $(DIRS); do \
		$(MAKE) test_cover DIR=$$currDir; \
	done

lint_fix_all:
	for currDir in $(DIRS); do \
		$(MAKE) lint DIR=$$currDir; \
	done

lint_fix_all:
	for currDir in $(DIRS); do \
		$(MAKE) lint_fix DIR=$$currDir; \
	done

clean_all:
	for currDir in $(DIRS); do \
		$(MAKE) clean DIR=$$currDir; \
	done

.PHONY: all deps build tidy lint dev_all test_all dep lint_fix_all lint_fix lint_all \
	list_dirs test_cover test_cover_all clean_all clean