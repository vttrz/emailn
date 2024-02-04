package foundation

import (
	"github.com/gin-gonic/gin"
)

func AbortWithStatusError(ctx *gin.Context, statusCode int, err string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"error": err,
	})
	return
}

func Success(ctx *gin.Context, statusCode int, msg interface{}) {
	ctx.JSON(statusCode, msg)
	return
}
