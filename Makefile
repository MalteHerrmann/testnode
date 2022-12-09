#!/usr/bin/make -f
BUILD_DIR ?= $(CURDIR)/build

# ------------------
# Build
#
BUILD_TARGETS := build install
$(BUILD_TARGETS): go.sum $(BUILD_DIR)/
	go $@ $(BUILD_ARGS) ./...

$(BUILD_DIR)/:
	mkdir -p $(BUILD_DIR)/

build: BUILD_ARGS=-o $(BUILD_DIR)/

# ------------------
# Golang
#
go.sum: go.mod
	@echo "Ensure dependencies have not been modified ..." >&2
	go mod verify
	go mod tidy

# ------------------
# Testing
#
PACKAGES_UNIT=$(shell go list ./...)
TEST_TARGETS := test-unit test-unit-cover

# Unit tests without coverage report
test-unit: ARGS=-timeout=1m -race
test-unit: TEST_PACKAGES=$(PACKAGES_UNIT)

test-unit-cover: ARGS=-timeout=15m -race -coverprofile=coverage.txt -covermode=atomic
test-unit-cover: TEST_PACKAGES=$(PACKAGES_UNIT)

$(TEST_TARGETS): run-tests
run-tests:
	go test -mod=readonly $(ARGS) $(TEST_PACKAGES)

# ------------------
# Vulnerability Check
#
vulncheck: $(BUILD_DIR)/
	GOBIN=$(BUILD_DIR) go install golang.org/x/vuln/cmd/govulncheck@latest
	$(BUILD_DIR)/govulncheck ./...
