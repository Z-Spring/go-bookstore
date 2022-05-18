package global

import (
	"golang.org/x/time/rate"
	"time"
)

type JwtSettings struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type ServerSettings struct {
	RunMode string
}

type RateLimiterSettings struct {
	RoutePath string
	RateLimit rate.Limit
	Buckets   int
}

var (
	JwtSetting         *JwtSettings
	ServerSetting      *ServerSettings
	RateLimiterSetting RateLimiterSettings
)
