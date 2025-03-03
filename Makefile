
include .env
export $(shell sed 's/=.*//' .env)

migrate_up:
	@echo "Migrating up..."
	migrate -path ./db/migration -database "$(DATABASE_URL)" --verbose up

migrate_down:
	@echo "Migrating down..."
	migrate -path ./db/migration -database "$(DATABASE_URL)" --verbose down

.PHONY: migrate_up migrate_down