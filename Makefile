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
	docker run --name $(APPNAME) --net=host -e POSTGRES_DSN='host=localhost user=postgres dbname=rozetka' --rm $(USERNAME)/$(APPNAME):$(VERSION)

docker_github_run:
	docker run --name $(APPNAME) --net=host -e API_DEBUG='false' -e POSTGRES_DSN='host=localhost user=postgres dbname=rozetka' --rm ghcr.io/$(USERNAME)/$(APPNAME):latest