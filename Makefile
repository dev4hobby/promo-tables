path := .

define Comment
	- Run `make help` to see all the available options.
endef


.PHONY: help
help: ## 도움말
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

.PHONY: build
build: ## Build
	go build -v

.PHONY: clean
clean: ## Clean all build resources
	rm -rf promo-tables promo-tables.exe