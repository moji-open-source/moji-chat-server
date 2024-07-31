package domain

import "github.com/moji-open-source/moji-chat-server/domain/models"

type SigninService interface {
	GetUserByEmail(email string) (models.SysUserModel, error)
	CreateAccessToken(user *models.SysUserModel) (accessToken string, err error)
}
