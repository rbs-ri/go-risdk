PROTOC_VERSION ?= 3.6.1

.PHONY: proto-compile
proto-compile:
	PLATFORM=$(shell uname -m) PROTOC_VERSION=$(PROTOC_VERSION) docker compose -f docker/docker-compose.yaml run --build --rm protogen
