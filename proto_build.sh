#!/usr/bin/env bash

# Генерация gRPC микросервиса на основе RoboSdk.proto
export GOROOT=/usr/local/go-1.17.1_amd64
export GOARCH=386
export PATH=$GOROOT/bin:$PATH

protoc --version
protoc -I proto/ proto/RoboSdk.proto --go_out=plugins=grpc:.