
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := build

GORUN=go run
GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get

CONTAINER=golang
DOCKER_EXEC=docker exec
DOCKER_COMPOSE=docker compose

.PHONY: clean
clean: ## clean build files.
	@sudo rm ./build/*

.PHONY: build
build: ## build go files.
	@$(DOCKER_EXEC) $(CONTAINER) bash -c "$(GOBUILD) -o ./build/basic-go-shell ./src"

.PHONY: run
run: ## run go files.
	@$(DOCKER_EXEC) $(CONTAINER) bash -c "$(GORUN) ./src/*.go"

.PHONY: setup
setup: ## containers start.
	@$(DOCKER_COMPOSE) up -d

.PHONY: 
down: ## remove docker containers.
	@$(DOCKER_COMPOSE) down

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*#"} /^[a-zA-Z_-]+:.*?#/ {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

%:
	@echo 'command "$@" is not found.'
	@$(MAKE) help
	@exit 2