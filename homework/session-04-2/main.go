package main

import (
	"thanhtran-s04-2/controller"
	imagebucket "thanhtran-s04-2/pkg/bucket"
	"thanhtran-s04-2/pkg/validator"
	userstore "thanhtran-s04-2/store"

	"thanhtran-s04-2/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	userStore := userstore.New()
	imgBucket := imagebucket.New()
	// QUESTION: how do we know userStore satisfy UserStore interface?
	uc := usecase.New(userStore, imgBucket)
	hdl := controller.New(uc)

	srv := createServer(hdl)
	if err := srv.Start(":8090"); err != nil {
		log.Error(err)
	}
}

func createServer(hdl *controller.Handler) *echo.Echo {
	e := echo.New()
	e.Validator = validator.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	public := e.Group("/api/public")
	private := e.Group("/api/private")

	public.POST("/register", hdl.Register)
	public.POST("/login", hdl.Login)

	private.GET("/self", hdl.Self)
	private.POST("/self", hdl.UploadImage)

	return e
}
