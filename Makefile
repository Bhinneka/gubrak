SOURCES := $(shell find . -name '*.go' -type f -not -path './vendor/*'  -not -path '*/mocks/*')
TEST_OPTS=-covermode=atomic -v -cover

# Linter
.PHONY: lint-prepare
lint-prepare:
	@echo "Installing golangci-lint"
	@go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint:
	@golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...

# Testing
.PHONY: unittest
unittest:
	@go test -short $(TEST_OPTS) ./...

.PHONY: test
test:
	@go test $(TEST_OPTS) ./...

# Build and Installation
.PHONY: install
install:
	@go install ./...

.PHONY: uninstall
uninstall:
	@echo "Removing binaries and libraries"
	@go clean -i ./...

gubrak: $(SOURCES)
	go build -o gubrak github.com/Bhineka/gubrak/cmd/gubrak