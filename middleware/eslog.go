package middleware

import (
	"bookstore/es"
	"github.com/gin-gonic/gin"
)

func LogToEsMiddleware() gin.HandlerFunc {
	return es.TransmitLogToEs()
}
