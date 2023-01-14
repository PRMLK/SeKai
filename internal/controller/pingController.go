package controller

import "github.com/gin-gonic/gin"

func pingController(router *gin.Engine) {
	router.GET("/ping.sass", func(c *gin.Context) {
		_, err := c.Writer.WriteString("p {\n    color: red;\n}")
		if err != nil {
			return
		}
	})
}
