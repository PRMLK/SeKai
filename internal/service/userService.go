package service

import (
	response "SeKai/internal/controller/api/utils"
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
	if tempUserID, err := dao.UserLogin(loginParam); err != nil {
		response.Fail(c, nil, err.Error())
		return
	} else {
		userID = tempUserID
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
