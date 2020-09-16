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
run: ## Run the application
	make fmt
	go run ./cmd/main.go

.PHONY: test
test: ## Run all test
	go test -v ./...

.PHONY: coverage
coverage: ## Show test coverage
	@go test -coverprofile=coverage.out ./... > /dev/null
	go tool cover -func=coverage.out
	rm coverage.out

.PHONY: docs
docs: ## Generate the documentation
	go run docs/gen.go

.PHONY: godoc
godoc: ## Start local godoc server
	@echo "See Documentation:"
	@echo "\thttp://localhost:6060/pkg/github.com/erdaltsksn/git-bump"
	@echo "\n"
	@godoc -http=:6060

.PHONY: build
build: ## Build the app
	go build -o ./bin/git-bump cmd/main.go

.PHONY: clean
clean: ## Clean all generated files
	rm -rf ./bin/
	rm -rf ./vendor/
