package router

import (
	"github.com/moji-open-source/moji-chat-server/setup"
	"google.golang.org/grpc"
)

type RouterContext struct {
	setup.Application
	Server *grpc.Server
}

func Setup(ctx RouterContext) {
	UseUserRouter(ctx)
}
