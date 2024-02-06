package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}
func SuccessRe(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, 200, data, msg)
}
func FalseRe(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
func ResponseWithoutGin(httpStatus int, code int, data gin.H) {
	var ctx *gin.Context
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data})
}
