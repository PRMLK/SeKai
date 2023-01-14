package dao

import (
	"SeKai/internal/controller/api/utils"
	"SeKai/internal/model"
	"SeKai/internal/model/param"
	"SeKai/internal/util"
	"errors"
)

func UserLogin(loginParam param.LoginParam) (userID uint, err error) {
	var user model.User
	err = util.Datebase.Where(&model.User{UserName: loginParam.Username}).Find(&model.User{}).First(&user).Error
	if err != nil {
		return 0, err
	}
	if !utils.PasswordVerify(loginParam.Password, user.Password) {
		return 0, errors.New("密码错误")
	}
	return user.ID, nil
}
