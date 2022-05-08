package global

import "time"

type JwtSettings struct {
	Secret string
	Issuer string
	Expire time.Duration
}

var JwtSetting *JwtSettings
