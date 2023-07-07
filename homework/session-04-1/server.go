package main

import (
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/api/public/register", register)
	e.POST("/api/public/login", login)
	e.GET("/api/private/self", self)

	e.Logger.Fatal(e.Start(":8090"))
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

	return c.String(http.StatusOK, "register")
}