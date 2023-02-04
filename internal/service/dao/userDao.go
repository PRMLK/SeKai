package dao

import (
	"SeKai/internal/controller/api/bcrypt"
	"SeKai/internal/logger"
	"SeKai/internal/model"
	"SeKai/internal/model/param"
	"SeKai/internal/util"
	"errors"
	"gorm.io/gorm"
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
		logger.ServerLogger.Debug("Hash错误: " + err.Error())
		return errors.New("系统错误")
	} else {
		hashPassword = tempHashPassword
	}
	result := util.Datebase.Where(&model.User{Username: registerParam.Username}).Find(&model.User{})
	if result.RowsAffected > 0 {
		logger.ServerLogger.Info("用户名已被注册: " + registerParam.Username)
		return errors.New("用户名已被注册")
	}
	if err := util.Datebase.Create(&model.User{Username: registerParam.Username, Email: registerParam.Email, Password: hashPassword}).Error; err != nil {
		return errors.New("系统错误")
	}
	return nil
}

func GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := util.Datebase.Where(&model.User{Username: username}).Find(&user).Error
	if err != nil {
		logger.ServerLogger.Debug("GetUserCode:" + err.Error())
		return model.User{}, errors.New("系统错误")
	}
	return user, nil
}

func GetUserByID(userId uint) (model.User, error) {
	var user model.User
	err := util.Datebase.Where(&model.User{
		Model: gorm.Model{
			ID: userId,
		},
	}).Find(&user).Error
	if err != nil {
		logger.ServerLogger.Debug("GetUserCode:" + err.Error())
		return model.User{}, errors.New("系统错误")
	}
	return user, nil
}

func SetUserAuthSecret(userId uint, secret string) error {
	err := util.Datebase.Model(&model.User{
		Model: gorm.Model{
			ID: userId,
		},
	}).Updates(&model.User{
		GoogleAuthSecret: secret,
	}).Error
	if err != nil {
		logger.ServerLogger.Debug(err)
		return errors.New("系统错误")
	}
	return nil
}

func UpdateUser(userId uint, userProfileParam *param.UserProfileUpdateParam) error {
	if err := util.Datebase.Model(&model.User{
		Model: gorm.Model{
			ID: userId,
		},
	}).Updates(&model.User{
		ProfilePhoto: userProfileParam.ProfilePhoto,
		SiteUrl:      userProfileParam.SiteUrl,
		Email:        userProfileParam.Email,
		LastName:     userProfileParam.LastName,
		Bio:          userProfileParam.Bio,
		FirstName:    userProfileParam.FirstName,
		Language:     userProfileParam.Language,
		Nickname:     userProfileParam.Nickname,
	}).Error; err != nil {
		logger.ServerLogger.Debug(err)
		return errors.New("系统错误")
	}
	return nil
}

func GetUserList() []model.User {
	var userList []model.User
	util.Datebase.Model(&model.User{}).Find(&userList)
	return userList
}
