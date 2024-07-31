package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/moji-open-source/moji-chat-server/domain/models"
	"github.com/moji-open-source/moji-chat-server/redisson"
)

type SigninService struct {
	models.UserRepository
	redisson.Redisson
}

func (s *SigninService) GetUserByEmail(email string) (models.SysUserModel, error) {
	return s.UserRepository.GetByEmail(email)
}

func (s *SigninService) CreateAccessToken(user *models.SysUserModel) (string, error) {
	userIdAlias := uuid.New().String()

	ctx := context.Background()
	key := redisson.AppendRedisKey(redisson.Authentication, fmt.Sprint(user.Uid))
	err := s.Redisson.Set(ctx, key, userIdAlias, time.Duration(12)*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return userIdAlias, nil
}
