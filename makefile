# Makefile

# Variables
GOOSE_DIR = db/postgres/schema/
DB_HOST = localhost
DB_PORT = 5432
DB_USER = admin
DB_PASS = root
DB_NAME = dataflow
DB_SSLMODE = disable

# Default target
.PHONY: migrate
migrate:
	goose -dir $(GOOSE_DIR) postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=$(DB_SSLMODE)" up

# Other possible targets (optional)
.PHONY: migrate-down
migrate-down:
	goose -dir $(GOOSE_DIR) postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=$(DB_SSLMODE)" down

.PHONY: status
status:
	goose -dir $(GOOSE_DIR) postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=$(DB_SSLMODE)" status
