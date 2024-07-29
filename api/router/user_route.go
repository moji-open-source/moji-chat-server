package router

import (
	"github.com/moji-open-source/moji-chat-server/abi/grpc_user"
	"github.com/moji-open-source/moji-chat-server/api/controller"
	"google.golang.org/grpc"
)

func UseUserRouter(svc *grpc.Server) {
	grpc_user.RegisterUserServer(svc, &controller.UserController{})
}
