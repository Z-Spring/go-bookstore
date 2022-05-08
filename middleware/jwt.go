package middleware

import (
	"bookstore/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func JWT() gin.HandlerFunc {
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})

		} else {
			_, err := auth.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					UnauthorizedTokenTimeout := errors.New("token已超时")
					c.JSON(http.StatusInternalServerError, gin.H{"error": UnauthorizedTokenTimeout})
				default:
					UnauthorizedTokenError := errors.New("token与实际不符")
					c.JSON(http.StatusInternalServerError, gin.H{"error": UnauthorizedTokenError})

				}
			}
		}
		//todo 什么意思？
		c.Next()
	}
}
