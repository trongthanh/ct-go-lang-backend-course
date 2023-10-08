package main

import (
	"context"
	"fmt"
	"gosocial/config"
	"gosocial/internal/controller"
	mongostore "gosocial/internal/store/mongo"
	"gosocial/internal/usecase"
	"gosocial/pkg/auth"
	googlestore "gosocial/pkg/bucket/google"
	"gosocial/pkg/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	config := config.LoadConfig()

	log.EnableColor()

	db, mongoErr := mongostore.Connect(config.MongoURI, config.MongoDB)
	if mongoErr != nil {
		fmt.Println("Error connecting to MongoDB:", mongoErr)
		log.Error(mongoErr)
	}
	fmt.Println("Connect to MongoDB:", config.MongoURI)
	userStore := mongostore.NewUserStore(db, "users")
	profileStore := mongostore.NewProfileStore(db, "profiles")
	postStore := mongostore.NewPostStore(db, "posts")
	imageStore := mongostore.NewImageStore(db, "images")
	imgBucket := googlestore.New(context.TODO(), config.GoogleBucketName, config.GoogleCredFile)
	fmt.Println("Connect to Google Cloud Storage:", config.GoogleBucketName)
	uc := usecase.New(config, userStore, profileStore, postStore, imageStore, imgBucket)
	hdl := controller.New(uc)

	srv := createServer(hdl)
	if err := srv.Start(":" + config.Port); err != nil {
		log.Error(err)
	}
}

func createServer(hdl *controller.Handler) *echo.Echo {
	e := echo.New()
	e.Validator = validator.New()
	e.Use(middleware.CORS())
	// Static
	e.Static("/*", "www/build")
	// SPA
	e.File("/auth", "www/build/index.html")

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${status} ${uri}, error=${error} \n",
	}))

	public := e.Group("/api/public")
	private := e.Group("/api/private")

	private.Use(auth.GetEchoJwtMiddleware(), auth.ExtraJwtMiddleware)

	public.POST("/user/signup", hdl.Register)
	public.POST("/user/login", hdl.Login)

	private.GET("/user/me", hdl.Self)

	// profiles
	public.GET("/profiles", hdl.GetProfiles)

	// posts
	private.GET("/post/all", hdl.GetPosts)
	private.GET("/post/:userid", hdl.GetPostsByUser)
	private.POST("/post/create", hdl.CreatePost)
	private.DELETE("/post/delete/:postid", hdl.DeletePost)
	private.POST("/post/like/:postid", hdl.LikePost)
	return e
}
