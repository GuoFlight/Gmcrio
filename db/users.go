package db

import (
	"Gmicro/logger"
	"Gmicro/utils"
	"context"
	"github.com/GuoFlight/gerror/v2"
)

var GUser User

type User struct {
	ID       int64
	Username string
	PwdHash  string `gorm:"column:pwd" json:"-"`
	Disable  bool
	Comment  string
}

func (User) TableName() string {
	return "user"
}
func (User) GetByUsername(ctx context.Context, username string) (User, bool, *gerror.Gerr) {
	var result User
	err := Conn.Where("username = ?", username).Find(&result).Error
	if err != nil {
		if err.Error() == "record not found" {
			return User{}, false, nil
		}
		return User{}, false, logger.ErrWithCtx(ctx, err.Error())
	}
	return result, true, nil
}
func (u User) CheckPwd(ctx context.Context, username, pwd string) (bool, *gerror.Gerr) {
	user, ok, gerr := u.GetByUsername(ctx, username)
	if gerr != nil || ok == false {
		return ok, gerr
	}
	if user.Disable {
		return false, nil
	}
	return utils.GBcrypt.CheckPwd(pwd, user.PwdHash), nil
}
