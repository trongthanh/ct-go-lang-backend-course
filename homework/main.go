package main

import (
	"fmt"
	"thanhtran/config"
	"thanhtran/internal/controller"
	mongostore "thanhtran/internal/store/mongo"
	"thanhtran/internal/usecase"
	"thanhtran/pkg/auth"
	imagebucket "thanhtran/pkg/bucket"
	"thanhtran/pkg/validator"

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
		MongoCollUser:    "users",
		MongoCollImage:   "images",
		GoogleCredFile:   "",
		GoogleBucketName: "",
	}

	fmt.Println("Connect to MongoDB:", config.MongoURI)
	db, mongoErr := mongostore.Connect(config.MongoURI, config.MongoDB)
	if mongoErr != nil {
		fmt.Println("Error connecting to MongoDB:", mongoErr)
		log.Error(mongoErr)
	}
	userStore := mongostore.NewUserStore(db, config.MongoCollUser)
	imageStore := mongostore.NewImageStore(db, config.MongoCollImage)
	imgBucket := imagebucket.New()
	uc := usecase.New(config, userStore, imageStore, imgBucket)
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
