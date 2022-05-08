package auth

import (
	"bookstore/global"
	"bookstore/utils"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type Claims struct {
	JwtSecret string `json:"jwtSecret,omitempty"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JwtSetting.Secret)
}

func GenerateToken(jwtSecret string) (string, error) {
	log.Println(global.JwtSetting.Expire)
	expireTime := time.Now().Add(global.JwtSetting.Expire)
	claims := Claims{
		JwtSecret: utils.EncodeMd5(jwtSecret),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    global.JwtSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
