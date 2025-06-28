package utils

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID uint, role string) (string, error) {
	exp := os.Getenv("JWT_EXPIRATION")
	duration, _ := time.ParseDuration(exp)
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(duration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

func ExtractRoleFromJWT(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("missing auth header")
	}
	tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return "", err
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}

	return role, nil
}

func ExtractUserIDFromJWT(c *gin.Context) (uint, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return 0, errors.New("missing auth header")
	}
	tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user_id in token")
	}

	return uint(userIDFloat), nil
}
