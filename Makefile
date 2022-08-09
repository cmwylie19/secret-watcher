# Makefile for building the secret-watcher server + docker image.

.DEFAULT_GOAL := docker-image

IMAGE ?= cmwylie19/secret-watcher:latest

image/secret-watcher: $(shell find . -name '*.go')
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o $@ ./cmd/secret-watcher

.PHONY: docker-image
docker-image: image/secret-watcher
	docker build -t $(IMAGE) image/

.PHONY: push-image
push-image: docker-image
	docker push $(IMAGE)