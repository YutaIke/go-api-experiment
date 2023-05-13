.PHONY: build build-local up down

DOCKDR_TAG := latest
build: ## Build docker image to deploy
	docker build -t go-api-experiment:${DOCKDR_TAG}

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down
