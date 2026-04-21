@REM Генерация gRPC микросервиса на основе RoboSdk.proto через Docker

if "%PROTOC_VERSION%"=="" set PROTOC_VERSION=3.6.1
if "%PLATFORM%"=="" set PLATFORM=x86_64

docker compose -f docker/docker-compose.yaml run --build --rm protogen
