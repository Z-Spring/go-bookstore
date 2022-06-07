package middleware

import (
	"bookstore/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			err   error
		)
		if s, exist := c.GetQuery("Authorization"); exist {
			token = s
		} else {
			token = c.GetHeader("Authorization")
		}
		if token == "" {
			err = errors.New("未登录，请登录后查看")
			// 这里必须有err.Error()，否则不会输出错误信息
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			_, err := auth.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					UnauthorizedTokenTimeout := errors.New("token已超时")
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": UnauthorizedTokenTimeout.Error()})
				default:
					UnauthorizedTokenError := errors.New("token与实际不符")
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": UnauthorizedTokenError.Error()})
				}
			}
			// TODO 要不要将用户信息返回
			//c.JSON(http.StatusOK, gin.H{"userName": claims.UserName})
		}
		c.Next()
	}
}
