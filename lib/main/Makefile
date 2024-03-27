include .env

kill_port:
	@echo "Killing process on port 8080..."
	@lsof -ti :8080 | xargs -r kill -9 || true

build:
	@echo "Building server..."
	@go build -o $(APP)

run:
	@echo "Running main.go..."
	@TURSO_AUTH_TOKEN=${TURSO_AUTH_TOKEN} TEST_DATABASE_URL=${TEST_DATABASE_URL} MONGODB_URI=${MONGODB_URI} go run .

db_migrate_up:
	@echo "Migrating up..."
	@cd ./internal/store/postgres/schema &&	goose postgres ${POSTGRES_URI} up

db_migrate_down:
	@echo "Migrating down..."
	@cd ./internal/store/postgres/schema &&	goose postgres ${POSTGRES_URI} down

db_generate_queries:
	@echo "Generating queries..."
	@sqlc generate