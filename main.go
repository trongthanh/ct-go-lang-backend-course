package main

import (
	"fmt"
	"gosocial/config"
	"gosocial/internal/controller"
	mongostore "gosocial/internal/store/mongo"
	"gosocial/internal/usecase"
	"gosocial/pkg/auth"
	imagebucket "gosocial/pkg/bucket"
	"gosocial/pkg/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	config := config.Config{
		Scheme:           "http://",
		Host:             "localhost",
		Port:             "8090",
		MongoURI:         "mongodb://localhost:27017",
		MongoDB:          "gocourse_db",
		GoogleCredFile:   "",
		GoogleBucketName: "",
	}

	fmt.Println("Connect to MongoDB:", config.MongoURI)
	db, mongoErr := mongostore.Connect(config.MongoURI, config.MongoDB)
	if mongoErr != nil {
		fmt.Println("Error connecting to MongoDB:", mongoErr)
		log.Error(mongoErr)
	}
	userStore := mongostore.NewUserStore(db, "users")
	profileStore := mongostore.NewProfileStore(db, "profiles")
	imageStore := mongostore.NewImageStore(db, "images")
	imgBucket := imagebucket.New()
	uc := usecase.New(config, userStore, profileStore, imageStore, imgBucket)
	hdl := controller.New(uc)

	srv := createServer(hdl)
	if err := srv.Start(":" + config.Port); err != nil {
		log.Error(err)
	}
}

func createServer(hdl *controller.Handler) *echo.Echo {
	e := echo.New()
	e.Validator = validator.New()

	// Static
	e.Static("/*", "www/build")
	// SPA
	e.File("/auth", "www/build/index.html")

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	public := e.Group("/api/public")
	private := e.Group("/api/private")

	private.Use(auth.GetEchoJwtMiddleware(), auth.ExtraJwtMiddleware)

	public.POST("/user/signup", hdl.Register)
	public.POST("/user/login", hdl.Login)

	private.GET("/self", hdl.Self)
	// private.POST("/upload-image", hdl.UploadImage)
	// private.POST("/change-password", hdl.ChangePassword)

	return e
}
