# Golang =====================================

GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
BINARY_NAME=auth-server
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build ## go test & go build
.PHONY: build
build: ## build go binary
	$(GOBUILD) -o $(BINARY_NAME) -v
.PHONY: test
test: ## go test
	$(GOTEST) -v ./...
.PHONY: clean
clean: ## remove go bainary
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
.PHONY: run
run: ## run go bainary
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
.PHONY: doc
doc: ## exec godoc localhost:6060
	godoc -http=:6060

# MySQL ======================================

.PHONY: initdb
initdb: ## create database and insert first data
	mysql -p -u root blogo < init/init.sql

# Docker =====================================

.PHONY: up
up: ## docker-compose up
	docker-compose up -d

.PHONY: down
down: ## docker-compose down
	docker-compose down

# Help =====================================

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

