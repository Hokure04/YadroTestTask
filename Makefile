export PROJECT_ROOT=$(shell pwd)

APP_NAME=impulse
CMD_PATH=./cmd
CONFIG_PATH=config.json
EVENTS=events

run:
	@go run $(CMD_PATH) $(CONFIG_PATH) $(EVENT_PATH)

build:
	@mkdir -p bin
	@go build -o bin/$(APP_NAME) $(CMD_PATH)

start: build
	@./bin/$(APP_NAME) $(CONFIG_PATH) $(EVENTS_PATH)

clean:
	@rm -rf bin