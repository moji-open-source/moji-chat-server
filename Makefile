export INCLUDE_DIR=/usr/local/Cellar/protobuf/26.1/include:.

.PHONY: build-proto, run

build-proto:
	@rm -rf abi && mkdir abi
	@set -ex && protoc -I=${INCLUDE_DIR} --go_out=. --go-grpc_out=. ./proto/**/*.proto 

run:
	@go run main.go
