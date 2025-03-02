migrate_up:
	@echo "Migrating up..."
	migrate -path ./db/migration -database "postgresql://admin:y7jHf&DNWG15@localhost:5030/main?sslmode=disable" --verbose up

migrate_down:
	@echo "Migrating down..."
	migrate -path ./db/migration -database "postgresql://admin:y7jHf&DNWG15@localhost:5030/main?sslmode=disable" --verbose down

.PHONY: migrate_up migrate_down