package jwt

import (
	"errors"
	"face_ui/utils/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `jsong:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte(setting.ApplicationSetting.JwtSecret)

func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)

	claims := Claims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "face_api",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (string, error) {
	if token == "" {
		return "", errors.New("no token is found")
	}

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims.Username, nil
		}
	}
	return "", err
}
