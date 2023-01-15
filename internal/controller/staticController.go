package controller

import "github.com/gin-gonic/gin"

func staticController(router *gin.Engine) {
	router.StaticFile("/stylesheets/styles.css", "./themes/frontStage/plain-sekai/template/header/stylesheets/styles.css")
	router.StaticFile("/backstage/user/stylesheets/styles.css", "./themes/backStage/plain-sekai/template/header/stylesheets/styles.css")
}
