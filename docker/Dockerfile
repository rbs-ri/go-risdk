FROM golang:1.18-bullseye

RUN apt-get update && \
    apt-get install -y --no-install-recommends unzip wget ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Создаём ссылку на legacy protoc-gen-go: он генерирует gRPC в тот же *.pb.go.
RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.2 && \
    ln -sf /go/bin/protoc-gen-go /usr/local/bin/protoc-gen-go

ARG PLATFORM
ARG PROTOC_VERSION

# PROTOC ставим для архитектуры x86_64
RUN export ZIP=x86_64 && \
    if [ "${PLATFORM}" = "arm64" ] || [ "${PLATFORM}" = "aarch64" ]; then export ZIP=aarch_64; fi && \
    wget --quiet https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-${ZIP}.zip && \
    unzip -o protoc-${PROTOC_VERSION}-linux-${ZIP}.zip -d /usr/local bin/protoc && \
    unzip -o protoc-${PROTOC_VERSION}-linux-${ZIP}.zip -d /usr/local 'include/*'
