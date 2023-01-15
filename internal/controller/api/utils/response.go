package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	code: xxx,
	data: xxx,
	msg: xxx
}
*/

func response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	response(ctx, http.StatusOK, 0, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	response(ctx, http.StatusOK, 400, data, msg)
}

func Unauthorized(ctx *gin.Context, data gin.H, msg string) {
	response(ctx, http.StatusUnauthorized, 400, data, msg)
}
