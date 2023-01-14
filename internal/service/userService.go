package service

import (
	response "SeKai/internal/controller/api/utils"
	"SeKai/internal/logger"
	"SeKai/internal/model/param"
	"SeKai/internal/service/logic"
	"SeKai/internal/util"
	"github.com/gin-gonic/gin"
)

func LoginService(c *gin.Context) {
	var userID uint
	loginParam := new(param.LoginParam)
	if err := c.ShouldBindJSON(&loginParam); err != nil {
		logger.ServerLogger.Info(err)
		response.Fail(c, nil, "传递参数有误")
		return
	}

	if tempUserID, err := logic.UserLogin(loginParam); err != nil {
		logger.ServerLogger.Info(err)
		response.Fail(c, nil, "用户名或密码错误")
		return
	} else {
		userID = tempUserID
	}

	if tokenString, err := util.ReleaseToken(userID); err != nil {
		logger.ServerLogger.Info(err)
		response.Fail(c, nil, "未知错误")
		return
	} else {
		response.Success(c, gin.H{"token": tokenString}, "登陆成功")
	}
}
