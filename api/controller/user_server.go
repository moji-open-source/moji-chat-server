package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/moji-open-source/moji-chat-server/abi/grpc_user"
	"github.com/moji-open-source/moji-chat-server/domain"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

type UserController struct {
	domain.SigninService
	grpc_user.UnimplementedUserServer
}

func (c *UserController) Signin(ctx context.Context, request *grpc_user.UserSigninRequest) (*grpc_user.UserSigninResponse, error) {
	email := strings.TrimSpace(request.GetEmail())
	password := request.GetPassword()

	user, err := c.SigninService.GetUserByEmail(email)
	if err != nil {
		return nil, status.Errorf(http.StatusNotFound, "User not found with the gevin email")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, status.Errorf(http.StatusInternalServerError, "Invalid credentials")
	}

	return &grpc_user.UserSigninResponse{AccessToken: "token"}, nil
}
