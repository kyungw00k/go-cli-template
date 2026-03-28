MODULE := {{MODULE_PATH}}
BINARY := {{PROJECT_NAME}}
BUILD_DIR := build
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-s -w -X $(MODULE)/internal/cli.Version=$(VERSION)"

.PHONY: build install test lint clean

build:
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) ./cmd/$(BINARY)

install: build
	mkdir -p $(HOME)/.local/bin
	cp $(BUILD_DIR)/$(BINARY) $(HOME)/.local/bin/$(BINARY)

test:
	go test ./... -v

lint:
	golangci-lint run

clean:
	rm -rf $(BUILD_DIR)
