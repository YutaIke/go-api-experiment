.PHONY: build build-local up down
GO := go1.20.1

DOCKDR_TAG := latest
build: ## Build docker image to deploy
	docker build -t go-api-experiment:${DOCKDR_TAG}

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up

down: ## Do docker compose down
	docker compose down

test: ## Run test
	$(GO) test -race -shuffle=on ./...
