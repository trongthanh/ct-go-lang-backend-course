package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("ct-secret-key")

const jwtIssuer = "ct-backend-course"

func GenerateToken(username string, expireDuration time.Duration) (string, error) {

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		Issuer:    jwtIssuer,
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}
