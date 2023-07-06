package main

import (
	"fmt"
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

func ValidateJWTToken(tokenString string) (string, error) {
	// Define the expected issuer and secret key
	expectedIssuer := "ct-backend-course"
	secretKey := mySigningKey // Replace "your-secret-key" with your actual secret key used for signing the JWT tokens

	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method of the token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse JWT token: %v", err)
	}

	// Verify the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	subject := token.Claims.(jwt.MapClaims)["sub"].(string)

	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Verify the issuer claim
	if issuer, ok := claims["iss"].(string); !ok || issuer != expectedIssuer {
		return "", fmt.Errorf("invalid token issuer")
	}

	// Verify the expiration time claim
	if ok := claims.VerifyExpiresAt(time.Now().Unix(), true); !ok {
		return "", fmt.Errorf("token has expired")
	}

	return subject, nil
}
