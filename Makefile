.PHONY: test
test: ## test package
	@go test  -bench -v -run '' -coverprofile=fmtcoverage.html github.com/finfini/sdk/...

.PHONY: show-coverage
show-coverage: ## show test coverage in browser
	@go tool cover -html=fmtcoverage.html

.PHONY: lint
lint:  ## Lint this codebase
	@go mod tidy
	@goimports -v -w .
	@gofmt -e -s -w .
	@golint .

.PHONY: help
.DEFAULT_GOAL := help
help:
	@echo  "[!] Available Command: "
	@echo  "-----------------------"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

