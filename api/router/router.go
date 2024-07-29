package router

import "google.golang.org/grpc"

func Setup(gRPCSvc *grpc.Server) {
	UseUserRouter(gRPCSvc)
}
