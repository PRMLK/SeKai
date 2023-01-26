package middleware

import (
	"github.com/gin-gonic/gin"
)

func LoadMiddleware(router *gin.Engine) {
	// HttpLogger
	router.Use(loggerMiddleware())
}
