# Variables
GOOSE_DIR = db/postgres/schema/
DB_HOST = localhost
DB_PORT = 5432
DB_USER = admin
DB_PASS = root
DB_NAME = dataflow
DB_SSLMODE = disable
GO_FILE = main.go

# Check and install goose
.PHONY: ensure-goose
ensure-goose:
	@if ! [ -x "$$(command -v goose)" ]; then \
		echo "goose is not installed. Installing..."; \
		go install github.com/pressly/goose/v3/cmd/goose@latest; \
	else \
		echo "goose is already installed."; \
	fi

# Default target
.PHONY: migrate
migrate: ensure-goose
	goose -dir $(GOOSE_DIR) postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=$(DB_SSLMODE)" up

# Other possible targets
.PHONY: migrate-down
migrate-down: ensure-goose
	goose -dir $(GOOSE_DIR) postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=$(DB_SSLMODE)" down

.PHONY: status
status: ensure-goose
	goose -dir $(GOOSE_DIR) postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=$(DB_SSLMODE)" status

# Run the main.go file
.PHONY: run
run:
	go run $(GO_FILE)