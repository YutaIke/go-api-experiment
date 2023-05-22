.PHONY: build build-local up-all up-infra up-app down-all down-infra down-app test lint 
GO := go1.20.1
DATA_SOURCE_URL := mysql://root:root@localhost:13306/go-api-experiment-local?charset=utf8&parseTime=True&loc=Local
MIGRATION_FILE_PATH := file://ent/migrate/migrations

DOCKDR_TAG := latest
build: ## Build docker image to deploy
	docker build -t go-api-experiment:${DOCKDR_TAG}

build-local: ## Build docker image to local development
	docker compose build --no-cache

up-all: up-infra up-app

up-infra: ## Do docker compose up db and redis
	docker compose up db -d
	docker compose up redis -d

up-app: ## Do docker compose up app with hot reload
	docker compose up app

down-all: down-infra down-app

down-infra: ## Do docker compose down db and redis
	docker compose rm -fsv db
	docker compose rm -fsv redis

down-app: ## Do docker compose down app
	docker compose rm -fsv app

test: ## Run test
	$(GO) test -race -shuffle=on ./...

lint: install-golangci-lint install-reviewdog
	golangci-lint run --config .golangci.yml | reviewdog -f=golangci-lint -diff "git diff HEAD^"   

install-golangci-lint: ## Install golangci-lint
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

install-reviewdog: ## Install reviewdog
	$(GO) install github.com/reviewdog/reviewdog/cmd/reviewdog@latest

clean-db-volume: down-infra
	docker volume rm go-api-experiment_db_data

## ent
.PHONY: ent-generate ent-generate-migration ent-migrate-apply-local ent-migrate-down-local ent-migrate-status
ent-generate: ## Generate ent code
	$(GO) generate ./ent

ent-generate-migration: ## Migrate ent schema
	$(GO) run -mod=mod ./cmd/migration/main.go test ## TODO: change test to use input args

ent-migrate-apply-local: ## Apply ent migration to local db
	 atlas migrate apply 
	 	--dir $(MIGRATION_FILE_PATH) \
		--url $(DATA_SOURCE_URL)

ent-migrate-down-local: ## Down ent migration to local db
	atlas schema apply \
		--to $(MIGRATION_FILE_PATH) \
		--url $(DATA_SOURCE_URL) \
		--dev-url "docker://mysql/8/go-api-experiment-local" \
		--exclude "atlas_schema_revisions"
	atlas migrate set \
		--dir $(MIGRATION_FILE_PATH) \
		--url $(DATA_SOURCE_URL)

ent-migrate-status: ## Show ent migration status
	atlas migrate status \
		--dir $(MIGRATION_FILE_PATH) \
		--url $(DATA_SOURCE_URL)