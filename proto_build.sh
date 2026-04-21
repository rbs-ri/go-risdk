#!/usr/bin/env bash

# Генерация gRPC микросервиса на основе RoboSdk.proto через Docker
set -euo pipefail

PROTOC_VERSION="${PROTOC_VERSION:-3.6.1}"
PLATFORM="${PLATFORM:-$(uname -m)}"

PROTOC_VERSION="$PROTOC_VERSION" PLATFORM="$PLATFORM" \
	docker compose -f docker/docker-compose.yaml run --build --rm protogen
