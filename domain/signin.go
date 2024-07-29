package domain

import "github.com/moji-open-source/moji-chat-server/domain/models"

type SigninService interface {
	GetUserByEmail(email string) (models.SysUserModel, error)
}
