CURRENT_VERSION = $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.1.0")
NEXT_VERSION = v0.1.0
GO_FILES = $(shell find . -type f -name '*.go')
GOFMT_FILES = $(shell find . -name '*.go' | grep -v vendor)

default: build

.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@gofmt -w $(GOFMT_FILES)
	@goimports -w $(GOFMT_FILES)

.PHONY: lint
lint:
	@echo "Linting code..."
	@golangci-lint run ./...

.PHONY: test
test: deps
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: integration
integration: build
	@echo "Running integration tests..."
	@go test -v ./testing/integration/...

.PHONY: build
build: deps fmt
	@echo "Building..."
	@go build -o dist/tflint-ruleset-ibm

.PHONY: install
install: build
	@echo "Installing..."
	@mkdir -p ~/.tflint.d/plugins
	@cp dist/tflint-ruleset-ibm ~/.tflint.d/plugins/

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf dist/
	@rm -f coverage.txt

.PHONY: release
release:
	@echo "Releasing..."
	@git tag $(NEXT_VERSION)
	@git push origin $(NEXT_VERSION)
	@goreleaser release --rm-dist

.PHONY: docs
docs:
	@echo "Generating docs..."
	@go run ./tools/docs-gen

.PHONY: all
all: clean deps fmt lint test build