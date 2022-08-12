# Makefile for building the secret-watcher server + docker image.
SHELL=bash

# The stage to build, one of: dev, test, prod
# Default to dev if not set.
# dev, test are arm64 only
# prod is amd64 only

ifndef STAGE
	export STAGE :=dev
endif

ifeq ($(STAGE),dev)
	export ARCH :=arm64
else ifeq ($(STAGE),test)
	export ARCH :=arm64
else ifeq ($(STAGE),prod)
	export ARCH :=amd64
else
	$(error STAGE must be one of: dev, test, prod)
endif


IMAGE ?= cmwylie19/secret-watcher:${ARCH}

#---------------------------
# Build the secret-watcher binary

.PHONY: build/secret-watcher
build/secret-watcher: $(shell find . -name '*.go')
	GOARCH=${ARCH} CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o $@ ./cmd/secret-watcher


#---------------------------
# Build the docker image

.PHONY: build/image
build/image: build/secret-watcher
	docker build -t $(IMAGE) build/
	rm build/secret-watcher

#--------------------------------
# Push the docker image to dockerhub
.PHONY: push-image
push-image: build/image
	docker push $(IMAGE)

#--------------------------------
all: build/secret-watcher build/image push-image
	@echo "Done building secret-watcher for ${ARCH}"