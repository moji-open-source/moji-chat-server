package router

import (
	"github.com/moji-open-source/moji-chat-server/setup"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func Setup(env *setup.Env, db *gorm.DB, gRPCSvc *grpc.Server) {
	UseUserRouter(env, db, gRPCSvc)
}
