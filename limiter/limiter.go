package limiter

import (
	"golang.org/x/time/rate"
	"time"
)

type MyLimiter struct {
	Limiter      *rate.Limiter
	LastGetToken time.Time
	RoutePath    string
	RoutePathLimiter
}

type RoutePathLimiter struct {
	LimiterBuckets map[string]int
}

func (m MyLimiter) GetLimiter() *MyLimiter {
	_, ok := m.LimiterBuckets[m.RoutePath]
	if !ok {
		m.LimiterBuckets[m.RoutePath] = m.Limiter.Burst()
		return &m
	}
	return &m
}

func (m MyLimiter) GetRoutePathBuckets(routePath string) (bucket int, ok bool) {
	bucket, ok = m.LimiterBuckets[routePath]
	return
}

func (m MyLimiter) Allow() bool {
	m.LastGetToken = time.Now()
	return m.Limiter.Allow()
}
