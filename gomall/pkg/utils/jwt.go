package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/yqihe/91-mall/gomall/pkg/constants"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(username string) (string, error) {
	expireTime := time.Now().Add(24 * 7 * time.Hour) // 过期时间为7天
	nowTime := time.Now()                            // 当前时间
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间戳
			IssuedAt:  nowTime.Unix(),    // 当前时间戳
			Issuer:    "91-mall",         // 颁发者签名
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenStruct.SignedString([]byte(constants.JWTValue))
}
