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
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			err = errors.New("token为空")
			// 这里必须有err.Error()，否则不会输出错误信息
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
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
		}
		c.Next()
	}
}
