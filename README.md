# moji-chat-server

## Before Start

First of all, You need to install [Protocol Buffer](https://github.com/protocolbuffers/protobuf/releases) complier.

You need to add the bin directory to the environment on the after download.

You can use brew to install protubuf, If you use MacOs/Linux

```shell
brew install protobuf
```

Finally, Install the following plugins all is done!

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
