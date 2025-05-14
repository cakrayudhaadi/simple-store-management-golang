package commons

import (
	"github.com/gin-gonic/gin"
)

func ResponseSuccessWithData(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(
		statusCode,
		ApiResponse{
			Success: true,
			Message: message,
			Data:    data,
		},
	)
}

func ResponseSuccessWithoutData(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(
		statusCode,
		ApiResponse{
			Success: true,
			Message: message,
		},
	)
}

func ResponseError(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(
		statusCode,
		ApiResponse{
			Success: true,
			Message: message,
		},
	)
}
