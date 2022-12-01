package middleware

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var (
	SECRET_KEY = []byte("RAHASIA")
)

func GenerateToken(userID int, role string) (string, error) {
	claim := jwt.MapClaims{}
	claim["id_user"] = userID
	claim["role_user"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (int, string, error) {
	var userID int
	var userRole string
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(SECRET_KEY), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID = int(claims["id_user"].(float64))
		userRole = string(claims["role_user"].(string))
		return userID, userRole, nil
	} else {
		return userID, userRole, err
	}
}
