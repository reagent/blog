.DEFAULT_GOAL := run

DB_USER ?= $(shell whoami)
DB_HOST ?= localhost
DB_NAME ?= blog
DB_PORT ?= 5432

MIGRATION_PATH=./db/migrations

.PHONY: create-db
create-db:
	createdb -h $(DB_HOST) -U $(DB_USER) -p $(DB_PORT) $(DB_NAME)

.PHONY: drop-db
drop-db:
	dropdb -h $(DB_HOST) -U $(DB_USER) -p $(DB_PORT) $(DB_NAME)

.PHONY: migrate
migrate:
	migrate \
		-url postgres://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable \
		-path $(MIGRATION_PATH) \
		up

.PHONY: rollback
rollback:
	migrate \
		-url postgres://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable \
		-path $(MIGRATION_PATH) \
		migrate -1

.PHONY: create-migration
create-migration:
	@test '$(name)' != ''
	@migrate \
		-url postgres://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable \
		-path $(MIGRATION_PATH) \
		create $(name)

.PHONY: build
build:
	go build -o ./blog main.go

.PHONY: run
run: build
	DB_HOST=$(DB_HOST) DB_USER=$(DB_USER) DB_NAME=$(DB_NAME) ./blog
