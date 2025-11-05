package tools

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func ParseToken(tokenString string, secret string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		val := claims["userId"].(float64)
		exp := int64(claims["exp"].(float64))
		if exp <= time.Now().Unix() {
			return 0, errors.New("token expired")
		}
		return int64(val), nil
	} else {
		return 0, errors.New("token invalid")
	}
}
