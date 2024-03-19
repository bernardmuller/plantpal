# PID = /tmp/domain-app.pid
# GO_FILES = $(wildcard *.go)
# APP = ./domain-app
include .env

# serve: start
# 	@fswatch -x -o --event Created --event Updated --event Renamed -r -e '.*' -i '\.go$$' . | while read; do make restart || make kill; done

kill_port: 
	@echo "Killing process on port 8080..."
	@lsof -ti :8080 | xargs -r kill -9 || true

# before:
# 	@echo "actually do nothing"

# build: $(GO_FILES)
# 	@go build -o $(APP)

# $(APP): $(GO_FILES)
# 	@go build $? -o $@

# start: build
# 	@echo "Starting $(APP)..."
# 	@echo "Generating templates..."
# 	@templ generate

# 	# @sh -c "$(APP) & echo $$! > $(PID)"
# 	@$(APP) & echo $$! > $(PID)

# restart: kill_port before build start

# .PHONY: start serve restart kill_port before  # let's go to reserve rules names
build:
	@echo "Building server..."
	@go build -o $(APP)

run: kill_port
	@echo "Generating templates..."
	@templ generate
	@echo "Running main.go..."
	@TURSO_AUTH_TOKEN=${TURSO_AUTH_TOKEN} TEST_DATABASE_URL=${TEST_DATABASE_URL} MONGODB_URI=${MONGODB_URI} go run main.go