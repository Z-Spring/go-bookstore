package global

import "time"

type JwtSettings struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type ServerSettings struct {
	RunMode string
}

var (
	JwtSetting    *JwtSettings
	ServerSetting *ServerSettings
)
