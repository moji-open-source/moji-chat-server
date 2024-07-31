.PHONY: build-proto, run

build-proto:
	@rm -rf abi && mkdir abi
	@set -ex && protoc --go_out=. --go-grpc_out=. ./proto/**.proto 

dev:
	@docker-compose up redis postgres -d && go run main.go
