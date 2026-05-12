package util

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/ini.v1"
)

type JwtClaims struct {
	UserId int    `json:"userId"`
	Phone  string `json:"phone"`
	jwt.RegisteredClaims
}

func getJwtSecret() []byte {
	config, err := ini.Load("./config/app.ini")
	if err != nil {
		return []byte("shopwebgo-jwt-secret")
	}
	secret := config.Section("").Key("jwt_secret").String()
	if secret == "" {
		return []byte("shopwebgo-jwt-secret")
	}
	return []byte(secret)
}

func GenerateToken(userId int, phone string) (string, error) {
	claims := JwtClaims{
		UserId: userId,
		Phone:  phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJwtSecret())
}

func ParseToken(tokenStr string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return getJwtSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func GetUserFromJWT(c *gin.Context) (int, string, bool) {
	tokenStr, err := c.Cookie("token")
	if err != nil || tokenStr == "" {
		return 0, "", false
	}
	claims, err := ParseToken(tokenStr)
	if err != nil {
		return 0, "", false
	}
	return claims.UserId, claims.Phone, true
}
