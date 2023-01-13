package chunkLoader

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func InitLoader(router *gin.Engine) {
	templates := template.New("")
	inlineAssetsLoader(templates)
	outlineAssetsLoader(templates)
	router.SetHTMLTemplate(templates)
}
