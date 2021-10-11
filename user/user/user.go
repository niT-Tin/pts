package user

import (
	"context"
	"net/http"
	userpb "pasteProject/user/api/gen/go"
)

type Service struct {
	userpb.UnimplementedUserServiceServer
}

func (s *Service) Login(ctx context.Context, req *userpb.UserRequest) (resp *userpb.UserResponse, err error) {
	return &userpb.UserResponse{
		StatusCode: http.StatusOK,
		Token:      "This is a token " + req.Username,
	}, nil
}
