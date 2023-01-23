package service

import (
	"SeKai/internal/config"
	response "SeKai/internal/controller/api/utils"
	"SeKai/internal/controller/api/utils/googleAuth"
	"SeKai/internal/logger"
	"SeKai/internal/model/param"
	"SeKai/internal/service/dao"
	"SeKai/internal/util"
	"github.com/gin-gonic/gin"
)

/*
	错误处理逻辑：
	被发现的异常或错误应当立刻被写入Logger以供调试
	需要返回给请求者的信息可以通过Error向上传递
*/

func LoginService(c *gin.Context) {
	var userID uint
	loginParam := new(param.LoginParam)
	// 参数检查
	if err := c.ShouldBindJSON(&loginParam); err != nil {
		logger.ServerLogger.Info("Bind error: " + err.Error())
		response.Fail(c, nil, "传递参数有误")
		return
	}

	// 登录逻辑
	if loginParam.Type == "password" {
		if tempUserID, err := dao.UserLogin(loginParam); err != nil {
			response.Fail(c, nil, err.Error())
			return
		} else {
			userID = tempUserID
		}
	} else if loginParam.Type == "googleAuth" {
		user, err := dao.GetUserByUsername(loginParam.Username)
		if err != nil {
			response.Fail(c, nil, "用户名或Code错误")
		}
		if !googleAuth.VerifyCode(util.HashGoogleSecret(user.GoogleAuthSecret), loginParam.Code) {
			response.Fail(c, nil, "Code不正确")
			return
		}
	} else {
		response.Fail(c, nil, "传递Type有误")
	}

	// 释放token
	if tokenString, err := util.ReleaseToken(userID); err != nil {
		logger.ServerLogger.Info(err)
		response.Fail(c, nil, err.Error())
		return
	} else {
		response.Success(c, gin.H{"token": tokenString}, "登陆成功")
	}
}

func RegisterService(c *gin.Context) {
	registerParam := new(param.RegisterParam)
	// 参数检查
	if err := c.ShouldBindJSON(&registerParam); err != nil {
		logger.ServerLogger.Info(err)
		response.Fail(c, nil, "传递参数有误")
		return
	}

	// 注册逻辑
	if err := dao.UserRegister(registerParam); err != nil {
		response.Fail(c, nil, err.Error())
		return
	} else {
		response.Success(c, nil, "注册成功")
	}
}

func SetGoogleAuthSecret(c *gin.Context) {
	newSecret := googleAuth.GetSecret()
	userId := uint(c.GetInt64("userId"))

	if err := dao.SetUserAuthSecret(userId, newSecret); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	user, err := dao.GetUserByID(userId)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, gin.H{
		"secret": "otpauth://totp/" + user.Username + "?secret=" + util.HashGoogleSecret(newSecret) + "&issuer=" + config.ApplicationConfig.SiteConfig.SiteName,
	}, "设置成功")
}

func Profile(c *gin.Context) {
	user, _ := dao.GetUserByID(uint(c.GetInt64("userId")))

	response.Success(c,
		gin.H{"profile": gin.H{
			"userId":       user.ID,
			"Username":     user.Username,
			"Nickname":     user.Nickname,
			"CreatedAt":    user.Model.CreatedAt,
			"Bio":          user.Bio,
			"Language":     user.Language,
			"FirstName":    user.FirstName,
			"LastName":     user.LastName,
			"Email":        user.Email,
			"ProfilePhoto": user.ProfilePhoto,
			"SiteUrl":      user.SiteUrl,
		}}, "查询成功")
}

func UpdateProfile(c *gin.Context) {
	userProfileParam := new(param.UserProfileUpdateParam)
	userId := uint(c.GetInt64("userId"))
	// 参数检查
	if err := c.ShouldBindJSON(&userProfileParam); err != nil {
		logger.ServerLogger.Info(err)
		response.Fail(c, nil, "传递参数有误")
		return
	}

	if err := dao.UpdateUser(userId, userProfileParam); err != nil {
		response.Fail(c, nil, err.Error())
	}
	response.Success(c, nil, "更新成功")
}
