SHELL = /bin/bash
OS = $(shell uname -s | tr '[:upper:]' '[:lower:]')

# Build variables
API_BINARY_NAME = robots-api
SIMULATOR_BINARY_NAME = robots-simulator
BUILD_DIR ?= bin

# Go variables
export CGO_ENABLED ?= 0
export GOOS ?= $(OS)
export GOARCH ?= amd64
GOFILES = $(shell find . -type f -name '*.go' -not -path "*/mock/*.go" -not -path "*.pb.go")

.PHONY: all
all: dep generate build-api build-simulator #install ## Runs dep generate build-api build-simulator

.PHONY: clean
clean: ## Clean the working area and the project
	@rm -rf $(BUILD_DIR)/

.PHONY: dep
dep: ## Install dependencies
	@go install github.com/golang/mock/mockgen@v1.6.0
	@go get -u -d github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	@go mod tidy
	@go get -v -t ./...

.PHONY: generate
generate:
	@go generate ./...

.PHONY: build-api
build-api: GOARGS += -o $(BUILD_DIR)/$(API_BINARY_NAME) ## Build API
build-api:
	@go build -v $(GOARGS) ./api/main.go

.PHONY: build-simulator
build-simulator: GOARGS += -o $(BUILD_DIR)/$(SIMULATOR_BINARY_NAME) ## Build Simulator
build-simulator:
	@go build -v $(GOARGS) ./simulator/main.go

.PHONY: test
test: ## Run unit tests
	@cd api && go test  -covermode=count ./...
	@cd simulator && go test  -covermode=count ./...
	@cd shared && go test  -covermode=count ./...

.PHONY: install
install: ## Install the binaries to /usr/local/bin
	@sudo cp $(BUILD_DIR)/$(API_BINARY_NAME) /usr/local/bin
	@sudo cp $(BUILD_DIR)/$(SIMULATOR_BINARY_NAME) /usr/local/bin

.PHONY: lint
lint: ## run golanci-lint locally
	@docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v

.PHONY: format
format: ## Format the source
	@goimports -w $(GOFILES)

.PHONY: list
list: ## List all make targets
	@$(MAKE) -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
.DEFAULT_GOAL := help
help: ## Get help output
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
