@REM Генерация gRPC микросервиса на основе RoboSdk.proto

set GOROOT=C:\go-1.17.1_amd64
set GOARCH=386
set PATH=C:\Mingw64\lib;C:\Mingw64\bin;%GOROOT%\bin

protoc -I proto/ proto/RoboSdk.proto --go_out=plugins=grpc:.
