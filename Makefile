
# See `make help` for a list of all available commands.
# Configuration
PROJECT_NAME ?= auth-service
LOG_LEVEL = INFO
BUILD_TIMESTAMP := $(shell date +%Y-%m-%d-%H-%M-%S)
CI_COMMIT_SHORT_SHA := $(shell git rev-parse --short HEAD)
ENV_FILE ?= .env
.ONESHELL:
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables

# include env file
#-include $(ENV_FILE)

.PHONY: up
up: build start

.PHONY: down
down: stop

.PHONY: build
build:
	@echo "Building backend..."
	@go build -o ./bin/authservice ./cmd/api/*
	@echo "Back end built!"

.PHONY: start
start:
	@echo "Starting backend..."
	@ ./bin/authservice &
	@echo "Back end started!"

.PHONY: stop
stop:
	@echo "stopping backend ..."
	@-pkill -SIGTERM -f "authservice"
	@echo "stopped backend ..."