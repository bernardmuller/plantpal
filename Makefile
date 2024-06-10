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

# Live Reload
#watch:
#	@if command -v air > /dev/null; then \
#	    air; \
#	    echo "Watching...";\
#	else \
#	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
#	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
#	        go install github.com/cosmtrek/air@latest; \
#	        air; \
#	        echo "Watching...";\
#	    else \
#	        echo "You chose not to install air. Exiting..."; \
#	        exit 1; \
#	    fi; \
#	fi