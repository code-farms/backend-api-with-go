package auth

import (
	"strconv"
	"time"

	"github.com/code-farms/go-backend/configs"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret) 
	if err != nil {
		return "", err
	}

	return tokenString, nil
}