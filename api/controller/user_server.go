package controller

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/moji-open-source/moji-chat-server/abi/grpc_user"
	"github.com/moji-open-source/moji-chat-server/domain"
	"github.com/moji-open-source/moji-chat-server/setup"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserController struct {
	domain.SigninService
	grpc_user.UnimplementedUserServiceServer
	*setup.Env
}

func (c *UserController) Signin(
	ctx context.Context,
	request *grpc_user.SigninRequest,
) (*grpc_user.SigninResponse, error) {
	email := strings.TrimSpace(request.GetEmail())
	password := request.GetPassword()

	user, err := c.SigninService.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(http.StatusNotFound, "User not found with the gevin email")
		}
		return nil, status.Errorf(http.StatusInternalServerError, "unkown system error")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, status.Errorf(http.StatusInternalServerError, "Invalid credentials")
	}

	accessToken, err := c.SigninService.CreateAccessToken(&user)
	if err != nil {
		return nil, status.Errorf(http.StatusInternalServerError, err.Error())
	}

	return &grpc_user.SigninResponse{AccessToken: accessToken}, nil
}
