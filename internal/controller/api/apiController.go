package api

import "github.com/gin-gonic/gin"

func APIController(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		userAPIController(v1)
		postAPIController(v1)
	}
}
