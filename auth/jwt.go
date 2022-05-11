package auth

import (
	"bookstore/global"
	"bookstore/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JwtSetting.Secret)
}

func GenerateToken(username, password string) (string, error) {
	expireTime := time.Now().Add(global.JwtSetting.Expire)
	claims := Claims{
		UserName: utils.EncodeMd5(username),
		Password: utils.EncodeMd5(password),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    global.JwtSetting.Issuer,
		},
	}
	//token, err2 := jwt.New(jwt.SigningMethodHS256).SignedString(GetJWTSecret())
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
		//  tokenClaims.Claims.(*Claims)    类型断言  x.(y)
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			fmt.Println("ParseToken: ", claims.UserName)
			return claims, nil
		}
	}

	return nil, err
}
