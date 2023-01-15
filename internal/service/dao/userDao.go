package dao

import (
	"SeKai/internal/controller/api/bcrypt"
	"SeKai/internal/logger"
	"SeKai/internal/model"
	"SeKai/internal/model/param"
	"SeKai/internal/util"
	"errors"
)

func UserLogin(loginParam *param.LoginParam) (userID uint, err error) {
	var user model.User
	err = util.Datebase.Where(&model.User{Username: loginParam.Username}).Find(&model.User{}).First(&user).Error
	if err != nil {
		logger.ServerLogger.Info("数据库查询: " + err.Error())
		return 0, errors.New("用户名或密码错误")
	}
	if !bcrypt.PasswordVerify(loginParam.Password, user.Password) {
		return 0, errors.New("用户名或密码错误")
	}
	return user.ID, nil
}

func UserRegister(registerParam *param.RegisterParam) error {
	var hashPassword string
	if tempHashPassword, err := bcrypt.PasswordHash(registerParam.Password); err != nil {
		logger.ServerLogger.Warning("Hash错误: " + err.Error())
		return errors.New("未知错误")
	} else {
		hashPassword = tempHashPassword
	}
	result := util.Datebase.Where(&model.User{Username: registerParam.Username}).Find(&model.User{})
	if result.RowsAffected > 0 {
		logger.ServerLogger.Info("用户名已被注册: " + registerParam.Username)
		return errors.New("用户名已被注册")
	}
	if err := util.Datebase.Create(&model.User{Username: registerParam.Username, Email: registerParam.Email, Password: hashPassword}).Error; err != nil {
		logger.ServerLogger.Warning("数据库插入错误: " + err.Error())
		return errors.New("未知错误")
	}
	return nil
}
