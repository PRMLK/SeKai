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
	newPostParam := new(param.NewPostParam)
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

}
