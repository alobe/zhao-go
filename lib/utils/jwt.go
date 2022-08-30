package utils

import (
	"time"
	"zhao-go/config"

	"github.com/dgrijalva/jwt-go"
)

func getSecret() (secret []byte) {
	secret = []byte(config.Config.Jwt.Secret)
	return
}

type Claims struct {
	entity interface{}
	jwt.StandardClaims
}

func GetToken(entity interface{}) string {
	expiresAt := time.Now().Add(time.Hour * 6).Unix()
	claims := Claims{
		entity: entity,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    config.Config.Jwt.Issuer + ":",
		},
	}
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(getSecret())
	return token
}

func parseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getSecret(), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
