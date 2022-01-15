USERNAME := $(shell git config user.name)
APPNAME := $(shell basename `git rev-parse --show-toplevel`)
BUILDTIMESTAMP := $(shell date -u "+%Y-%m-%d %H:%M:%S")
GOVERSION := $(shell go version | sed -r 's/go version go(.*)\ .*/\1/')
VERSION := $(shell git describe --tags)

demo:
	@echo "hello" $(VERSION)

docker_build:
	docker build -t $(USERNAME)/$(APPNAME):$(VERSION) -f ./Dockerfile .

docker_run:
	docker run --rm --name $(APPNAME) --net=host -e POSTGRES_DSN=$(POSTGRES_DSN) \
		-e TELEGRAM_TOKEN=$(TELEGRAM_TOKEN) -e LOG_LEVEL=debug $(USERNAME)/$(APPNAME):$(VERSION)

docker_push:
	docker push $(USERNAME)/$(APPNAME):$(VERSION)

docker_github_run:
	docker run --name $(APPNAME) --net=host -e API_DEBUG='false' --rm ghcr.io/$(USERNAME)/$(APPNAME):latest

build_api:
	java -jar ~/.local/bin/openapi-generator-cli.jar generate -g go -i openapi/openapi-manual.yaml -o api \
      --type-mappings integer=int64 --type-mappings number=float64 --additional-properties=hideGenerationTimestamp=true,packageName=api

