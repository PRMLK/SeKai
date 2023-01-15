package logic

import (
	"SeKai/internal/model/param"
	"SeKai/internal/service/dao"
)

func UserLogin(loginParam *param.LoginParam) (userID uint, err error) {
	userID, err = dao.UserLogin(loginParam)
	return userID, err
}

func UserRegister(registerParam *param.RegisterParam) error {
	return dao.UserRegister(registerParam)
}
