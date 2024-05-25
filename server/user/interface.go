package user

import "github.com/moji-open-source/moji-chat-server/abi/user"

type GrpcUserServer struct {
	user.UnimplementedUserServer
}
