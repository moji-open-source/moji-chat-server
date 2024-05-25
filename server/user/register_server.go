package user

import (
	"github.com/moji-open-source/moji-chat-server/abi/user"
	"google.golang.org/grpc"
)

func RegisterServer(server *grpc.Server) {
	user.RegisterUserServer(server, &GrpcUserServer{})
}
