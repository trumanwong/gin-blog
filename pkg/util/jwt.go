package util

import (
	"fmt"
	"gin-blog/pkg/settings"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secret = []byte(settings.JwtSecret)

func GenerateToken(username, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"username": username,
		"password": password,
		"expire_at": time.Now().Add(3 * time.Hour).Format("2006-01-02 15:04:05"),
	})
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}