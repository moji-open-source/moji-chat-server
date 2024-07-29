.PHONY: build-proto, run

build-proto:
	@rm -rf abi && mkdir abi
	@set -ex && protoc --go_out=. --go-grpc_out=. ./proto/**.proto 

run:
	@go run main.go
