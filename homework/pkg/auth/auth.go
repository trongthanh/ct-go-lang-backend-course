package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
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

func GetEchoJwtMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		ContextKey: "auth",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			fmt.Println("error:", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		},
		SigningKey: mySigningKey,
	}

	return echojwt.WithConfig(config)
}

/**
 * This middleware help check issuer mismatch
 * and store the username in the Context
 */
func ExtraJwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		auth := c.Get("auth").(*jwt.Token)
		claims := auth.Claims.(*jwt.RegisteredClaims)
		username := claims.Subject
		issuer := claims.Issuer

		if issuer != jwtIssuer {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token. Issuer mismatch")
		}

		// store the username in the context for private handlers
		c.Set("username", username)

		next(c)

		return nil
	}
}
