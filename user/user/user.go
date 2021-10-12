package user

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	"pasteProject/errs"
	"pasteProject/repositories"
	userpb "pasteProject/user/api/gen/go"
)

var (
	user       repositories.IUserRepo
	errHandler errs.IErr
)

func Init(db *gorm.DB) {
	user = repositories.NewUserRepo(db)
	errHandler = errs.NewErrs(db)
}

func init() {
	Init(errs.GetDB())
}

type Service struct {
	userpb.UnimplementedUserServiceServer
}

func (s *Service) Login(ctx context.Context, req *userpb.UserRequest) (resp *userpb.UserResponse, err error) {
	errs.Refresh()
	Init(errs.GetDB())
	return &userpb.UserResponse{
		StatusCode: http.StatusOK,
		Token:      "This is a token " + req.Username,
	}, nil
}
