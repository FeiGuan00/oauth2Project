package jwtutil

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTUtil struct {
	EncodingString string
}

func (j JWTUtil) getToken(appId string, appSecret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2022, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(j.EncodingString)

	fmt.Println(tokenString, err)
	if err != nil {
		return tokenString
	} else {
		return ""
	}
}

func (j JWTUtil) verifyToken(token string) {

}
