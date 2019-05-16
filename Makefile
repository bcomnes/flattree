.PHONY: all build deps generate help test validate
AN_ENV_VAR=foo

help: ## Show this help.
		@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: deps generate test build

build:
	go build ./...

deps:
	GO111MODULE=off go get -u github.com/myitcv/gobin && go mod download

generate:
	go generate

test:
	gobin -m -run github.com/kyoh86/richgo test -v ./...
