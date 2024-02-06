package middleware

import (
	"PaSer/common"
	"PaSer/model"
	"PaSer/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{"code": 401, "msg": "权限不足"}, "")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{"code": 401, "msg": "权限不足 "}, "")
			ctx.Abort()
			return
		}

		//验证通过，获取userID
		userId := claims.UserId
		DB := common.GetDB_User()
		var user model.User
		DB.First(&user, userId)

		//用户信息不存在
		if user.ID == 0 {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{"code": 401, "msg": "权限不足 "}, "")
			ctx.Abort()
			return
		}
		//用户存在，将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
