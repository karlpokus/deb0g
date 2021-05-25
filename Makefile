VERSION := $(shell git describe --always --dirty --tags)
GOOS    := $(shell go env GOOS)
GOARCH  := $(shell go env GOARCH)

.PHONY: build docker push

build:
	CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=$(VERSION)" -o bin/server

docker: build
	docker build -t pokus2000/deb0g:$(VERSION) .

push: build docker
	docker push pokus2000/deb0g:$(VERSION)
