package service

import (
	"SeKai/internal/config"
	response "SeKai/internal/controller/api/utils"
	"SeKai/internal/logger"
	"SeKai/internal/model/param"
	"SeKai/internal/service/dao"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewPostService(c *gin.Context) {
	newPostParam := new(param.PostParam)
	// 参数检查
	if err := c.ShouldBindJSON(&newPostParam); err != nil {
		logger.ServerLogger.Info("Bind error: " + err.Error())
		response.Fail(c, nil, "传递参数有误")
		return
	}
	// post逻辑
	if postId, err := dao.NewPost(newPostParam, uint(c.GetInt64("userId"))); err != nil {
		logger.ServerLogger.Debug(config.LanguageConfig.ServerLogger.NewPostError + ": " + err.Error())
		response.Fail(c, nil, err.Error())
		return
	} else {
		response.Success(c, gin.H{"PostId": postId}, "创建成功")
	}
}

func ShowPostService(c *gin.Context) {
	postIdString, _ := c.Params.Get("id")

	postIdInt, err := strconv.ParseInt(postIdString, 10, 64)
	if err != nil {
		response.Fail(c, nil, "postId有误")
		return
	}

	if post, err := dao.GetPost(uint(postIdInt)); err != nil {
		response.Fail(c, nil, err.Error())
	} else {
		response.Success(c, gin.H{"post": post}, "查询成功")
	}
}

func EditPostService(c *gin.Context) {
	postIdString, _ := c.Params.Get("id")
	postParam := new(param.PostParam)

	// 参数检查
	if err := c.ShouldBindJSON(&postParam); err != nil {
		logger.ServerLogger.Info("Bind error: " + err.Error())
		response.Fail(c, nil, "传递参数有误")
		return
	}
	postIdInt, err := strconv.ParseInt(postIdString, 10, 64)
	if err != nil {
		response.Fail(c, nil, "postId有误")
		return
	}

	// 查询是否存在该post
	if _, err := dao.GetPost(uint(postIdInt)); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	// 更新
	if err := dao.UpdatePost(uint(postIdInt), postParam, uint(c.GetInt64("userId"))); err != nil {
		response.Fail(c, nil, err.Error())
	} else {
		response.Success(c, nil, "更新成功")
	}
}

func DelPostService(c *gin.Context) {
	postIdString, _ := c.Params.Get("id")

	// 参数检查
	postIdInt, err := strconv.ParseInt(postIdString, 10, 64)
	if err != nil {
		response.Fail(c, nil, "postId有误")
		return
	}

	// 查询是否存在该post
	if _, err := dao.GetPost(uint(postIdInt)); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	// 删除
	if err := dao.DelPost(uint(postIdInt)); err != nil {
		response.Fail(c, nil, err.Error())
	} else {
		response.Success(c, nil, "更新成功")
	}
}
