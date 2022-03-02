
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

.PHONY: build
build:
	@echo "Building back end..."
	@go build -o tmp/main ./cmd/api/*
	@echo "Back end built!"

.PHONY: start
start:
	@echo "Starting back end..."
	@ ./tmp/main &
	@echo "Back end started!"

.PHONY: down
down:
	@echo "stopping backend ..."
	@-pkill -SIGTERM -f "main"
	@echo "stopped backend ..."