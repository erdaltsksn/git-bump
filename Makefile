.POSIX:

.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## Get dependencies
	go get -v -t -d ./...

.PHONY: fmt
fmt: ## Run all formatings
	go mod vendor
	go mod tidy
	go fmt ./...

.PHONY: run
run: ## Run all examples
	make fmt
	go run ./cmd/main.go

.PHONY: test
test: ## Run all test
	go test -v ./...

.PHONY: docs
docs: ## Generate documentation
	go run docs/gen.go

.PHONY: build
build: ## Build git-bump app
	go build -o ./bin/git-bump cmd/main.go

.PHONY: clean
clean: ## Clean all generated files
	rm -rf ./bin/
	rm -rf ./vendor/
