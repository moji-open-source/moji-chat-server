package services

import "github.com/moji-open-source/moji-chat-server/domain/models"

type SigninService struct {
	models.UserRepository
}

func (s *SigninService) GetUserByEmail(email string) (models.SysUserModel, error) {
	return s.UserRepository.GetByEmail(email)
}
