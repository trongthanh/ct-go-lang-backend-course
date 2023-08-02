package main

import (
	"thanhtran/controller"
	"thanhtran/pkg/auth"
	imagebucket "thanhtran/pkg/bucket"
	"thanhtran/pkg/validator"
	userstore "thanhtran/store"

	"thanhtran/usecase"

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

	// Static
	e.Static("/", "public")

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	public := e.Group("/api/public")
	private := e.Group("/api/private")

	private.Use(auth.GetEchoJwtMiddleware(), auth.ExtraJwtMiddleware)

	public.POST("/register", hdl.Register)
	public.POST("/login", hdl.Login)

	private.GET("/self", hdl.Self)
	private.POST("/upload-image", hdl.UploadImage)

	return e
}
