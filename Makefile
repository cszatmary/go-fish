.DEFAULT_GOAL = build
GOFISH = go run main.go
COVERPKGS = ./...

# Get all dependencies
setup:
	@echo Installing tool dependencies
	@shed install
# Self-hoisted!
	@$(GOFISH) install
.PHONY: setup

build:
	@go build
.PHONY: build

# Clean all build artifacts
clean:
	@rm -rf coverage
	@rm -f go-fish
.PHONY: clean

fmt:
	@shed run goimports -- -w .
.PHONY: fmt

check-fmt:
	@./scripts/check_fmt.sh
.PHONY: check-fmt

lint:
	@shed run golangci-lint run ./...
.PHONY: lint

# Remove version installed with go install
go-uninstall:
	@rm $(shell go env GOPATH)/bin/go-fish
.PHONY: go-uninstall

# Run tests and collect coverage data
test:
	@mkdir -p coverage
	@go test -coverpkg=$(COVERPKGS) -coverprofile=coverage/coverage.txt ./...
	@go tool cover -html=coverage/coverage.txt -o coverage/coverage.html
.PHONY: test

# Run tests and print coverage data to stdout
test-ci:
	@mkdir -p coverage
	@go test -coverpkg=$(COVERPKGS) -coverprofile=coverage/coverage.txt ./...
	@go tool cover -func=coverage/coverage.txt
.PHONY: test-ci
