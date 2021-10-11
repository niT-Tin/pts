package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pasteProject/errs"
	"pasteProject/models"
	"time"
)

type IUserRepo interface {
	CreateU(*models.User) error
	DeleteU(*models.User) error
	UpdateU(*models.User) error
	SelectU(string) (*models.User, error)
}

type UserRepo struct {
	Db *gorm.DB
}

// NewUserRepo New User Repo
func NewUserRepo(Db *gorm.DB) IUserRepo {
	return &UserRepo{Db: Db}
}

func (u *UserRepo) CreateU(user *models.User) error {
	return NoRepeat("create user failed", "user_repo", u.Db.Create(user))
}

func (u *UserRepo) DeleteU(user *models.User) error {
	return NoRepeat("delete user failed", "user_repo", u.Db.Where("user_name = ?", user.UserName).Delete(&models.User{}))
}

func (u *UserRepo) SelectU(name string) (*models.User, error) {
	var user models.User
	e := NoRepeat("select user failed", "user_repo", u.Db.Where("user_name = ?", name).First(&user))
	return &user, e
}

func (u *UserRepo) UpdateU(user *models.User) error {
	return NoRepeat("update user failed", "user_repo", u.Db.Where("user_name = ?", user.UserName).Updates(user))
}

func NoRepeat(msg, where string, d *gorm.DB) error {
	var e = errors.New(msg)
	if d.RowsAffected == 0 {
		es := errs.NewErrs(errs.GetDB())
		es.ReciteErrors(errs.Err{
			//ErrBasic: e,
			Message: msg,
			When:    time.Now(),
			Where:   where,
		})
		return e
	}
	return nil
}
