package limiter

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type MyLimiter struct {
	Limiter      *rate.Limiter
	LastGetToken time.Time
	RoutePath    string
	RoutePathLimiter
}

type RoutePathLimiter struct {
	//LimiterBuckets map[string]int
	LimiterBuckets sync.Map
}

func (m *MyLimiter) GetLimiter() *MyLimiter {
	//_, ok := m.LimiterBuckets[m.RoutePath]
	_, ok := m.LimiterBuckets.Load(m.RoutePath)
	if !ok {
		//m.LimiterBuckets[m.RoutePath] = m.Limiter.Burst()
		m.LimiterBuckets.Store(m.RoutePath, m.Limiter.Burst())
		return m
	}
	return m
}

func (m *MyLimiter) GetRoutePathBuckets(routePath string) (ok bool) {
	//bucket, ok = m.LimiterBuckets[routePath]
	_, ok = m.LimiterBuckets.Load(routePath)
	return
}

func (m *MyLimiter) Allow() bool {
	m.LastGetToken = time.Now()
	return m.Limiter.Allow()
}
