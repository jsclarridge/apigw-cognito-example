.PHONY: help
help: ## Print the help documentation
	@grep -E '^[/a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: envrc_local
envrc_local: ## Initialize .envrc.local file
	cp .envrc.local.template .envrc.local

.PHONY: build
build: ## Build the binary
	go build -o bin/callapi main.go

.PHONY: run
run: build ## Run the app
	./bin/callapi

default: help
