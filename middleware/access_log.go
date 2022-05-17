package middleware

import (
	"bookstore/mylog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Log(logger *zap.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return mylog.GinLogWithConfig(logger, &mylog.Config{TimeFormat: "", UTC: utc})
}
