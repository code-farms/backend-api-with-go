# Build the Go application and generate the binary in the bin directory
build:
	@go build -o bin/backend-api cmd/main.go

# Run the tests in all subdirectories
test:
	@go test -v ./...

# Run the compiled application after building
run: build
	@./bin/backend-api

# Declare "migration" as a phony target to avoid conflicts with files or directories named "migration".
.PHONY: migration migrate-up migrate-down

# Create a new database migration.
# Usage: `make migration <migration-name>`
# Example: `make migration add-user-table`
migration:
	@echo "Creating migration: $(filter-out $@,$(MAKECMDGOALS))"
	# This command creates migration files (up and down) in the specified directory.
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

# Catch-all rule to avoid errors when passing arguments (e.g., migration names) to the Makefile.
%:
	@:

# Apply all pending migrations to the database.
# Usage: `make migrate-up`
migrate-up:
	# This runs the custom Go migration script located at cmd/migrate/main.go with the "up" command.
	@go run cmd/migrate/main.go up

# Rollback the most recent migration.
# Usage: `make migrate-down`
migrate-down:
	# This runs the custom Go migration script located at cmd/migrate/main.go with the "down" command.
	@go run cmd/migrate/main.go down
