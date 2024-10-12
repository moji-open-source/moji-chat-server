package router

import (
	"github.com/moji-open-source/moji-chat-server/abi/grpc_user"
	"github.com/moji-open-source/moji-chat-server/api/controller"
	"github.com/moji-open-source/moji-chat-server/domain/models"
	"github.com/moji-open-source/moji-chat-server/redisson"
	"github.com/moji-open-source/moji-chat-server/services"
)

func UseUserRouter(ctx RouterContext) {
	signin := &services.SigninService{
		UserRepository: models.NewUserRepository(ctx.Database),
		Redisson:       redisson.Redisson{Client: ctx.Redis},
	}
	grpc_user.RegisterUserServiceServer(ctx.Server, &controller.UserController{Env: ctx.Env, SigninService: signin})
}
