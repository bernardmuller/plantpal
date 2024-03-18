PID = /tmp/domain-app.pid
GO_FILES = $(wildcard *.go)
APP = ./domain-app


serve: kill_port start
	@fswatch -x -o --event Created --event Updated --event Renamed -r -e '.*' -i '\.go$$' -i '\.templ$$' -i '\.js$$' -i '\.ts$$'  . | xargs -n1 -I{}  make restart || make kill_port

kill_port:
	@echo "Killing process on port 8080..."
	@lsof -ti :8080 | xargs -r kill -9 || true

before:
	@echo "actually do nothing"

build: $(GO_FILES)
	@go build -o $(APP)

$(APP): $(GO_FILES)
	@go build $? -o $@

start:
	@echo "Starting $(APP)..."
	@echo "Generating templates..."
	@templ generate

	# @sh -c "$(APP) & echo $$! > $(PID)"
	@$(APP) & echo $$! > $(PID)

restart: kill_port before build start

.PHONY: start serve restart kill_port before  # let's go to reserve rules names
