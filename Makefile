export PROJECT_ROOT=$(shell pwd)

APP_NAME=impulse
CONFIG_PATH=config.json
EVENTS_PATH=events

run:
	@go run . $(CONFIG_PATH) $(EVENTS_PATH)

build:
	@mkdir -p bin
	@go build -o bin/$(APP_NAME) .

start: build
	@./bin/$(APP_NAME) $(CONFIG_PATH) $(EVENTS_PATH)

test:
	@go test ./...

test-v:
	@go test -v ./...

clean:
	@rm -rf bin