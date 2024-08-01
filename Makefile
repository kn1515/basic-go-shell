
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
help: ## display this help screen
    @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

%:
	@echo 'command "$@" is not found.'
	@$(MAKE) help
	@exit 2