.PHONY: help
help: ## Print the help documentation
	@grep -E '^[/a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: envrc_local
envrc_local: ## Initialize .envrc.local file
	@cp .envrc.local.template .envrc.local

pre-commit-install: ## Install pre-commit
	pip install pre-commit

lint: ## Run lint check
	pre-commit run --all-files

.PHONY: build
build: ## Build the binary
	@go build -o bin/callapi main.go

.PHONY: run
run: ## Run the app
	@./bin/callapi

default: help
