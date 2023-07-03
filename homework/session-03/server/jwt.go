package main

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("ct-secret-key")

func GenerateToken(username string, expireDuration time.Duration) (string, error) {

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		Issuer:    "ct-backend-course",
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

func ValidateToken(tokenString string) (string, error) {

	// Parse the JWT token with custom claims
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil // Replace "your-secret-key" with your actual secret key used for signing the JWT tokens
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", jwt.ValidationError{}
	}

	// Extract the username from the claims
	username := claims.Subject

	return username, nil
}
