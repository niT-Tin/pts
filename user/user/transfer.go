package user

import (
	"net/http"
	"pasteProject/models"
	userpb "pasteProject/user/api/gen/go"
)

func P2MUser(u *userpb.UserRequest) *models.User {
	return &models.User{
		UserName:     u.Username,
		EncryptedPwd: u.Password,
	}
}

func M2PUser(u *models.User) *userpb.UserResponse {
	return &userpb.UserResponse{
		StatusCode: http.StatusOK,
	}
}
