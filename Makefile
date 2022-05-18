path := .

define Comment
	- Run `make help` to see all the available options.
endef


.PHONY: help
help: ## What you see here..
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## Init tables and insert mockdata
	go run main.go -init

.PHONY: flush
flush: ## Drop all tables
	go run main.go -flush

.PHONY: test
test: ## For CI and Local test
	go test ./test -v

.PHONY: test-clean
test-clean: ## Expires all test results
	go clean -testcache

.PHONY: build-docker
build-docker: ## Build docker image
	docker-compose up --build

.PHONY: build-go
build-go: ## Build
	go build -v

.PHONY: clean
clean: ## Clean all build resources
	rm -rf promo-tables promo-tables.exe
	go clean -testcache