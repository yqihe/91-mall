package auth

import (
	"context"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/dgrijalva/jwt-go"
	"github.com/yqihe/91-mall/gomall/pkg/constants"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iss":      "91-mall",
		"username": username,
	})
	return token.SignedString([]byte(constants.JWTValue))
}

func CheckAuth(ctx context.Context) (string, error) {
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		return "", err
	}
	var clientClaims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			return "", errno.InvalidAlgorithmError
		}
		return []byte(constants.JWTValue), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errno.InvalidTokenError
	}
	return clientClaims.Username, nil
}

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errno.NoMetadataInContextError
	}
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", errno.NoAuthorizationInMetadataError
	}
	return token[0], nil
}
