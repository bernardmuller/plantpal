include .env

run-plants-service:
	go run services/plants/main.go

run-web-service:
	go run services/web/main.go

run-all-services:
	@echo "Running all services..."
	@make run-plants-service & make run-web-service

gen-plants:
	@protoc \
		--proto_path=services "services/plants/plantspb/plants.proto" \
		--go_out=services --go_opt=paths=source_relative \
  		--go-grpc_out=services --go-grpc_opt=paths=source_relative

db_migrate_up:
	@echo "Migrating up..."
	@cd ./internal/store/postgres/schema &&	goose postgres ${POSTGRES_URI} up

db_migrate_down:
	@echo "Migrating down..."
	@cd ./internal/store/postgres/schema &&	goose postgres ${POSTGRES_URI} down

db_generate_queries:
	@echo "Generating queries..."
	@sqlc generate