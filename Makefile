export PROJECT_ROOT=$(shell pwd)

APP_NAME=impulse
CMD_PATH=./cmd
CONFIG_PATH=config.json
EVENTS_PATH=events

run:
	@go run $(CMD_PATH) $(CONFIG_PATH) $(EVENTS_PATH)

build:
	@mkdir -p bin
	@go build -o bin/$(APP_NAME) $(CMD_PATH)

start: build
	@./bin/$(APP_NAME) $(CONFIG_PATH) $(EVENTS_PATH)

test:
	@go test ./...

test-v:
	@go test -v ./...

clean:
	@rm -rf bin