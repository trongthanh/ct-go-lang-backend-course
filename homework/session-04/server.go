package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(LoggerMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/api/public/register", register)
	e.POST("/api/public/login", login)

	private := e.Group("/api/private")
	// example taken from: https://github.com/labstack/echox/blob/master/cookbook/jwt/custom-claims/server.go
	config := echojwt.Config{
		ContextKey: "auth",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			fmt.Println("error:", err)
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		},
		SigningKey: mySigningKey,
	}

	private.Use(echojwt.WithConfig(config), jwtMiddleware)

	private.GET("/self", self)

	e.Logger.Fatal(e.Start(":8090"))
}

/**
 * This middleware help check issuer mismatch
 * and store the username in the Context
 */
func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func register(c echo.Context) (err error) {

	user := new(UserInfo)
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(user); err != nil {
		return err
	}

	if err = userStore.Save(*user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string
}

func login(c echo.Context) (err error) {

	loginReq := new(LoginRequest)
	if err = c.Bind(loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user UserInfo

	user, err = userStore.Get(loginReq.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid username or password")
	}

	if user.Password != loginReq.Password {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid username or password")
	}

	var token string
	token, err = GenerateToken(user.Username, 24*time.Hour)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := LoginResponse{Token: token}

	return c.JSON(http.StatusOK, resp)
}

func self(c echo.Context) (err error) {
	username := c.Get("username").(string)
	var user UserInfo
	user, err = userStore.Get(username)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error: getting user info")
	}

	return c.JSON(http.StatusOK, user)
}

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Start timer
		start := time.Now()

		// Call the next handler
		err := next(c)

		// Calculate request duration
		duration := time.Since(start)

		status := c.Response().Status
		if err != nil {
			// QUESTION:
			status = err.(*echo.HTTPError).Code
		}

		// Log the HTTP event
		logMessage := fmt.Sprintf("[%s] %s %s - %d %s",
			time.Now().Format("2006-01-02 15:04:05"),
			c.Request().Method,
			c.Request().URL.Path,
			status,
			duration.String(),
		)
		fmt.Println(logMessage)

		return err
	}
}
