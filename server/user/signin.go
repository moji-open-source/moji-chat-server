package user

import (
	"context"

	"github.com/moji-open-source/moji-chat-server/abi/user"
)

func (GrpcUserServer) Signin(
	ctx context.Context, req *user.UserSigninReq,
) (*user.UserSigninResp, error) {
	return &user.UserSigninResp{}, nil
}
