package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func backStageController(router *gin.Engine) {
	router.GET("backstage/user/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", gin.H{
			"sekaiPageTitle": "Home",
			"sekaiSiteRoot":  "localhost:12070",
		})
	})
}
