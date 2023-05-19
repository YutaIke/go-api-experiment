.PHONY: build build-local up down test lint
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

lint: install-golangci-lint install-reviewdog
	golangci-lint run --config .golangci.yml | reviewdog -f=golangci-lint -diff "git diff HEAD^"   

install-golangci-lint: ## Install golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

install-reviewdog: ## Install reviewdog
	go install github.com/reviewdog/reviewdog/cmd/reviewdog@latest
