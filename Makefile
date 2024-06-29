BIN := main
BIN_DIR := ./bin
GOPATH := $(GOPATH)

.PHONY: help
help: ## Display usage information
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Clean the project
	@echo "Cleaning the project..."
	@rm -rf $(BIN_DIR)

.PHONY: build
build: ## Build the project
	@go build -o $(BIN_DIR)/$(BIN) .

.PHONY: run
run: build ## Run the project
	@$(BIN_DIR)/$(BIN)
