package router

import (
	"github.com/moji-open-source/moji-chat-server/abi/grpc_user"
	"github.com/moji-open-source/moji-chat-server/api/controller"
	"github.com/moji-open-source/moji-chat-server/domain/models"
	"github.com/moji-open-source/moji-chat-server/services"
	"github.com/moji-open-source/moji-chat-server/setup"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func UseUserRouter(env *setup.Env, db *gorm.DB, svc *grpc.Server) {
	grpc_user.RegisterUserServer(svc, &controller.UserController{Env: env, SigninService: &services.SigninService{
		UserRepository: models.NewUserRepository(db),
	}})
}
