package middleware

import (
	"bookstore/limiter"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RateLimiter(m *limiter.MyLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := m.GetLimiter()
		ok := m.GetRoutePathBuckets(m.RoutePath)
		if ok {
			if !l.Allow() {
				log.Println(m.RoutePath, ": 限流啦")
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": m.RoutePath + " 被限流"})
			}
		}
		c.Next()
	}

}
