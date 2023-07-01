# Makefile

## HELP:
.PHONY: help
## help: Show this help message.
help:
	@echo "Usage: make [target]\n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## :
## BUILD:

.PHONY: build
## build: Build Go server code.
build:
	go build -o /dev/null ./...

## :
## DEPENDENCIES:

.PHONY: dep-clean
## dep-clean: Clean up dependency files.
dep-clean:
	@rm go.mod go.sum

.PHONY: dep-get
## dep-get: Get Go modules.
dep-get:
	go mod tidy

.PHONY: dep-init
## dep-init: Initialize Go modules.
dep-init:
	go mod init

.PHONY: dep-update
## dep-update: Update Go modules.
dep-update:
	go get -u ./...
	go mod tidy

## :
## RUN:

.PHONY: run
## run: Run logo2go.
run:
	go run ./cmd/logo2go/...

## :
## TEST:

.PHONY: test
## test: Run Go tests.
test:
	go test -coverprofile=cover.out ./cmd/logo2go/...

.PHONY: test-clean
## test-clean: Delete test output files.
test-clean:
	@rm cover.out

.PHONY: test-coverage
## test-coverage: View test coverage report.
test-coverage:
	go tool cover -html=cover.out

## :
